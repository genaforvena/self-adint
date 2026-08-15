package main

// Minimal OpenRTB 2.5 subset — only the fields Step 1 reads or preserves.
// The receiver never round-trips the whole request; it decodes what it needs
// and, for a matched request, extracts the payload fields the brief §A names.

type BidRequest struct {
	ID     string `json:"id"`
	Imp    []Imp  `json:"imp"`
	App    *App   `json:"app"`
	Device Device `json:"device"`
	User   *User  `json:"user"`
	// Test holds an out-of-band instruction used ONLY by the synthetic
	// harness to exercise the panic path (§ correction 1b). A real exchange
	// never sends it; it is ignored on the wire and never persisted.
	Test string `json:"__panic,omitempty"`
}

type Imp struct {
	ID       string  `json:"id"`
	BidFloor float64 `json:"bidfloor"`
}

type App struct {
	Bundle string `json:"bundle"`
}

type Device struct {
	IFA string `json:"ifa"`
	Geo *Geo   `json:"geo"`
}

type Geo struct {
	// geo.type: 1=GPS/Location Services, 2=IP, 3=User provided (OpenRTB).
	// The brief wants the CLASS of precision, never the coordinates — so
	// lat/lon are deliberately NOT in this struct and never decoded.
	Type int `json:"type"`
}

type User struct {
	Data []Data   `json:"data"`
	Ext  *UserExt `json:"ext"`
}

type Data struct {
	ID      string    `json:"id"`   // data provider id
	Name    string    `json:"name"` // data provider name
	Segment []Segment `json:"segment"`
	Ext     *DataExt  `json:"ext"`
}

type DataExt struct {
	SegTax int `json:"segtax"`
}

type Segment struct {
	ID string `json:"id"`
}

type UserExt struct {
	EIDs    []EID  `json:"eids"`
	Consent string `json:"consent"` // TCF string, if it arrives
}

type EID struct {
	Source string `json:"source"`
	UIDs   []UID  `json:"uids"`
}

type UID struct {
	ID string `json:"id"`
}

// --- outbound ---

type BidResponse struct {
	ID      string    `json:"id"`
	SeatBid []SeatBid `json:"seatbid"`
	Cur     string    `json:"cur"`
}

type SeatBid struct {
	Bid []Bid `json:"bid"`
}

type Bid struct {
	ID    string  `json:"id"`
	ImpID string  `json:"impid"`
	Price float64 `json:"price"`
	AdM   string  `json:"adm"`  // fallback creative — must be servable (§1a)
	CrID  string  `json:"crid"`
	NURL  string  `json:"nurl"` // win notice; exchange calls it if we win
}

// Record is the only thing that ever reaches disk, and only for a request
// whose device.ifa matched the target. Foreign-ifa requests never construct
// one (§ privacy invariant). It carries the brief §A payload, not coordinates.
type Record struct {
	TS      string       `json:"ts"`
	Source  string       `json:"source"`
	ReqID   string       `json:"req_id"`
	Bundle  string       `json:"bundle"`
	GeoType int          `json:"geo_type"`
	Consent bool         `json:"consent_present"`
	Segments []SegGroup  `json:"segments"`
	EIDs    []string     `json:"eid_sources"`
}

type SegGroup struct {
	Provider string   `json:"provider"`
	SegTax   int      `json:"segtax"`
	IDs      []string `json:"ids"`
}
