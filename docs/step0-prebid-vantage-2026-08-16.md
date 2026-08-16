# Step 0 — prebid-bidder-strategy-study: what the sweep found, and what the vantage forbids

**Question of the study** (operator, board `[design]` 2026-08-16T02:49:12Z): classify DSPs by
the SHAPE of the marginal bid distribution at fixed input, framed by Guerre–Perrigne–Vuong.
**Question of step 0**: which runet sites actually run a client-side auction we can read.
**Artifact**: `data/pbjs-probe.jsonl` (append-only, every row carries its schema and the egress
IP). **Read it with** `tools/adint-stage0-report`, never off a console tail.

---

## The corpus limitation, stated first because it conditions every number below

**Correction, same night (05:35Z): the Montréal exit was a PROXY's, not this node's.** This
shell exports `HTTP_PROXY=HTTPS_PROXY=ALL_PROXY=http://127.0.0.1:8118` — privoxy, forwarding to
socks5 on 1080 — and both curl and Chromium honour it. So the vantage named below was measured
through that proxy and every capture took the same path:

| path | exit | where |
|---|---|---|
| proxied (the default here) | `38.49.216.141` | Montréal, CA, AS26832 Rica Web Services |
| direct (the node's own route) | `77.246.104.228` | Amsterdam, NL, AS216071 Servers Tech Fzco |

The node card said NL and the node is NL; I "corrected" it to Montréal off the proxied
measurement and had to put it back. What follows was true of the path the sweeps took, and the
qualitative conclusion survives — both exits are `hosting: true` datacenter ASNs outside the geo
being measured — but the country, the ASN and every `ERR_TUNNEL_CONNECTION_FAILED` belong to
privoxy, not to this node.

This node's sweeps egressed at **Montréal, Canada — 38.49.216.141, AS26832 Rica Web Services,
`hosting: true`** (ip-api, 2026-08-16), through the proxy above. Not the geo being measured.

Both halves matter and they compound:

- **Wrong geo.** Runet publishers sell runet audiences (and the direct path, Amsterdam, is no
  closer to being runet than Montréal). Demand for a Canadian impression is
  thin to absent regardless of how well the page runs.
- **Datacenter ASN.** `hosting: true` is what invalid-traffic filters exist to drop. An SSP that
  answers a residential IP will decline a hosting one, and it declines it *silently* — as a
  no-bid, which is indistinguishable from "this bidder had no demand" in exactly the column the
  study depends on.

So the central number of this sweep — **zero observed auctions** — is a statement about the
vantage, not about the publishers. It must never be written down as "runet sites do not run
prebid". Eight sites additionally answered 401/403/507/521 and one errored; those rows say
nothing at all about prebid and are excluded from the denominator by construction.

A tunnel on this node is **not** the fix: the router already carries the VPN, so a host-side
tunnel nests inside it and loops the egress (the mesh's nesting invariant; it has already cost
one hand-rescued outage). The fix is a different *body*, not a different route.

## What the sweep counted (schema 4, 31 sites, one visit each)

| verdict | n | meaning |
|---|---|---|
| `none` | 19 | page rendered, no prebid global in any frame |
| `pbjs-idle` | 3 | a global exists, no auction ran in our window |
| `auction` | 0 | bidders actually asked |
| `blocked` | 8 | 4xx/5xx — the page did not render |
| `error` | 1 | no usable response |

**Observable denominator = 22.** `blocked` and `error` are excluded: dividing by a 403
manufactures an absence rate, and the first version of this tool did exactly that — it folded
4xx/5xx into `none` and wrote nine geo-blocked sites down as publishers without prebid.

## What it found instead: the wrapper is the market

| marker | sites (of 22 observable) |
|---|---|
| `Ya` / `yaContextCb` | 22 |
| `Ya.adfoxCode` | 21 |
| `YaHeaderBiddingSettings` | 16 |

Prebid's own global appears on 3, and is empty every time. **`prebid.js` is not the vantage of
this market** — the operator's pre-named branch ("if it is mostly Adfox HB with its own wrapper,
take the wrapper apart first") is confirmed, and it is confirmed by a live read of the page,
not by grepping markup.

## The reading that survives the vantage: the roster

`YaHeaderBiddingSettings` is configuration the page *ships*. It does not depend on anyone
deciding to bid, so it arrives intact at a Canadian datacenter IP. 16 rosters, **41 distinct
bidder identities**. Reach across sites:

```
mytarget 13 · betweendigital 12 · mediasniper 12 · adriver 11 · astralab 11 · buzzoola 11
hybrid 10 · sape 10 · roxot 10 · gnezdo 8 · smi2 8 · getintent 6 · otm 6 · videonow 6
bidvol 5 · ohmybid 5 · alfasense 4 · solta 4 · sparrow 4 · … tail to 1
```

Two cautions that belong with the numbers, not below them:

- **`adfox_*` entries are Adfox's own adapters**, not third-party DSPs (`adfox_adsmart` 11,
  `adfox_imho-video` 5, `adfox_yandex_*`). Counting them alongside independent bidders would
  overstate the field's breadth.
