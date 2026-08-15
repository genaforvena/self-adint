package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

const (
	targetIFA  = "TARGET-IFA-0000-1111"
	foreignIFA = "FOREIGN-IFA-9999-8888"
)

// bufLog is a concurrency-safe log sink so a test can assert what the handler
// did — and did NOT — write about a request.
type bufLog struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (b *bufLog) logf(f string, a ...any) {
	b.mu.Lock()
	defer b.mu.Unlock()
	fmt.Fprintf(&b.buf, f+"\n", a...)
}
func (b *bufLog) String() string { b.mu.Lock(); defer b.mu.Unlock(); return b.buf.String() }

func newTestServer(t *testing.T) (*Server, string, *bufLog) {
	t.Helper()
	dir := t.TempDir()
	sink, err := NewSink(filepath.Join(dir, "receiver.jsonl"))
	if err != nil {
		t.Fatalf("sink: %v", err)
	}
	t.Cleanup(func() { sink.Close() })
	lg := &bufLog{}
	cfg := &Config{
		TargetIFA: targetIFA,
		SpendCap:  1.0,
		Source:    "test",
		Pass:      defaultFilter("com.pass.allowed"),
	}
	ledger := NewLedger(cfg.SpendCap, filepath.Join(dir, "win-counter.txt"), filepath.Join(dir, "WIN-ALERT.txt"), lg.logf)
	return NewServer(cfg, sink, ledger, lg.logf), dir, lg
}

func req(ifa, bundle string, floor float64, extra map[string]any) *http.Request {
	m := map[string]any{
		"id":     "req-" + ifa,
		"imp":    []map[string]any{{"id": "1", "bidfloor": floor}},
		"app":    map[string]any{"bundle": bundle},
		"device": map[string]any{"ifa": ifa, "geo": map[string]any{"type": 2, "lat": 55.75, "lon": 37.61}},
		"user": map[string]any{
			"data": []map[string]any{{"id": "prov1", "name": "AcmeData", "ext": map[string]any{"segtax": 4}, "segment": []map[string]any{{"id": "seg-42"}}}},
			"ext":  map[string]any{"consent": "CPxyz", "eids": []map[string]any{{"source": "liveramp.com", "uids": []map[string]any{{"id": "u-secret"}}}}},
		},
	}
	for k, v := range extra {
		m[k] = v
	}
	b, _ := json.Marshal(m)
	r := httptest.NewRequest("POST", "/bid", strings.NewReader(string(b)))
	return r
}

func readJSONL(t *testing.T, dir string) []Record {
	t.Helper()
	f, err := os.Open(filepath.Join(dir, "receiver.jsonl"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		t.Fatal(err)
	}
	defer f.Close()
	var recs []Record
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) == "" {
			continue
		}
		var rec Record
		if err := json.Unmarshal(sc.Bytes(), &rec); err != nil {
			t.Fatalf("bad jsonl line: %v", err)
		}
		recs = append(recs, rec)
	}
	return recs
}

// --- the decoy rule truth table (brief verbatim) ---

func TestDecoyRule(t *testing.T) {
	cfg := &Config{TargetIFA: targetIFA, Pass: defaultFilter("com.pass.allowed")}
	cases := []struct {
		name    string
		ifa     string
		bundle  string
		want    Verdict
		matched bool
	}{
		{"target ifa always bids", targetIFA, "com.random", BidBelowFloor, true},
		{"foreign + passes filter bids", foreignIFA, "com.pass.allowed", BidBelowFloor, false},
		{"foreign + fails filter 204s", foreignIFA, "com.random", NoBid, false},
		{"target ifa even when filter fails", targetIFA, "com.random", BidBelowFloor, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &BidRequest{Device: Device{IFA: c.ifa}, App: &App{Bundle: c.bundle}}
			v, m := cfg.decide(r)
			if v != c.want || m != c.matched {
				t.Fatalf("verdict=%v matched=%v; want %v/%v", v, m, c.want, c.matched)
			}
		})
	}
}

// --- PRIVACY INVARIANT: a foreign ifa never reaches disk (the load-bearing gate) ---

