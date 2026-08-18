# step0c — sampling frame and schedule for the RU header-bidding study

date: 2026-08-18 · decided by: the `adint` window · status: **frame and schedule fixed; capture not yet
re-run under them.** Everything below that is a count was measured on this node on this date and carries
its `n`. Nothing here is a result about the Russian market — the study that would produce one is what
this document schedules.

## 0. Why this document exists, and what it is allowed to claim

The lane turned toward header bidding because the western literature on it covers US/EU inventory and
does not contain the Russian SSPs at all: `Yandex`, `adriver`, `buzzoola`, `betweendigital`, `sape`,
`smi2`, `MyTarget/VK`, `astralab`, `gnezdo`, `hybrid`, `bidvol`, `alfasense` are visible by name in our
own capture and absent from the papers. That is a real gap and it is worth filling properly.

**The size of what we hold today, stated before anything else, because the presentation must never
outrun the data:**

| dataset | loads | sites | days | hours (UTC) | candidate rows | rows with a price |
|---|---:|---:|---:|---|---:|---:|
| `hb-auction.jsonl` (16.08) | 418 | 31 | 1 | 4–7, 10–11 | 43408 | — |
| `hb-auction-ru.jsonl` (18.08, RU vantage) | **57** | 15 | **1** | **15–16** | 5259 | **13** |
| `hb-auction-nlctl.jsonl` (18.08, NL control) | 14 | **1** | 1 | 16–17 | 1155 | — |

Thirteen observable prices, from one day, two hours of the clock, on fifteen sites. The word *market*
is not available to us yet, and no chart in this repository may be captioned as though it were. Every
figure we publish prints its `n` and its coverage in the same line as the number.

## 1. The frame

### 1.1 What we have now is a convenience sample, not a frame

The 31 sites captured so far were assembled by hand during step 0. No selection rule was ever written
down, which means the set cannot be audited, cannot be reproduced by a stranger, and cannot support a
statement of the form "on the Russian market, …" — only "on these 31 sites, …". Their Tranco ranks
range from **#1146** (`pikabu.ru`) to **#56784** (`eva.ru`), i.e. across a factor of ~50 in popularity,
which is the signature of an ad-hoc list rather than a stratum.

### 1.2 The universe: Tranco, dated and pinned

The ranked universe is the Tranco list, because it is built for exactly this — a reproducible, dated,
citable sampling frame that a third party can re-fetch and diff.

```
list_id      V3JQN
created_on   2026-08-01T22:00:02Z
window       2026-07-03 .. 2026-08-01 (30 days), combination: dowdall
providers    crux, farsight, majestic, radar, umbrella
download     https://tranco-list.eu/download/V3JQN/1000000
sha256       c68e363d5cc412a544024d947826db636fe84e56e16d95a3ae32f82684dcdd2a
```

Derived and committed: **`ref/ru-universe-tranco-V3JQN.tsv`** — every registrable domain in the top
50 000 whose TLD is `.ru`, `.su` or `.xn--p1ai`. **n = 1944.** Nothing hand-picked, nothing removed.

### 1.3 The naive rule fails, and its failure is a finding

