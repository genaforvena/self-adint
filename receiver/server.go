package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
)

// Sink is the JSONL artifact writer. By construction the ONLY caller is the
// matched-request path, so nothing but the target device's own rows are ever
// handed to it (§ privacy invariant — the filter is in the parser, upstream
// of this buffer, not a query on it).
type Sink struct {
	mu sync.Mutex
	f  *os.File
	en *json.Encoder
}

func NewSink(path string) (*Sink, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return &Sink{f: f, en: json.NewEncoder(f)}, nil
}

func (s *Sink) Write(rec Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.en.Encode(rec) // Encode appends a newline → JSONL
}

func (s *Sink) Close() error {
	if s.f == nil {
		return nil
	}
	return s.f.Close()
}

// Server wires the config, sink and ledger behind the bid/win endpoints.
type Server struct {
	cfg    *Config
	sink   *Sink
	ledger *Ledger
	logf   func(string, ...any)
}

func NewServer(cfg *Config, sink *Sink, ledger *Ledger, logf func(string, ...any)) *Server {
	if logf == nil {
		logf = func(string, ...any) {}
	}
	return &Server{cfg: cfg, sink: sink, ledger: ledger, logf: logf}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/bid", s.handleBid)
	mux.HandleFunc("/win", s.handleWin)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) })
	return mux
}

// handleBid is the hot path. Its recover() is the leak path the brief §1b
// warns about: a panic handler that dumps the request body would write a
// foreign ifa to a log even though the filter kept it off the JSONL. So the
// recover logs the request id at most — NEVER the body, NEVER the ifa.
func (s *Server) handleBid(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if p := recover(); p != nil {
			// Nothing request-derived is logged here — not the body, not the
			// ifa, and not even req.ID: an OpenRTB id is opaque and may embed
			// PII, so correlating on it would reopen the very leak §1b names.
			// Only our own panic value (a programmer-controlled string) is logged.
			s.logf("panic in /bid recovered: %v", p)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req BidRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Test-only panic injection (§1b). It is placed AFTER ifa is decoded so
	// the panic path runs with a foreign ifa in scope — the exact condition
	// under which a careless handler would leak it.
	if req.Test == "panic" {
		panic("synthetic panic (test harness)")
	}

	verdict, matched := s.cfg.decide(&req)

	// Persist BEFORE responding, and only for a matched request.
	if matched {
		if err := s.sink.Write(s.cfg.extract(&req)); err != nil {
			s.logf("sink write failed (req_id=%q): %v", req.ID, err)
		}
	}

	// Hard spend cap: past it we stop bidding entirely (§1a).
	if verdict == NoBid || s.ledger.capReached() {
		w.WriteHeader(http.StatusNoContent) // 204
		return
	}

	resp := s.cfg.buildBid(&req)
	// Point the win notice at our own endpoint so a win is always booked.
	resp.SeatBid[0].Bid[0].NURL = "/win?p=${AUCTION_PRICE}&id=" + req.ID
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// handleWin books the (rare) win. Price comes from the exchange's settlement
// macro; here the synthetic harness passes it as ?p=.
func (s *Server) handleWin(w http.ResponseWriter, r *http.Request) {
	price := parsePrice(r.URL.Query().Get("p"))
	over := s.ledger.recordWin(price)
	if over {
		s.logf("hard spend cap reached after win; bidding halts")
	}
	// Serve the fallback creative so a win never fails to render (§1a).
	w.Header().Set("Content-Type", "text/html")
	_, _ = io.WriteString(w, fallbackCreative)
}
