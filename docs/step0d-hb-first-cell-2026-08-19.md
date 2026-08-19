# step0d — the first frame built under step0c's rule, and the three places the rule did not survive contact

date: 2026-08-19 · by: the `adint` window · status: **frame walk in progress; the rule was
corrected twice while running it, one paired cell was captured, and the walk's own first
null result expired on schedule (§4).** Every count below is
re-derivable — the commands are named beside each claim. Tallies are deliberately NOT typed
in: run `python3 tools/adint-status` and `tools/adint-frame-compare`, which read the ledgers
on disk. A number typed into prose is true on the day it is typed.

## 0. What this document is

`step0c` fixed a sampling frame and a schedule on paper, and its own status line said
"capture not yet re-run under them". It could not be: **no tool existed that could execute
stage B**, so there was no frame to capture under. This is stage B made runnable
(`tools/adint-frame-stageb`), the first walk of it, and what the walk immediately revealed
about the rule it was executing.

Nothing here is a result about the Russian market. Two sites are in the frame at the time of
writing. What is here is a corrected instrument and three measurements about our own
blindness — which is what a study owes before it owes numbers.

## 1. The rule admitted ad SERVING, and the study measures header BIDDING

The first walk admitted seven of the top 34 `.ru` domains: `mail.ru`, `dzen.ru`, `nic.ru`,
`yandex.ru`, `reg.ru`, `ok.ru`, `imgsmail.ru`. Reading the ledger back — which is the only
reason this was catchable — **every one of them has `YaHeaderBiddingSettings` absent, no
prebid global, and zero named bidders.** All seven admitted on `Ya.adfoxCode` and
`yaContextCb`.

Those are not the same thing as header bidding:

- `Ya.adfoxCode` places an **Adfox advertisement**. `yaContextCb` queues the **Yandex
  advertising context**. A page carrying both sells advertising whose demand is resolved
  *inside Yandex*; nobody else is asked.
- **Header bidding is the publisher asking several bidders from within the page.** The
  observable that says so is the wrapper's own configuration: `YaHeaderBiddingSettings` with
  its bidder roster, or a prebid global in any frame.

So the frame contained **zero instances of the phenomenon the study exists to measure**,
while looking like a healthy frame of ad-supported sites. A 14-day paired capture run against
it would have measured two registrars and an image CDN, at full cost, and produced a table
nobody could tell was empty.

**Two widenings, and closing one is what made the other invisible.** §1.4 named the signal as
"`Ya` / `YaHeaderBiddingSettings`". Bare `window.Ya` is Yandex's *generic* namespace — Yandex
Metrika sets it, and Metrika is on much of the Russian web — so that rule fills a frame with
analytics. Narrowing it to the advertising globals closed the **Metrika axis** and left the
**serving-versus-auction axis** untouched, and the rule then *looked* corrected. Fixing one
width tells you nothing about the other.

Admission now requires an auction configuration. `ad-serving-only` is kept as its own
verdict rather than folded into `no-wrapper`, because these publishers demonstrably do sell
advertising and erasing that would be a second false claim in place of the first.

Reclassifying the 34 domains already walked cost **no page loads at all** — the ledger keeps
the probe row verbatim, so `--reclassify` re-derives the verdict and writes a new file at the
new schema rather than editing the old one. Seven `admit` became `ad-serving-only`, and the
frame in the top 34 is **empty**.

**That emptiness is a result, not a failed run.** Header bidding does not live at the top of
the `.ru` rank order. The first genuine admission is `magnit.ru` at rank **#721** with 3
bidders; the second is `pikabu.ru` at rank **#1146** with 19. This is §1.3's finding —
"Tranco ranks by resolver traffic, so it answers *what is resolved*, not *what a human reads
with ads on it*" — arriving a second time, from the other side, and much later in the list
than the two-stage rule assumed.

**Cross-check, in the direction that would have caught the opposite error.** If the new
criterion were too narrow it would reject real header bidding. All 15 sites of the 18.08 RU
capture where header bidding was actually observed — `pikabu.ru`, `www.kp.ru`,
`www.fontanka.ru`, `www.sports.ru`, `www.mk.ru`, `www.woman.ru`, `3dnews.ru`, `74.ru`,
`eva.ru`, `overclockers.ru`, `ura.news`, `www.e1.ru`, `www.ixbt.com`, `www.sport-express.ru`,
`www.zr.ru` — carry a roster. The criterion selects exactly the observed population.

## 2. `error` was pooling two opposite facts

Of the first eight navigation failures, **seven were DNS, certificate or TLS failures on CDN
and nameserver apexes** (`yccdn.ru`, `geobasket.ru`, `tld-servers.ru`, `edgecdn.ru`,
`sberbank.ru`, …) and **one was a timeout**. "This domain serves no page at its apex" is a
structural fact about the domain — §1.3's infrastructure finding again, from yet another
side. "We could not reach it" is our own blindness. Split into `no-web-apex` and
`unreachable`. An unrecognised failure stays `error`, so a growing residue is the signal that
the table needs another line rather than a silent misfiling.

