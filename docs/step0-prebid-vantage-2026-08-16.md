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

---

## 2026-08-16 11:00Z — the price distribution on the proven route, and the hosts the wrapper never names

**39 loads, every one with its exit verified from inside the browser as `77.246.104.228`**
(6 passes over the four sparrow sites plus the 15-site sweep above). 79 non-zero sparrow prices
now, still all RUB, still sparrow alone.

### The round values are a set, not an atom, and the split is per placement

| kind | n | values |
|---|---|---|
| round (≤2 decimals) | 12 | 2.5 ×5 · 3.0 ×3 · 4.0 ×3 · 7.5 ×1 — **every one a multiple of 0.5** |
| computed (15–18 decimals) | 67 | **none** lands on a 0.5 grid |

Hand-verified against the raw bodies, each with the price echoed a second time inside the
returned `displayCode`: `"cpm": 4` at `320828`, `"cpm": 3` at `320828`, `"cpm": 7.5` at
`326146`.

Three readings die here, and it is worth naming them because each was mine:

* **"RUB 2.5 is THE atom."** There are four round values, not one.
* **"The atom is that slot's fixed price."** `326146` returned 2.5, 4.0, 4.396624, 4.809109 and
  7.5 across five visits; `319804` returned 2.5 three times and then 6.363761.
* **"A round value is a page-wide fallback."** Atoms and computed values arrive in the SAME
  response set — 74.ru at 07:22:44Z answered 3 round and 5 computed prices in one load, e1.ru at
  10:44:06Z answered 1 and 6. The split is per placement, per moment.

What survives: 4 of the 19 repeatedly-observed placements ever emit a round value, and neither
the site nor the price level predicts which. `319806` (woman.ru) and `326116` (74.ru) sit in the
same 1–7 RUB band as every atom-bearing slot and have never emitted one; `320828` has emitted
nothing else in three visits. **A fixed price and a floor-follower still cannot be separated**:
the Adfox protocol carries no `floorData`, so the floor these numbers might be tracking is not
on the wire at all.

### The hosts the publisher's config never names

`tools/adint-attribute-by-adapter` (new) searches Prebid.js for each unattributed candidate host,
then **fetches the proposed file and requires the host literally in its bytes** before naming
anybody. 17 hosts, ≥5 requests each:

| verdict | host | requests | adapter code | file |
|---|---|---|---|---|
| exact | px.adhigh.net | 422 | getintent | modules/getintentBidAdapter.js |
| exact | kimberlite.io | 365 | kimberlite | modules/kimberliteBidAdapter.js |
| exact | ad.vqserve.com | 57 | viqeo | modules/viqeoBidAdapter.js |
| exact | pb.adriver.ru | 14 | adriver | modules/adriverBidAdapter.js |
| exact | ads.betweendigital.com | 14 | between | modules/betweenBidAdapter.js |
| exact | ssp.hybrid.ai | 14 | vox | modules/voxBidAdapter.js |
| exact | ssp-rtb.sape.ru | 14 | rtbsape | modules/rtbsapeBidAdapter.js |
| domain | hb.bumlam.com | 419 | mediasniper | modules/mediasniperBidAdapter.js |
| domain | yhb.p.otm-r.com | 288 | otm | modules/otmBidAdapter.js |
| none | hb-bidder.skcrtxr.com · ssp.al-adtech.com · sp.ohmy.bid · r.ussp.io · adfox-hb-bidder.rutarget.ru · pbs.alfasense.com · ssp.bidvol.com · widget.sparrow.ru | 29–420 | — | — |

Two cautions the table cannot carry on its own. **`domain` is a weaker claim than `exact`**:
`hb.bumlam.com` is not the `sapi.bumlam.com` mediasniper's adapter publishes, so it is evidence
about the company and none about that endpoint. And **an adapter's `BIDDER_CODE` is a code name,
not a brand** — `ssp.hybrid.ai` cites `vox`, which is Prebid's name for that adapter, while this
study's roster calls the same participant `hybrid`.

Eight hosts stay unattributed, and they stay that way. `r.ussp.io` is the reason the tool
verifies bytes: GitHub's code search answers a query for `ussp` with `anzuSSP` — a different
company that a ranking-trusting tool would have written down as the owner.

---

## 2026-08-16 11:45Z — the census: 84 verified loads, 32 bidders answering, one returning a bid

Three more passes over the 15 sites, `--via direct`, run-idx 20/21/22. The whole schema-5 corpus
is now **84 loads, every one with its exit read from inside the browser as `77.246.104.228`** —
one arm, no unverified rows.

**Question (a): does anybody but sparrow ever return a bid?** No. 32 bidders answered at least
one auction message; exactly one ever returned a bid object (sparrow: 89 priced, 7 zero). This
is the sentence the study turns on and it is now printed by `adint-hb-report` under the schema it
was computed on, because computing it by hand over the whole log pools schema-1 rows and answers
*three* — `hb.bumlam.com`'s `["INCORRECT_UNPUT"]` (a vendor's error token in the bids array) and
a publisher's own `currency.json` (a USD rate of 84.54), both already fixed in the collector and
both back the moment a reader ignores the schema line.

**Question (b): does the 0.5-grid property survive a larger n?** Yes, and the vocabulary did not
grow. 89 sparrow prices at schema 5: **17 round — 2.5 ×3, 3.0 ×6, 4.0 ×7, 7.5 ×1, all multiples
of 0.5** — against **72 computed values, none of which lands on that grid**. Four distinct round
values at n=12 this morning; the same four at n=17.

**And the split is predicted by price level, which I said this morning it was not.** Per
placement, schema 5:

| placement | site | n | max | ever round |
|---|---|---|---|---|
| 326146 | 74.ru | 6 | 7.5000 | yes |
| 326116 | 74.ru | 5 | 6.8241 | yes |
| 326112 | 74.ru | 5 | 4.4352 | yes |
| 320828 | e1.ru | 5 | 4.0000 | yes |
| 326110 | 74.ru | 2 | 3.6099 | yes |
| 320892 | fontanka.ru | 6 | 1.9008 | — |
| …13 more, max 1.90 down to 0.0534 | | | | — |

Clean separation, no overlap: every placement whose price ever reached **3.61 or above** emits
round values; every placement topping out at **1.90 or below** never has. That looks circular —
a round value of 2.5 or more is itself large — so the same cut with round values EXCLUDED: the
round-emitting placements' *computed* prices reach 3.61–6.82, the others' top out at 1.90.
`320828` is the extreme case: five observations, all round, no computed price ever seen.

Two of my own claims die here, both small-n artifacts of this morning: `326116` "has never
emitted a round value" (it now has, 3.0 and 4.0), and "neither the site nor the price level
predicts which placements do" (the level predicts it perfectly across these 18 placements).

What this still cannot decide: whether the round numbers are a rate card, a floor that only binds
on expensive inventory, or a rounding rule above a threshold. The Adfox protocol carries no
`floorData` — the floor, if there is one, is not on the wire. And the whole finding is one
bidder's behaviour: **one curve, not a set of shapes.**