- **The wrapper's timeout is per publisher and it varies**: 3000ms (lenta, gazeta, championat),
  1500 (kp), 1300 (ura), 1200 (3dnews, sport-express), 1000 (seven sites), 900 (mk), 500 (zr). The design already
  requires `timeout` to be a column separate from `no-bid`; this says the *cutoff itself is set
  by the publisher*, so selectivity cannot be compared across sites without conditioning on it.

## A correction to the design: one visit is a sample, not a site property

Sites disagreed with themselves across identical visits:

- `aif.ru` — `pbjs` → `pbjs-idle` → `none` over three visits.
- `fontanka.ru`, `iz.ru` — answered 200 on early runs, 403 by the third: we are being
  fingerprinted from run to run.

Any step-0 corpus therefore needs a **repeat column**; a single-visit verdict feeds stratification
on noise. The reporter states stability as UNMEASURED when no site was visited twice, rather than
letting silence read as stable.

## Three instrument defects found and closed (each gate seen RED against a mutant)

1. **The read was one frame wide, the verdict was site wide** (`030474b`). `page.evaluate` runs
   in the main frame only, so 19 `none` verdicts were claims about the top document printed as
   claims about the publisher — and wrappers routinely run the auction inside an ad iframe. Now
   every frame is read, with `auction-in-frame` / `pbjs-in-frame` kept apart from `auction`
   because the collector that reads the page's global cannot reach a frame. Fixture is
   cross-origin on purpose (`localhost` iframe in a `127.0.0.1` page).
2. **One bidder, two spellings.** `biddersMap` renders `betweenDigital`; a bid renders
   `betweendigital`. A join on the raw string splits one bidder into two. Canonical union is
   published beside the raw forms, never instead of them.
3. **A roster fallback that dropped bidders and looked right doing it.** The report took the
   stored `bidders` field when present and fell back to `biddersMap` alone when absent — silently
   losing every bidder that appears only in the adUnits. `ohmybid` is asked on 6 of 7 slots on
   3dnews.ru and was missing from the count; `roxot` moved 9 → 10, `getintent` from unlisted to
   6. Nothing in the output looked broken, which is the whole hazard. The union is recomputed
   every time and the stored field is *checked* against it, never preferred to it.

## Addendum, same night: the auction DOES run here — and every bidder declines

The collector (`tools/adint-hb-capture`) reads the wire instead of the JS API, and on
lenta.ru it finds the auction the pbjs probe could not see:

```
POST hb.bumlam.com/yandex/            {"places":[{"id":"begun_block_…","placementId":"41139",
POST ssp.al-adtech.com/api/adfox/bids   "sizes":[[970,250]],"codeType":"combo",…}]}  ->  {"bids":[]}
```

Three things follow, and they sharpen the headline above rather than contradict it:

- **The market's protocol is not OpenRTB.** No `imp`, no `bidfloor`, no `seatbid`. A collector
  built to the expected schema recorded 197 rows and zero candidates while two bidders were
  answering the whole time.
- **`{"bids":[]}` is a real no-bid** — a bidder that was asked, answered, and declined. It is
  neither silence nor an absent auction, and the three must never be folded.
- **Attribution is exact here.** `placementId → bidder` comes from the publisher's own wrapper
  config and each endpoint is one bidder: 41139 → mediasniper, 69eb5feb… → astralab, both
  confirmed live. No host list, no substring matching.

So the vantage's effect is now visible in its actual shape: the wrapper fires, the bidders
answer, and the answer is always no. Everything the design needs except the price column —
who is asked, when, with what floor, who answers, how fast, who stays silent — is measurable
from this node today. Only the prices need the right body.

## What this forces

Guerre–Perrigne–Vuong estimates a bid distribution. From this vantage there are **no bids** —
not few, none — so there is nothing to estimate, and no amount of further sweeping from here
changes that.

The mechanism to collect them is nonetheless clear and geo-independent: on an Adfox HB page the
prices do not live in a JS API, they live in the **network** — the wrapper issues one request per
bidder and takes back a response, and the design's required columns (`originalCpm` /
`originalCurrency`, floor, timeout distinct from no-bid) are present there. The collector can be
built and gated against a local fixture from here. Only the *answers* need the right body.