## 3. A single inclusion probe misses roughly one real HB site in five

Restricted to the 18 sites where a bidder roster was *ever* observed, the per-load detection
rate is **255/299 at `window_s = 8`** and **42/56 at `window_s = 45`**. So a single probe
fails to see a genuine wrapper on the order of 15–25 % of attempts.

**The two figures must not be read as a window effect.** The 45 s set is the RU-vantage arm
and its control; the 8 s set is the earlier NL sweep. Window and vantage are confounded in
exactly the way §2.1 and §3 warn about, and this data cannot separate them. What survives
that caveat is the part that matters for the frame: **a one-shot rejection is unreliable**,
which is why every non-admitting verdict describing a page we actually saw is re-probed at
the full 45 s window before being written down. Two attempts, not one — and they are not
independent (same vantage, minutes apart), so the residual miss rate is real and unquantified.

## 4. The frame IS vantage-dependent, and the hole falls exactly on the publishers

`blocked` is the verdict most likely to be geo-shaped, and a frame built from one vantage
would carry that vantage's blindness into *both* arms of a paired study. So the walk is run
from both.

**The first reading said no.** At 43 domains reached by both walks (ranks 13–865): 42 agreed,
1 differed, and every `blocked` verdict was identical from Amsterdam and from Moscow. That
was reported as a null result — with its coverage bound stated in the same breath: *"that
rank band is infrastructure, telecoms and marketplaces; the cases most likely to be
geo-shaped are news publishers above rank 1100, which the Moscow walk has not reached."*

**Two hundred ranks later the caveat expired exactly as written.** At 63 domains (ranks
13–1347): **59 agree, 4 differ (6.3 %)** — and the headline understates it.

| rank | domain | `nl-direct` | `ru-mobile` |
|---:|---|---|---|
| #615 | `cdnvideo.ru` | `no-wrapper` | `ya-generic-only` |
| **#1163** | **`rbc.ru`** | **`blocked` (401)** | **`admit` — 10 bidders** |
| **#1169** | **`gismeteo.ru`** | **`blocked` (403)** | **`admit` — 2 bidders** |
| #1315 | `russianpost.ru` | `unreachable` | `no-web-apex` |

Two of the seven domains where either vantage said `blocked` **become admissions from the
other side**. `rbc.ru` and `gismeteo.ru` are real publishers running real header-bidding
auctions with readable bidder rosters, and from Amsterdam they do not appear to exist.

**This is not a rounding error in a frame. It is a hole in it, and it falls precisely on the
class of site the study is about.** The other two disagreements are between *rejection*
reasons and change nothing; these two change the frame's membership.

**Consequences, and they bind the design rather than decorating it:**

1. **The frame must be built from the RU vantage**, or from the union of both with every
   disagreement published. An `nl-direct` frame is a frame of *Russian sites Amsterdam is
   allowed to load*, which is a different population and never says so.
2. **A paired study on an NL-built frame compares the arms on sites selected by one of
   them** — and the excluded ones are systematically the news publishers, i.e. exactly where
   display demand concentrates.
3. **`blocked` counts are not portable between vantages.** Any rejection-ledger statistic must
   carry the vantage it was measured from, which it now does per row.

Note what worked here: nothing was retracted, because the earlier claim was never made
larger than its evidence. The coverage bound was doing real work, and it is the reason this
reads as a claim expiring on schedule rather than as a result being reversed.

    tools/adint-frame-compare data/frame-stageb-nl-direct-<date>-schema3.jsonl \
                              data/frame-stageb-ru-mobile-<date>-schema3.jsonl

## 5. The paired cell: the channel confound is not a caveat, it is the whole signal

**Six simultaneous pairs**, two sites (`pikabu.ru`, `magnit.ru`), three replicates, MSK
morning zone — a zone the RU arm had never observed. Both arms launched on one site at a
time and the next site waits for both: **median arm-duration gap 8.8 s**, and the first pair
0.0 s. `window_s = 45` on both. Every load carries its exit address as the *browser* read it,
and 12 of 12 are verified.

| arm | egress (browser-read) | `no-bid` | `no-answer` | `priced` | `unread` | `binary` | `redirect` |
|---|---|---:|---:|---:|---:|---:|---:|
| `nl-direct` | 77.246.104.228 · Amsterdam, datacentre | **75** | 7 | 1 | 71 | 113 | 10 |
| `ru-mobile` | 91.79.81.62 · MTS Moscow, HSDPA | 49 | **34** | 0 | 45 | 87 | 13 |

Note first what does *not* move: `unread` and `binary` — our own read failures — land at
comparable rates on both arms. The difference is concentrated exactly where the design said
it would be.

**Silence rate**, the fraction of bidders that were asked and said nothing rather than
declining, over near-identical denominators:

> `nl-direct` **8.5 %** (7 of 82)  ·  `ru-mobile` **41.0 %** (34 of 83)

