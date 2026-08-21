# Which sites each vantage cannot see, and what that costs the paired design

*2026-08-21. Measured from the published corpus in `public-data/` — 13 cells, 338 pairs,
two arms. Regenerate every figure here with `python3 tools/adint-status`; the per-site
tables below are derived from the same files and are reproduced in §5.*

## 0. Why this document exists

The study's claim is a **contrast**: the same page, at the same minute, loaded from a
Russian mobile vantage and from a foreign one, with the two bidder rosters compared. A pair
where one arm never saw the page carries no contrast at all. It is not a weak data point —
it is not a data point.

Until 2026-08-21 the progress figure quoted in every commit and handoff was counted in
**loads attempted**, which is not the study's unit. Recounted in readable pairs it is
roughly half (`55cc09c`). This document is the other half of that correction: not *how many*
pairs are unusable, but **which sites lose them, on which arm, and why** — because the loss
is not spread evenly and it is not random.

## 1. The headline

Across the whole published corpus:

| | NL/foreign arm | RU mobile arm |
|---|---:|---:|
| loads answered `4xx` by the publisher | **50** | **0** |

Fifty refusals on one arm and none on the other. Every `401`/`403` in the corpus was served
to the foreign vantage; the domestic vantage has never once been declined by a publisher in
this frame.

That asymmetry is the reason the two arms fail in **different shapes**, and the two shapes
need different words. The capture already keeps them apart and this document uses its
vocabulary:

- **declined** — the publisher's server answered, with a `4xx`. It saw us and said no. This
  is a fact *about the publisher's policy toward that vantage*.
- **no response** — the navigation never returned anything. This is a fact *about our own
  link*, and it says nothing whatever about the site.

Conflating them would let our own blind spots read as publisher behaviour, which is exactly
the mistake this study cannot afford to make: it is a study *of* differential treatment by
vantage, so an instrument that treats vantages differently is indistinguishable from the
result, unless the two are held apart by construction.

## 2. What the foreign arm cannot see

| site | seen / loads | how it fails |
|---|---:|---|
| `rbc.ru` | **0 / 24** | `401` × 24 |
| `iz.ru` | 10 / 23 | `403` × 13 |
| `gismeteo.ru` | 15 / 25 | `403` × 10 |
| `hh.ru` | 16 / 22 | `403` × 3, no response × 3 |

`rbc.ru` is the extreme case and deserves naming plainly: **the foreign arm has never once
seen it.** Twenty-four attempts, twenty-four `401`s, across every cell and every daypart. It
answers in about two seconds, so this is not a timeout dressed as a refusal — it is a
decision, made quickly, every time.

`hh.ru` refusing a foreign vantage is independently known on this mesh from unrelated work,
which is mild corroboration that the mechanism is the publisher and not our exit.

## 3. What the RU mobile arm cannot see

| site | seen / loads | how it fails |
|---|---:|---|
| `aif.ru` | 5 / 23 | no response × 18 |
| `rbc.ru` | 5 / 23 | no response × 18 |
| `magnit.ru` | 8 / 22 | no response × 14 |
| `matchtv.ru` | 8 / 22 | no response × 14 |
| `pikabu.ru` | 11 / 25 | no response × 14 |

Zero of these are refusals. They are navigations that returned nothing inside a 120 s
timeout, on a mobile tether that is the node's only uplink.

**This is our blindness, and there is a direct measurement to that effect.** On 2026-08-21
at 02:1x UTC, from inside the vantage's own network namespace, `aif.ru` answered `curl` with
`200` in **1.8 s** while the browser arm was timing out on it. A `curl` of the HTML is not a
browser load of a full page with its subresources, so the two are not in contradiction: the
site is reachable and the *page* does not finish over that link. That is page weight against
link capacity, not a publisher decision.

The honest limit: this was spot-checked for `aif.ru` only. The other four are the same
failure *class* on the same link, but each has not been individually shown to be reachable
while unloadable. Treat "our link" as well-supported for the class and demonstrated for one
member.

A second measurement from the same hour is a caution rather than a result. The vantage was
briefly **walled** — in-country destinations answering normally while every foreign host
tested was dead, including `securepubads.g.doubleclick.net` (`713f038`). No cell was running
in that window, so no published row is affected. But a vantage can enter a state where the
RU arm would see a systematically thinner roster for reasons that have nothing to do with
the Russian ad market, and nothing in a row's own fields would show it. `adint-ru-vantage-keep`
now records `foreign=BLACKHOLED` on every 5-minute tick so the state is at least on the record
when it recurs.

