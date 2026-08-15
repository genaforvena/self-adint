package main

import "time"

// Verdict is the decoy bidder's decision for one request.
type Verdict int

const (
	NoBid Verdict = iota // → HTTP 204
	BidBelowFloor
)

// Filter decides whether a non-matching request still passes the geo/bundle
// gate the seat is configured for. In production this gate lives ON THE
// EXCHANGE (brief: "гео-фильтр настраивается на стороне биржи, не в коде");
// here it is a local stand-in so synthetic traffic exercises both branches.
type Filter func(*BidRequest) bool

// Config is the receiver's whole configuration surface.
type Config struct {
	TargetIFA string  // the operator's own device.ifa — the ONLY ifa ever persisted
	Floor     float64 // fallback floor when imp carries none
	SpendCap  float64 // HARD cap enforced on our side (§1a); bidding stops at it
	Source    string  // label written into every Record
	Pass      Filter  // geo/bundle gate for non-matching traffic
}

// decide implements the brief's decoy rule verbatim:
//
//	если ifa не совпал и не прошёл гео/bundle-фильтр → 204
//	иначе                                            → бид ниже флора
//
// It returns the verdict and whether the request matched the target ifa.
// matched is what gates persistence; it is NOT the same question as "do we
// bid" — the decoy bids on filter-passing foreign traffic to keep the seat
// alive, but only a matched request is ever written.
func (c *Config) decide(r *BidRequest) (v Verdict, matched bool) {
	matched = r.Device.IFA != "" && r.Device.IFA == c.TargetIFA
	if !matched && !c.pass(r) {
		return NoBid, false
	}
	return BidBelowFloor, matched
}

func (c *Config) pass(r *BidRequest) bool {
	if c.Pass == nil {
		return false
	}
	return c.Pass(r)
}

// bidPrice returns a price strictly BELOW the imp floor. With no floor the
// brief's decoy still bids; that bid can win outright in a first-price
// auction with no floor (§1a), which is exactly the rare event the ledger
// and spend cap exist to catch — so we do not silently bid 0.
func (c *Config) bidPrice(r *BidRequest) float64 {
	floor := c.Floor
	if len(r.Imp) > 0 && r.Imp[0].BidFloor > 0 {
		floor = r.Imp[0].BidFloor
	}
	if floor <= 0 {
		return 0.01 // no floor to undercut; smallest non-zero bid
	}
	return floor * 0.5
}

// buildBid constructs the decoy response with a SERVABLE fallback creative
// (§1a: winning and failing to serve inflates error rate, a costlier axis
// than the bid rate the decoy protects).
func (c *Config) buildBid(r *BidRequest) BidResponse {
	impID := ""
	if len(r.Imp) > 0 {
		impID = r.Imp[0].ID
	}
	return BidResponse{
		ID:  r.ID,
		Cur: "USD",
		SeatBid: []SeatBid{{Bid: []Bid{{
			ID:    r.ID + "-b",
			ImpID: impID,
			Price: c.bidPrice(r),
			AdM:   fallbackCreative,
			CrID:  "decoy-fallback-1",
			NURL:  "", // wired by the server to its own /win endpoint
		}}}},
	}
}

// fallbackCreative is a minimal always-servable creative. A real seat would
// point this at a hosted 1x1; inline keeps Step 1 self-contained.
const fallbackCreative = `<div style="width:1px;height:1px"></div>`

// extract builds the on-disk Record from a MATCHED request. It is only ever
// called when decide() returned matched=true, so a foreign ifa can never
// reach it. It copies precision CLASS (geo.type) and segment/eid structure,
// never coordinates and never raw uids beyond their source graph.
func (c *Config) extract(r *BidRequest) Record {
	rec := Record{
		TS:     time.Now().UTC().Format(time.RFC3339Nano),
		Source: c.Source,
		ReqID:  r.ID,
	}
	if r.App != nil {
		rec.Bundle = r.App.Bundle
	}
	if r.Device.Geo != nil {
		rec.GeoType = r.Device.Geo.Type
	}
	if r.User != nil {
		for _, d := range r.User.Data {
			g := SegGroup{Provider: providerName(d)}
			if d.Ext != nil {
				g.SegTax = d.Ext.SegTax
			}
			for _, s := range d.Segment {
				g.IDs = append(g.IDs, s.ID)
			}
			rec.Segments = append(rec.Segments, g)
		}
		if r.User.Ext != nil {
			rec.Consent = r.User.Ext.Consent != ""
			for _, e := range r.User.Ext.EIDs {
				rec.EIDs = append(rec.EIDs, e.Source)
			}
		}
	}
	return rec
}

func providerName(d Data) string {
	if d.Name != "" {
		return d.Name
	}
	return d.ID
}
