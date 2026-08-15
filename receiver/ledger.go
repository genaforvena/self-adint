package main

import (
	"fmt"
	"os"
	"sync"
)

// Ledger is the win/spend artifact (§1a). "win rate нулевой, расход нулевой"
// is an assumption about the design, not a property of it: a below-floor bid
// wins whenever no floor is applied. So the win counter is a real artifact —
// asserted by the tests, persisted to disk — not a hope, and the first
// non-zero win raises a loud, out-of-band alert.
type Ledger struct {
	mu        sync.Mutex
	Wins      int
	Spend     float64
	Cap       float64
	counterFP string
	alertFP   string
	alerted   bool
	log       func(string, ...any)
}

func NewLedger(cap float64, counterFP, alertFP string, logf func(string, ...any)) *Ledger {
	if logf == nil {
		logf = func(string, ...any) {}
	}
	return &Ledger{Cap: cap, counterFP: counterFP, alertFP: alertFP, log: logf}
}

// capReached reports whether the hard spend cap is hit — bidding must stop
// (return 204) past it, so the decoy can never overspend even on a win storm.
func (l *Ledger) capReached() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Cap > 0 && l.Spend >= l.Cap
}

// recordWin books a win notice: increment the counter, add the price, alert
// on the FIRST non-zero win, and persist the counter so it survives a crash.
// Returns true if this win crossed the hard cap.
func (l *Ledger) recordWin(price float64) (overCap bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Wins++
	l.Spend += price
	if !l.alerted {
		l.alerted = true
		l.raiseAlert(price)
	}
	l.persist()
	return l.Cap > 0 && l.Spend >= l.Cap
}

func (l *Ledger) raiseAlert(price float64) {
	msg := fmt.Sprintf("FIRST NON-ZERO WIN: price=%.4f spend=%.4f cap=%.4f", price, l.Spend, l.Cap)
	l.log("ALERT %s", msg)
	if l.alertFP != "" {
		_ = os.WriteFile(l.alertFP, []byte(msg+"\n"), 0o644)
	}
}

func (l *Ledger) persist() {
	if l.counterFP == "" {
		return
	}
	_ = os.WriteFile(l.counterFP, []byte(fmt.Sprintf("wins=%d spend=%.6f cap=%.6f\n", l.Wins, l.Spend, l.Cap)), 0o644)
}

func (l *Ledger) snapshot() (wins int, spend float64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Wins, l.Spend
}
