# step0d — the first frame built under step0c's rule, and the three places the rule did not survive contact

date: 2026-08-19 · by: the `adint` window · status: **frame walk in progress; the rule was
corrected twice while running it, and one paired cell was captured.** Every count below is
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

## 4. The frame is NOT vantage-dependent at this depth — a null result, stated as one

`blocked` is the verdict most likely to be geo-shaped, and a frame built from one vantage
would carry that vantage's blindness into *both* arms of a paired study. So the walk is run
from both.

On the **25 domains both walks had reached** (ranks 13–457): **25 agree, 0 differ.** All four
`blocked` verdicts — `okcdn.ru` 403, `vkuserphoto.ru` 418, `ozon.ru` 403, `wildberries.ru`
498 — are identical from Amsterdam and from a Moscow MTS address. At this depth those
refusals are **bot detection, not geography**.

**Coverage binds the claim.** That rank band is infrastructure, telecoms and marketplaces.
The cases most likely to be geo-shaped are news publishers above rank 1100, where the NL walk
already records `rbc.ru` 401 and `gismeteo.ru` 403 — and the RU walk has not reached them.
The comparison is computed on the intersection only: a domain one walk has reached and the
other has not is a difference in *progress*, and pooling those would make a slow walk look
like a censoring vantage.

    tools/adint-frame-compare data/frame-stageb-nl-direct-<date>.jsonl \
                              data/frame-stageb-ru-mobile-<date>.jsonl

## 5. The paired cell: the channel confound is not a caveat, it is the whole signal

One simultaneous pair, `pikabu.ru`, both arms launched in the same minute by
`tools/adint-paired-run` (**barrier wait 0.0 s**), same 19-bidder roster read off the page,
same wrapper deadline `wrapper_timeout_ms = 1000`, `window_s = 45` on both:

| arm | egress (read by the browser) | `no-bid` | `no-answer` | candidates |
|---|---|---:|---:|---:|
| `nl-direct` | 77.246.104.228 · Amsterdam, datacenter | **19** | 2 | 33 |
| `ru-mobile` | 91.79.81.62 · MTS Moscow, HSDPA | 4 | **17** | 32 |

The arms invert. `no-bid` is a bidder answering *no*; `no-answer` is a bidder asked and
silent inside our window. Measured RTT from inside the `ruvantage` namespace is **400–1100 ms**
against a wrapper deadline of **1000 ms** — so on this link a bidder that answers perfectly
promptly still misses the publisher's own deadline, and the page records silence.

**Any statement of the form "Russian bidders are less responsive", read off this arm, is an
artifact of the phone's link.** This is §3's confound — vantage and channel perfectly
confounded, a difference between them being *two main effects wearing the interaction's
name* — no longer argued but measured, on one pair, in one minute.

It also settles that §3's third arm is load-bearing rather than tidy. Without `nl-shaped`
(the NL egress degraded with `tc netem` to the RU arm's measured profile) there is no way to
attribute the gap, and the study cannot report a single number about Russian demand.

**A cheap tell for every future run:** compare `wrapper_timeout_ms` against the arm's RTT
*before* reading its outcome mix at all. If the deadline sits inside the RTT distribution,
that arm is not measuring bidding — it is measuring its own latency.

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