**Open question for the operator** (not a task anyone here can close): a RU-resident, residential
vantage — his phone on LTE is the mesh's only candidate. Until then step 0 yields the roster and
the wrapper's structure, which is real and is enough to design the collector against, and it does
not yield a single bid.

---

## Correction, 2026-08-16 10:30Z: the "direct" arm was never direct, and the silence it measured was privoxy's

The section above puts the proxy in the right place, but it left a claim standing that was not
true: that a browser launched with `--no-proxy-server` takes the node's own route. It does not.
Measured on this node the same afternoon, three launches of the same chromium:

| launch | exit read from INSIDE the browser | `ads.betweendigital.com` |
|---|---|---|
| default (inherits the shell's proxy) | `38.49.216.141` | `ERR_TUNNEL_CONNECTION_FAILED` |
| `--no-proxy-server` only | `38.49.216.141` | `ERR_TUNNEL_CONNECTION_FAILED` |
| proxy vars stripped from the child's env | `77.246.104.228` | HTTP 400 (a bare GET, answered) |

The inherited `HTTP_PROXY`/`HTTPS_PROXY`/`ALL_PROXY` beats the flag. So the `--treatment proxy`
sweep — 117 loads labelled `proxy`, 117 labelled `direct` — ran both arms through privoxy. It is
an **A/A null wearing a treatment's name**, and the difference it showed (29 priced rows on one
side, 4 on the other) is the corpus disagreeing with itself, not an effect of vantage.

The consequence for the bidder roster is larger than the arm. `ERR_TUNNEL_CONNECTION_FAILED` is
an error only a proxy CONNECT can produce: it says privoxy refused the host, and privoxy ships an
ad-domain blocklist. Every "this bidder is 100% error" line in this study's earlier notes was a
statement about that blocklist.

**What the four blocked bidders actually do, asked from the node's own route** (15 sites,
`--via direct`, 15/15 loads with the exit verified from inside the browser as `77.246.104.228`):

| bidder | through privoxy | direct, verified |
|---|---|---|
| betweendigital | 562 × `error` | 26 no-bid · 10 redirect (cookie-sync) · 7 no-answer · **0 error** |
| MyTarget/VK | 131 × `error` | 13 no-bid · 5 no-answer · 1 unparsed · **0 error** |
| videonow | 24 × `error` | 1 no-bid · 2 no-answer |
| segmento | 36 × `error` | 2 no-answer · 1 error |

Hand-verified against the raw bodies: betweendigital's `no-bid` is HTTP 200 with an empty `bids`
array (a real decline), its `redirect` is a 302 cookie-sync hop (`/match?bidder_id=45632`) that
is attribution-only and never an auction answer, and its `no-answer` is `ERR_ABORTED` past the
wrapper's deadline with `net_fault: false` — honest silence, not our network. 31 distinct bidders
answered at least one auction message in these 15 loads.

**What did not change.** Still exactly one bidder returns a price: sparrow, 14 more non-zero
prices here (47 total), all RUB. Everyone else declines. The study's aim — the SHAPE of a
marginal bid distribution *per DSP* — remains out of reach from any vantage this node has: one
curve is not a set of shapes.

**What the new lane corrects in the price story.** The earlier note that sparrow "returns the same
computed value for one placement hours apart, so the price is a function of the slot" does not
survive three visits: placement `326146` (74.ru) returned `4.809109`, then `4.0`, then
`4.396624`; `320892` (fontanka.ru) returned `0.573035`, `1.146862` ×3, `1.029160`. Repeats happen
and are exact, but a placement's price MOVES. The round atoms are still there and still atoms —
`2.5` at `319804` across four loads, `4.0` at `320828` (raw body `"cpm": 4`, repeated as
`'cpm': 4` inside the returned `displayCode`) — but they appear at placements that compute a long
decimal at other moments, so "the atom is that slot's fixed price" is refuted too. A fixed price
and a floor-follower still cannot be separated here: the Adfox protocol carries no `floorData`.

**How the tooling now prevents the repeat.** The arm is measured where it exists — from inside
the browser, per load, before the site is opened (`browser_egress_ip` + `vantage_verified` on the
load and on every request row, schema 5). A sweep whose two arms come back with the same exit
refuses to start. The old gate that looked like coverage compared `egress_ip` against
`browser_used_proxy` — two fields the same code writes from one argument — and every other leg
runs against `127.0.0.1`, which chromium exempts from any proxy, so a localhost fixture cannot
tell a bypassed proxy from an obeyed one. Commits `f0d53af`, `5915151`, `945836d`.