### The same fact, measured a second way, independently

`no-answer` is our window's verdict. The bidder's own clock says the same thing without
reference to it. Per-bidder latency — `t_response_s − t_request_s` — against each page's own
`wrapper_timeout_ms`:

| arm | median bidder latency | answers arriving after the wrapper's deadline |
|---|---:|---:|
| `nl-direct` | **171 ms** | 81 of 595 — **14 %** |
| `ru-mobile` | **659 ms** | 198 of 483 — **41 %** |

The late-arrival fraction on the RU arm (41 %) and its silence rate (41.0 %) are two
different measurements of one phenomenon and they agree. Wrapper deadlines observed on these
pages are **500 and 1000 ms**; measured RTT from inside the `ruvantage` namespace is
**400–1100 ms**. On that link a bidder that answers perfectly promptly still misses the
publisher's own deadline, and the page records silence.

**So "Russian bidders are less responsive" cannot be read off this arm at all.** That is §3's
confound — vantage and channel perfectly confounded, a difference between them being *two
main effects wearing the interaction's name* — no longer argued but measured. §3's third arm
(`nl-shaped`: the datacentre link degraded with `tc netem` to the RU arm's measured profile)
is what separates them, and it is now required rather than prudent.

**A cheap tell for every future run:** compare `wrapper_timeout_ms` against the arm's RTT
*before* reading its outcome mix at all. If the deadline sits inside the RTT distribution,
that arm is not measuring bidding — it is measuring its own latency.

### We see bids the auction threw away

One price was observed: `hybrid` (`ssp.hybrid.ai`) returned **70.0 RUB CPM** on `magnit.ru`
from the NL arm — with a bidder latency of **1747 ms against that page's 500 ms wrapper
deadline**. It is a real bid and it is *not* a price the auction could have used. Our window
is deliberately longer than the wrapper's, so this will be common; pooling late answers with
in-time ones would overstate both the depth of demand and the clearing level. Every request
row now carries `bidder_latency_ms` and `answered_after_wrapper_deadline`, so "what the
auction saw" and "what the bidder was willing to pay" stay separable.

**And that flag was wrong on its first attempt, which is why it is worth stating how.** It
first compared `t_response_s` — measured from *navigation start* — against a deadline that
runs from *auction start*, and reported **100 % of answers late on both arms**. A flag that
fires on everything has stopped being a measurement, and that is the only thing that made it
visible; nothing about the number itself looked wrong.

## 6. Two failures of ours, recorded because they were invisible

**Three walks wrote to one ledger for several minutes.** Launcher shells had been killed and
the message read `Session terminated, killing shell` — which is about the *shell*; the walks
outlived it. The only symptom was one domain appearing three times. A longer overlap would
have inflated every verdict count by however many walkers there were, in a file that reads as
one honest walk. Now: a non-blocking `flock` on the ledger whose refusal is loud and exits
non-zero, and de-duplication at *derivation* — the ledger stays the append-only record of
what was actually done, the dataset is one row per domain, and the number of collapsed rows
is printed, because a silent de-duplication looks exactly like a clean run.

**A `no-wrapper` verdict ages differently from an `admit`.** `admit` stands: we saw the
wrapper. `no-wrapper` only ever means *we looked with this detector*, and it silently changes
meaning the moment the detector learns a signal — without changing value. `drom.ru` (#324, a
large ad-supported classifieds site) came back `no-wrapper` at the full 45 s window; Google's
`googletag` and Amazon's `apstag` were invisible to the probe entirely. They are now recorded
as evidence — never as admission criteria, since the study's object is Adfox and prebid — and
`--audit-no-wrapper` exists to go back to the rejections when the detector improves.

## 7. What this changes in step0c

0. **§2.5 gains a frame clause: the FRAME is a measurement from a vantage, not just the
   capture.** It must be built from the RU vantage or from the published union of both, and
   every rejection statistic carries the vantage it was taken from (see §4).
1. **§1.4's condition (ii) is replaced.** Admission requires an auction configuration, not an
   advertising global. The intermediate population is preserved as `ad-serving-only`.
2. **§1.6's category source is upgraded from "open" to load-bearing.** Unstratified, the walk
   admits whatever the rank order offers; the corrected criterion happens to exclude
   infrastructure as a side effect, but that is luck, not stratification.
3. **§2.3's cost estimate needs re-deriving.** It assumed 20 sites per category were available
   near the top of the list. At the observed density the walk must go far deeper, and the
   binding cost of the study is the FRAME, not the capture.
4. **§3's third arm is promoted from prudent to required** (see §5 above).
5. **§2.4's simultaneity is achievable and cheap** — one barrier per site, measured skew 0.0 s
   — but only because the runner idles the fast arm. Two independent walks cannot deliver it.

## 8. What is still not claimed

That we know anything about what the Russian header-bidding market pays. The frame holds two
sites. Every number above is about our instrument.
