package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Ledger is the win/spend artifact (§1a). "win rate нулевой, расход нулевой"
// is an assumption about the design, not a property of it: a below-floor bid
// wins whenever no floor is applied. So the win counter is a real artifact —
// asserted by the tests, persisted to disk — not a hope, and the first
// non-zero win raises a loud, out-of-band alert.
type Ledger struct {
	mu            sync.Mutex
	Wins          int
	Spend         float64
	Cap           float64
	ZeroPriceWins int
	counterFP     string
	alertFP       string
	zeroFP        string
	alerted       bool
	log           func(string, ...any)
}

func NewLedger(cap float64, counterFP, alertFP string, logf func(string, ...any)) *Ledger {
	if logf == nil {
		logf = func(string, ...any) {}
	}
	l := &Ledger{Cap: cap, counterFP: counterFP, alertFP: alertFP, log: logf}
	// The zero-price marker lives beside the alert file rather than taking a
	// parameter of its own: it is the same directory and the same kind of
	// out-of-band signal, and a second constructor argument nobody passes is a
	// feature that exists only in the tests.
	if alertFP != "" {
		l.zeroFP = filepath.Join(filepath.Dir(alertFP), "WIN-ZERO-PRICE.txt")
	}
	return l
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
	// The alert latch is ONE SHOT, so which event spends it decides whether the
	// operator ever hears about real money. Spending it on a zero was not
	// hypothetical: driven over the synthetic corpus on 2026-08-30 the first win
	// notice carried the settlement macro UNEXPANDED, parsePrice() correctly read
	// it as 0, and the alert fired anyway — writing the sentence "FIRST NON-ZERO
	// WIN: price=0.0000", which is false on its face, and then latching, so the
	// first actual spend arrived in silence. The alarm must be attached to the
	// event it names.
	if !l.alerted && price > 0 {
		l.alerted = true
		l.raiseAlert(price)
	}
	// A zero-price win is NOT nothing, and it must not be swallowed just because
	// it is not the alert's event. Either the exchange failed to substitute
	// ${AUCTION_PRICE} or it settled at zero; in the first case we have won an
	// impression whose cost we cannot see, which is exactly the state where a
	// spend counter quietly understates. It gets its own loud marker, and it is
	// counted apart from spend so the two can never be netted against each other.
	if price <= 0 {
		l.ZeroPriceWins++
		l.log("WIN WITH NO PRICE: wins=%d zero_price_wins=%d — settlement macro unexpanded, or a zero settlement; spend is a LOWER BOUND", l.Wins, l.ZeroPriceWins)
		if l.zeroFP != "" {
			_ = os.WriteFile(l.zeroFP, []byte(fmt.Sprintf(
				"WIN WITH NO PRICE: wins=%d zero_price_wins=%d spend=%.6f — spend is a LOWER BOUND\n",
				l.Wins, l.ZeroPriceWins, l.Spend)), 0o644)
		}
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
	// zero_price_wins rides in the counter itself: a reader who sees spend and not
	// this number cannot tell a cheap month from a month whose prices we never got.
	_ = os.WriteFile(l.counterFP, []byte(fmt.Sprintf("wins=%d spend=%.6f cap=%.6f zero_price_wins=%d\n",
		l.Wins, l.Spend, l.Cap, l.ZeroPriceWins)), 0o644)
}

func (l *Ledger) snapshot() (wins int, spend float64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Wins, l.Spend
}