The obvious rule — *take the top N `.ru` domains* — was executed against the universe before being
adopted, and it does not select the object of study. Of the **top 60** `.ru` domains, exactly **2**
(`pikabu.ru` #1146, `rbc.ru` #1163) are ad-supported content sites a person reads. The other 58 are:

- CDNs and image hosts — `okcdn.ru`, `yccdn.ru`, `edgecdn.ru`, `imgsmail.ru`, `cdnvideo.ru`, `cdn-vk.ru`,
  `vkuserphoto.ru`, `wbbasket.ru`, `geobasket.ru`
- ad-tech infrastructure itself — `adriver.ru`, `rtbcdn.ru`, `stbid.ru`, `ad-tech.ru`, `mhverifier.ru`
- registrars and hosting — `nic.ru`, `reg.ru`, `selectel.ru`, `timeweb.ru`, `netangels.ru`, `tld-servers.ru`
- telcos, banks, marketplaces, government, SaaS — `mts.ru`, `megafon.ru`, `rt.ru`, `sberbank.ru`,
  `ozon.ru`, `wildberries.ru`, `avito.ru`, `gosuslugi.ru`, `bitrix24.ru`, `kontur.ru`, `consultant.ru`

Tranco ranks by **resolver traffic**, so it answers *what is resolved*, not *what a human reads with
ads on it*. A top-N cut of it is a frame for infrastructure. This is not a defect in Tranco; it is a
defect in the naive rule, and it is why the rule below has two stages.

### 1.4 The rule, in two stages

> **Stage A — universe.** Every registrable domain at Tranco rank ≤ 50 000 with TLD `.ru` / `.su` /
> `.xn--p1ai`. (n = 1944, committed as `ref/ru-universe-tranco-V3JQN.tsv`.)
>
> **Stage B — inclusion probe.** Walk stage A in rank order and admit a domain to the frame iff a single
> automated probe of its apex URL finds **both**: (i) HTTP 200 with an HTML document, and (ii) an ad
> wrapper present in the page — `Ya` / `YaHeaderBiddingSettings` (Adfox) or a non-empty `pbjs`. Stop at
> 20 admitted domains **per category**. Every domain the probe touches is written to the ledger with its
> verdict, including the ones it rejects and why.

Stage B is mechanical, so a stranger re-runs it and gets our frame or a dated diff. It also *is* the
answer to "name the sites without header bidding separately, their absence is a result too": the
rejection ledger is that list, and it is published, not discarded.

**The rejection reasons must stay distinct — they are different facts.** Our existing capture already
shows all three, and collapsing them would be the study's first lie:

| reason | meaning | seen so far (n loads attempted) |
|---|---|---|
| `no-wrapper` | page loaded, no Adfox and no pbjs | `vedomosti.ru` on 5 of 9 loads with HTTP 200 |
| `blocked` | HTTP 401/403/5xx — **we never saw the page**, so we know nothing about its wrapper | `rbc.ru` 401×10, `rg.ru` 401×9, `quote.rbc.ru` 401×9, `drive2.ru` 403×9, `iz.ru` 403×9, `irecommend.ru` 521×9, `otzovik.com` 507×9 |
| `rate-limited` | HTTP 428/429 under our own load | `woman.ru` 428×10 |

A `blocked` domain is **not** evidence of absent header bidding. Reporting it as such would convert our
own access failure into a market fact. It stays in the ledger as blindness, named.

### 1.5 Three known holes in the rule, each with its declared handling

The TLD test in stage A drops sites that are unambiguously part of the object of study. Measured
against the 31 domains already captured: **24 are inside the universe, 7 are not**, for three different
reasons that must not share one bucket.

1. **RU-audience site on a non-`.ru` TLD** — `habr.com` (#3561), `championat.com` (#2751), `ixbt.com`
   (#5132), `otzovik.com` (#5732), `ura.news`. A TLD is a weak proxy for audience.
   *Handling:* a second, separately declared allowlist file, `ref/ru-audience-nonru-tld.tsv`, each entry
   carrying the evidence for its inclusion (a citable audience/geo source), reviewed as its own artifact.
   Results are reported **both with and without** this allowlist, so a reader can see exactly how much
   of any finding rests on our judgement rather than on the mechanical rule.
2. **Below the depth cut** — `eva.ru` at #56784 falls outside the top-50k, not outside the TLD rule.
   *Handling:* the cut is a declared parameter, not a fact. Sites already captured that fall below it are
   kept and **flagged `below-cut`** in every table, never silently mixed into the frame.
3. **Subdomain** — `quote.rbc.ru`. Tranco is a registrable-domain list, so subdomains are structurally
   unrankable.
   *Handling:* subdomains are out of the frame by construction. If one is captured it is reported under
   its parent with the subdomain named.

### 1.6 Categories

Stratification is by category because a market-wide average over news, автомобили and women's interest
is not a quantity anyone wants. The category source must itself be a citable artifact, not our opinion.
`liveinternet.ru/rating/ru/` is the RU-native candidate and is reachable from this node (HTTP 200,
0.84 s, measured 2026-08-18), but its category index is not in the served HTML — it renders client-side,
so extracting it is a scrape with its own failure modes. **Open, and deliberately not guessed:** either
render that page once and commit the resulting mapping as a dated artifact, or assign categories by hand
in a committed file with the assigner named. Until one is chosen and committed, stage B runs
**unstratified** in pure rank order, and every table says `unstratified` where a category would go.

## 2. The schedule

### 2.1 What one observation costs (measured, not assumed)

| arm | `window_s` | median `observed_s` | median wall gap load→load | n |
|---|---:|---:|---:|---:|
| NL 16.08 | 8 | 23.7 s | 22 s | 414 |
| RU 18.08 | 14 / 45 | 52.7 s | 94 s | 45 |
| NL control 18.08 | 45 | 96.1 s | 96 s | 13 |

**A confound already sits in that table and must be closed before any comparison is run:** the three
datasets used *different observation windows* (8 s / 14 s / 45 s). The window is what separates
`no-answer` from a real no-bid, so a difference in window mechanically produces a difference in outcome
mix. **All arms of the study run at a single fixed `window_s = 45`**, which is the longest we have
evidence for and exceeds the wrapper deadlines we have observed; any row whose window was shorter than
the wrapper's own timeout keeps carrying `window_shorter_than_wrapper`.

At `window_s = 45` one load costs **≈ 96 s of wall clock** (measured, n = 13). That, not traffic, is the
binding constraint — the operator has confirmed the Note3 SIM is **unlimited**, so the RU arm runs at
full size and no paired-subsample reduction is applied.

### 2.2 Time-of-day zones

The market's clock is Moscow time, so zones are defined in MSK (UTC+3) and stored in UTC:

| zone | MSK | UTC | covered today? |
|---|---|---|---|
| night | 00–06 | 21–03 | **no** |
| morning | 06–12 | 03–09 | partly (NL 16.08 at 4–7 UTC; RU never) |
| day | 12–18 | 09–15 | partly (NL 16.08 at 10–11 UTC; RU never) |
| evening | 18–24 | 15–21 | RU 15–16 UTC, NL control 16–17 UTC |

The RU arm has touched **one** zone. Half the clock — the entire night and most of the morning — has
never been observed from any vantage. Any current statement about "the RU market" is a statement about
two evening hours of one Tuesday.

### 2.3 Size

Per arm, per day: **20 sites × 4 zones × 4 replicates = 320 loads**, i.e. 80 loads per 6-hour zone
≈ 2.1 h of wall clock inside a 6 h window (36 % duty). The headroom is deliberate: it absorbs retries,
redirects and the slow tail without letting a zone bleed into the next one, which would put the zone
label itself in doubt.

Over **14 days**: **4480 loads per arm**. Against today's RU total of 57 loads that is a **79×** increase,
and it is the smallest design that gives each (site × zone) cell 4 replicates per day and 56 across the
run — enough for a per-cell distribution rather than a point.

### 2.4 Pairing and randomisation

- **Pairing is simultaneous, not merely same-hour.** The RU and NL arms load the *same site* at the
  *same minute*, from two network namespaces on this one node (RU via the Note3 tether netns, NL via the
  node's own route). Same-hour-but-not-same-minute is not a pair on a market where floors move.
- **Strict pairs available today: one.** (`kp.ru`, hour 16 UTC) is the only (site, hour) cell present in
  both the RU and the NL-control datasets. Against `hb-auction.jsonl` (16.08) the pair count is **zero**,
  because it is a different day. The 418-load dataset therefore **cannot** serve as the control arm.
- **Site order is randomised within each zone, per day, per arm, from a seed recorded in the run log.**
  Without it, position-in-zone is confounded with site.
- Both arms share one `load_id` prefix per pair so the pairing is recoverable from the data alone and
  not reconstructed later by timestamp arithmetic.

### 2.5 The 418-load dataset is not one arm

`hb-auction.jsonl` mixes two vantages and five arm labels: egress `38.49.216.141` on 217 loads and
`77.246.104.228` on 201, under arms `A`/`B`/`proxy`/`direct`/`direct-via`. Per this node's own operator
notes `38.49.216.141` is the **privoxy far end in Montréal (CA)**, not the node's NL route — so calling
those 418 loads "NL" is wrong on 217 of them. Worse, **`vantage_verified` is true on only 84 of 418**;
on the other 334 the egress was recorded from the shell and `browser_egress_ip` is null, so the browser's
own path was never confirmed — and on this node a browser can carry proxy environment its shell does not.
**Every load in the study runs with `vantage_verified: true` or it is discarded**, and the vantage is read
from the browser, never from the shell that launched it.

The RU arm's own egress is not stable either: 2 distinct MTS addresses across 57 loads
(`91.79.82.121` n=30, `91.79.90.229` n=27). Mobile IPs rotate; the address is recorded per load and is a
covariate, never an identifier of the arm.

## 3. The channel confound

The RU vantage rides HSDPA on the Note3 tether; the NL vantage is a datacenter link. Vantage and channel
are therefore perfectly confounded, and a difference between them is *two main effects wearing the
interaction's name*. The operator's brief offered two exits — put RU on wire, or degrade NL. RU-on-wire
would destroy the thing that makes the RU arm worth having (a resident Russian vantage), so:

**Three arms, not two.**

| arm | vantage | channel |
|---|---|---|
| `ru-mobile` | RU (Note3 tether) | HSDPA, as it is |
| `nl-direct` | NL (node route) | datacenter, as it is |
| `nl-shaped` | NL (node route) | datacenter degraded to the **measured** `ru-mobile` RTT / bandwidth / loss |

`nl-direct` vs `nl-shaped` isolates the channel effect inside a single vantage. `ru-mobile` vs
`nl-shaped` compares vantages at a matched channel. Neither comparison alone would have separated them.

**Prerequisite, and it is a real one:** the shaping target must be measured before it can be applied —
a per-zone RTT/bandwidth/loss profile of the `ru-mobile` path, since a mobile link at 03:00 MSK is not
the same link as at 20:00 MSK. Shaping is applied with `tc netem` on the NL namespace only. It touches
no routing table and no interface the mesh is reached over; it is namespace-local and reversible, and it
is **not** a substrate change to the node's own path.

## 4. Units, and what this market does not let us read

The design requires currency, floor-vs-bid, and `at` on every row. Measured against what the capture
actually produces:

- **currency — present.** `detail.prices[].currency`, e.g. `RUB`, read from the response itself.
- **price — present, with its raw path.** `detail.prices[]` carries `path` (`$[0].cpm`), `field`, `value`.
- **`at` (first vs second price) — absent, and structurally so.** `at` is an OpenRTB *bid request* field.
  These are Adfox `yhb` POSTs with `req_openrtb: false`; there is no `at` on the wire to read. This is a
  property of the market, not a gap in the tool, and it is stated wherever a price is shown.
- **floor — absent on exactly the rows where a price exists.** Both priced examples carry `floor: null`,
  because the floor column is read from OpenRTB `imp[].bidfloor` and these requests are not OpenRTB.
  **A price without its floor cannot be classified**: at the floor and flat are indistinguishable, and
  under a dynamic floor a floor-follower reads as a live valuation. Until a floor source is found for
  Adfox-shaped requests, prices from them are reported as *unclassified*, never as a valuation.
- **`cpm` is the only field this market offers.** The design says to prefer `originalCpm` and never
  `cpm`, precisely because an adjustment smears a discrete price list into a continuous one. Here there
  is no `originalCpm` to prefer. Whether the value has been wrapper-adjusted is **unknown**, and that
  word is printed beside it rather than resolved by assumption.
- **Outcomes stay separated.** RU 18.08, n = 5259 candidate rows: `no-bid` 1891, `no-answer` 1485,
  `binary` 682, `unread` 543, `redirect` 232, `unparsed` 211, `error` 144, **`priced` 13**, `bid-zero` 1.
  A timeout is an **error**, never a loss; `no-answer` is a claim about our observation window and is
  never renamed `timeout`.

**The standing blindness, printed next to every result:** this is a passive capture of what the wrapper
hands the browser. The server-side auction is not observable from here at all. We see the client's half.

## 5. What "presentable to a stranger" requires

Four artifacts, none of them cosmetic, each owed to a specific failure of a repository a stranger
cannot use:

1. **README as a shopfront** — question, method, numbers, how to re-check. The current one describes the
   file layout, i.e. it is written for us.
2. **A datasheet for the dataset** (per *Datasheets for Datasets*): provenance, the selection rule of
   §1.4 verbatim, **what did not enter and why** (§1.5's three rejection reasons, with counts), the
   unobservables of §4, licence CC0. Without a datasheet the data cannot be cited by anyone.
3. **A living report page** at a link, with charts, updated as the run accumulates — and every figure on
   it carrying its `n` and its coverage **in the same line as the number**. The RU arm is one day and 57
   loads; the page must make that impossible to miss.
4. **Reproducibility from the repo alone** — commands, fixtures, and gates, so a stranger needs the
   repository and not our account of it.

## 6. What is fixed here, and what is still open

**Fixed:** the universe and its pin (§1.2, committed); the two-stage rule and the published rejection
ledger (§1.4–1.5); zones, replicates, 14 days, 4480 loads/arm (§2.2–2.3); simultaneous pairing and
per-day randomisation (§2.4); a single `window_s = 45` across all arms (§2.1); `vantage_verified: true`
as an admission requirement (§2.5); the three-arm channel design (§3).

**Open, and named rather than assumed:** the category source (§1.6) — until it is committed, stage B
runs unstratified; the RU-audience allowlist for non-`.ru` TLDs (§1.7 hole 1) and its evidence column;
a floor source for Adfox-shaped requests (§4), without which those prices stay unclassified; and the
per-zone measurement of the `ru-mobile` channel that `nl-shaped` must be shaped to (§3).

**Not claimed anywhere:** that we know what the Russian header-bidding market pays. We have 13 prices.
