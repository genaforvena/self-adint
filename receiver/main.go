package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parsePrice(s string) float64 {
	// Guard the exchange macro left unexpanded (e.g. "${AUCTION_PRICE}").
	if s == "" || strings.ContainsAny(s, "${}") {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil || v < 0 {
		return 0
	}
	return v
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func envFloat(k string, def float64) float64 {
	if v := os.Getenv(k); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return def
}

// defaultFilter is the local stand-in for the exchange-side geo/bundle gate.
// It passes a request whose bundle is in ADINT_BUNDLE_ALLOW (comma-separated).
// Empty allowlist → passes nothing, so foreign traffic is 204'd by default.
func defaultFilter(allow string) Filter {
	set := map[string]bool{}
	for _, b := range strings.Split(allow, ",") {
		if b = strings.TrimSpace(b); b != "" {
			set[b] = true
		}
	}
	return func(r *BidRequest) bool {
		return r.App != nil && set[r.App.Bundle]
	}
}

func main() {
	cfg := &Config{
		TargetIFA: env("ADINT_TARGET_IFA", ""),
		Floor:     envFloat("ADINT_FLOOR", 0),
		SpendCap:  envFloat("ADINT_SPEND_CAP", 1.0), // $1 hard cap by default
		Source:    env("ADINT_SOURCE", "receiver"),
		Pass:      defaultFilter(env("ADINT_BUNDLE_ALLOW", "")),
	}
	if cfg.TargetIFA == "" {
		log.Fatal("ADINT_TARGET_IFA is required (the device.ifa to match)")
	}

	dataDir := env("ADINT_DATA_DIR", "data")
	_ = os.MkdirAll(dataDir, 0o755)
	sink, err := NewSink(dataDir + "/receiver.jsonl")
	if err != nil {
		log.Fatalf("open sink: %v", err)
	}
	defer sink.Close()

	ledger := NewLedger(cfg.SpendCap, dataDir+"/win-counter.txt", dataDir+"/WIN-ALERT.txt", log.Printf)
	srv := NewServer(cfg, sink, ledger, log.Printf)

	addr := env("ADINT_ADDR", "127.0.0.1:8788")
	log.Printf("adint receiver: listening on %s (cap=$%.2f, target set, allow=%q)",
		addr, cfg.SpendCap, env("ADINT_BUNDLE_ALLOW", ""))
	log.Fatal(http.ListenAndServe(addr, srv.Handler()))
}