## 4. What it costs, per site

Readable-pair yield — pairs where **both** arms saw the page:

| site | pairs | readable | yield |
|---|---:|---:|---:|
| `rbc.ru` | 24 | 0 | **0 %** |
| `aif.ru` | 23 | 4 | 17 % |
| `iz.ru` | 24 | 7 | 29 % |
| `magnit.ru` | 26 | 8 | 31 % |
| `matchtv.ru` | 25 | 8 | 32 % |
| `pikabu.ru` | 26 | 11 | 42 % |
| `mk.ru` | 25 | 13 | 52 % |
| `gismeteo.ru` | 25 | 13 | 52 % |
| `kp.ru` | 24 | 13 | 54 % |
| `hh.ru` | 24 | 14 | 58 % |
| `aviasales.ru` | 23 | 17 | 74 % |
| `sports.ru` | 24 | 18 | 75 % |
| `interfax.ru` | 22 | 17 | 77 % |
| `ria.ru` | 23 | 18 | 78 % |
| **total** | **338** | **161** | **48 %** |

The frame is not delivering 14 sites. It is delivering four sites at ~75 %, five in the
middle, four badly, and **one not at all**.

Two consequences follow, and they are different in kind.

**Budget.** `rbc.ru` has consumed 24 pairs of capture time and returned nothing, and by
construction it cannot return anything while the foreign arm is refused on every attempt.
Every cell spends two loads per replicate on it for a guaranteed empty result.

**Validity, and this is the serious one.** The yield column is not noise — it is a
*site-dependent* filter, and the sites it filters hardest are not a random subset. Any
statement of the form "RU publishers show bidder pattern X" computed over readable pairs is
computed over a corpus that has quietly dropped `rbc.ru` entirely, most of `aif.ru`, and
two-thirds of `magnit.ru` and `matchtv.ru`. If wrapper configuration correlates with
anything that also predicts page weight — publisher size, ad density, video — then the
filter is correlated with the outcome and the surviving corpus is biased, not merely
smaller. Nothing here establishes that such a correlation exists. Nothing here rules it out
either, and a shrunk corpus that *looks* clean is the more dangerous of the two.

## 5. Reproducing this

Every table above comes from `public-data/hb-paired-*.jsonl`, `kind == "load"` rows, with
one row per `(cell, arm)` — a published cell exists in both `data/` and `public-data/`, and
reading both counts it twice. That trap is live: the first draft of §3's per-site counts was
exactly 2× the truth for this reason, which is the same double-count `tools/adint-status`
already dedups against and the same one that inflated the coverage percentage before
`55cc09c`.

- `seen` = `http_status` in `[200, 400)`
- `declined` = `http_status` in `[400, 500)`
- `no response` = no `http_status` at all
- `readable pair` = a `pair_id` whose row in **both** arms is `seen`

`tools/adint-status` prints the per-cell version of the last of these as its bolded
`readable` column, and gates it in `--test`.

## 6. Open — these are decisions about the study, not fixes to a tool

1. **`rbc.ru`.** Drop it from the frame, or keep it and record it as a permanent hole?
   Keeping it costs two loads per replicate forever for a structurally guaranteed zero.
   Dropping it removes a top-rank RU publisher from a frame that claims to be rank-ordered,
   which is itself a bias — and one that would be invisible in the results. **Whichever way
   it goes, it must be written down in `ref/CANONICAL-FRAME` where the frame is declared,
   not decided here.** Note the trap in the tempting option: a frame that quietly drops the
   sites the instrument cannot see becomes a frame *defined by the instrument*, and reports
   full coverage of itself.
2. **Reporting yield beside every roster claim.** A per-site roster statement should carry
   the number of readable pairs it rests on. `ria.ru` at 18 and `aif.ru` at 4 currently read
   identically in any aggregate.
3. **Whether the RU arm's timeouts can be attributed per site.** A cheap control exists:
   at capture time, `curl` the same URL from inside the vantage's namespace and record the
   result beside the load. That separates "this link could not finish this page" from "this
   host was unreachable" per row, instead of by hand afterwards for one site.
