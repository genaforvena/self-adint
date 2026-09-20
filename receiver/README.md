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

The canonical Step 1 process harness is the repository-root entry point. It writes only the
explicitly selected private run directory; no committed `step1-2026-09-07` fixture is a status source.

    tools/adint-step1-run --seed 20260830 --out-dir data/step1-20260830 --offline-egress

The legacy `tools/adint-receiver-run` name delegates to that same command and defaults to the same
private output directory.

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
| `TestExpiredTmaxRefusesWork` | §1e: expired `tmax` returns 204 and writes no row | `tmax` budget check removed → red |
| `TestResponseCompletionGateRefusesAfterIngressBudget` | a late response is refused by the same private-ingress budget while the matched observation remains recorded | final gate removed → red |
| `TestDecoyRule` / `TestBidBelowFloor` | the truth table above; the bid is strictly below floor | — |

The privacy invariant is a **construction, not a discipline** (`CLAUDE.md` rule 2): the ifa filter
sits in the parser upstream of the sink, the sink's only caller is the matched path, and the panic
handler logs nothing request-derived. "win rate нулевой" is not assumed — the win counter is real,
because a below-floor bid wins whenever no floor is applied (§1a).

The receiver reads each request's `tmax` in milliseconds, subtracts measured arrival-to-decision
time, and refuses work when the remaining budget is zero or negative before persistence. A second
check immediately before response encoding refuses a bid that would complete after that same
private-ingress budget; a matched observation already written remains an observation even when
the late response is refused. The current process harness does not claim to exercise an expired
request; the focused contract checks are `TestExpiredTmaxRefusesWork` and
`TestResponseCompletionGateRefusesAfterIngressBudget`. `at` is carried by some synthetic requests
but is not an enforced receiver field.

## Not done here

The exchange-side geo/bundle gate is faked locally (`ADINT_BUNDLE_ALLOW`); on a real seat it lives
on the exchange. No TLS/whitelist/QPS — those belong to Step 3 (seat onboarding), which is the
operator's call alone.
