# Step 1 — receiver + decoy bidder (synthetic)

The local half of the project (`PLAN.md` Step 1 / brief §"Порядок работ" 1). No account, no
exchange, no network egress: an HTTP endpoint that receives OpenRTB bid requests, decides the
decoy's bid, and persists **only** the target device's payload. Runs on synthetic requests.

## Run

    ADINT_TARGET_IFA=<your device.ifa> \
    ADINT_BUNDLE_ALLOW=com.pass.a,com.pass.b \
    ADINT_DATA_DIR=data \
    ADINT_ADDR=127.0.0.1:8788 \
    go run .

`POST /bid` (OpenRTB BidRequest JSON) · `GET /win?p=<price>` (win notice) · `GET /healthz`.

Env: `ADINT_TARGET_IFA` (required — the only ifa ever written), `ADINT_FLOOR` (fallback floor),
`ADINT_SPEND_CAP` (hard cap, default $1.00), `ADINT_BUNDLE_ALLOW` (local stand-in for the
exchange-side geo/bundle gate), `ADINT_SOURCE`, `ADINT_DATA_DIR`, `ADINT_ADDR`.

## The decoy rule (brief, verbatim)

    если ifa не совпал и не прошёл гео/bundle-фильтр → 204
    иначе                                            → бид ниже флора

Bidding on filter-passing foreign traffic keeps the seat alive (a 204-on-everything seat throttles
in days). **Persistence is a different question from bidding:** only a request whose `device.ifa`
equals the target is ever written to disk — the decoy bids on foreign traffic but never records it.

## What reaches disk

`data/receiver.jsonl`, one matched request per line: `bundle`, `geo_type` (precision *class*, never
coordinates), `consent_present`, `segments[]` (provider/segtax/ids), `eid_sources[]`. Coordinates,
raw uids, and any foreign ifa are structurally absent — see the tests.

`data/win-counter.txt` (wins + spend, persisted every win) and `data/WIN-ALERT.txt` (written on the
first non-zero win — see §1a below).

## Gates (`go test`) — each seen red under a real mutation

| gate | what it asserts | seen red by |
|---|---|---|
| `TestForeignIFANeverWritten` | privacy invariant: a foreign ifa never reaches disk, even when bid on | `matched = true` → 3 rows persist, not 1 |
| `TestPanicPathNoLeak` | §1b: the panic handler leaks neither body, uid, nor a PII-bearing req id | logging `req_id` (which embedded the ifa) → red on first write |
| `TestExtractNoCoordinates` | precision class kept, lat/lon dropped | — (structural: `Geo` has no lat/lon field) |
| `TestWinCounterAndAlert` | §1a: win counter is an asserted artifact; first win alerts | — |
| `TestSpendCapHalts` | §1a: past the hard cap the decoy stops bidding (204) | — |
| `TestDecoyRule` / `TestBidBelowFloor` | the truth table above; the bid is strictly below floor | — |

The privacy invariant is a **construction, not a discipline** (`CLAUDE.md` rule 2): the ifa filter
sits in the parser upstream of the sink, the sink's only caller is the matched path, and the panic
handler logs nothing request-derived. "win rate нулевой" is not assumed — the win counter is real,
because a below-floor bid wins whenever no floor is applied (§1a).

## Not done here

The exchange-side geo/bundle gate is faked locally (`ADINT_BUNDLE_ALLOW`); on a real seat it lives
on the exchange. No TLS/whitelist/QPS — those belong to Step 3 (seat onboarding), which is the
operator's call alone.