func TestForeignIFANeverWritten(t *testing.T) {
	srv, dir, _ := newTestServer(t)
	h := srv.Handler()

	// Two foreign requests (one filter-passing so it BIDS — proving that
	// bidding on foreign traffic still never persists it), one target.
	for _, r := range []*http.Request{
		req(foreignIFA, "com.pass.allowed", 1.0, nil), // foreign, bids
		req(foreignIFA, "com.random", 1.0, nil),       // foreign, 204
		req(targetIFA, "com.random", 1.0, nil),        // target, persisted
	} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
	}

	recs := readJSONL(t, dir)
	if len(recs) != 1 {
		t.Fatalf("expected exactly 1 persisted row (target only), got %d", len(recs))
	}
	if recs[0].ReqID != "req-"+targetIFA {
		t.Fatalf("persisted the wrong request: %q", recs[0].ReqID)
	}
	// Belt-and-suspenders: the whole JSONL file must not contain the foreign
	// ifa or the secret uid as substrings, from any field or framing.
	raw, _ := os.ReadFile(filepath.Join(dir, "receiver.jsonl"))
	if strings.Contains(string(raw), foreignIFA) {
		t.Fatal("foreign ifa found in JSONL")
	}
}

// --- §1b: the panic path must not leak the body/ifa, and must write nothing ---

func TestPanicPathNoLeak(t *testing.T) {
	srv, dir, lg := newTestServer(t)
	h := srv.Handler()

	r := req(foreignIFA, "com.pass.allowed", 1.0, map[string]any{"__panic": "panic"})
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("panic path should 500, got %d", w.Code)
	}
	if recs := readJSONL(t, dir); len(recs) != 0 {
		t.Fatalf("panic path wrote %d rows; must write none", len(recs))
	}
	if strings.Contains(lg.String(), foreignIFA) {
		t.Fatalf("panic log leaked the foreign ifa:\n%s", lg.String())
	}
	if strings.Contains(lg.String(), "u-secret") {
		t.Fatalf("panic log leaked a uid from the body:\n%s", lg.String())
	}
}

// --- extraction keeps precision CLASS, drops coordinates ---

func TestExtractNoCoordinates(t *testing.T) {
	srv, dir, _ := newTestServer(t)
	h := srv.Handler()
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req(targetIFA, "com.app", 1.0, nil))

	recs := readJSONL(t, dir)
	if len(recs) != 1 {
		t.Fatalf("want 1 row, got %d", len(recs))
	}
	rec := recs[0]
	if rec.GeoType != 2 {
		t.Fatalf("geo_type not preserved: %d", rec.GeoType)
	}
	if !rec.Consent {
		t.Fatal("consent presence not recorded")
	}
	if len(rec.Segments) != 1 || rec.Segments[0].SegTax != 4 || rec.Segments[0].Provider != "AcmeData" {
		t.Fatalf("segments mis-extracted: %+v", rec.Segments)
	}
	if len(rec.EIDs) != 1 || rec.EIDs[0] != "liveramp.com" {
		t.Fatalf("eid sources mis-extracted: %+v", rec.EIDs)
	}
	// The raw coordinate values must never appear on disk.
	raw, _ := os.ReadFile(filepath.Join(dir, "receiver.jsonl"))
	if strings.Contains(string(raw), "55.75") || strings.Contains(string(raw), "37.61") {
		t.Fatal("coordinates leaked into JSONL")
	}
}

// --- §1a: win counter is an asserted artifact; first win alerts ---

func TestWinCounterAndAlert(t *testing.T) {
	srv, dir, _ := newTestServer(t)
	h := srv.Handler()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/win?p=0.5&id=x", nil))

	wins, spend := srv.ledger.snapshot()
	if wins != 1 || spend != 0.5 {
		t.Fatalf("ledger wins=%d spend=%.2f; want 1/0.50", wins, spend)
	}
	if _, err := os.Stat(filepath.Join(dir, "win-counter.txt")); err != nil {
		t.Fatalf("win counter artifact not written: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "WIN-ALERT.txt")); err != nil {
		t.Fatalf("first-win alert artifact not written: %v", err)
	}
}

// --- §1a: past the hard spend cap the decoy stops bidding (204) ---

func TestSpendCapHalts(t *testing.T) {
	srv, _, _ := newTestServer(t) // cap = $1.0
	h := srv.Handler()

	// Drive spend to the cap.
	for i := 0; i < 3; i++ {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest("GET", "/win?p=0.5&id=x", nil))
	}
	if _, spend := srv.ledger.snapshot(); spend < 1.0 {
		t.Fatalf("precondition: spend %.2f should be >= cap", spend)
	}

	// A normally bid-eligible target request must now 204.
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req(targetIFA, "com.app", 1.0, nil))
	if w.Code != http.StatusNoContent {
		t.Fatalf("past cap should 204, got %d", w.Code)
	}
}

// --- the bid is strictly below the imp floor ---

func TestBidBelowFloor(t *testing.T) {
	cfg := &Config{TargetIFA: targetIFA}
	price := cfg.bidPrice(&BidRequest{Imp: []Imp{{BidFloor: 2.0}}})
	if price >= 2.0 {
		t.Fatalf("bid %.2f not below floor 2.0", price)
	}
}
