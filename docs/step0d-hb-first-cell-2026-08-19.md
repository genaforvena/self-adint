# step0d — the first frame built under step0c's rule, and the three places the rule did not survive contact

date: 2026-08-19 · by: the `adint` window · status: **frame walk in progress; the rule was
corrected twice while running it, one paired cell was captured, and the walk's own first
null result expired on schedule (§4). §9 adds the second cell — the first against the
CANONICAL frame — and the three instrument defects it exposed, one of which silently cost
that cell four of its fourteen RU loads. §10 is the verdict the fixed instrument then
delivered on the VANTAGE: the phone's RTT straddles the wrapper's own deadline, so the RU
arm's dominant outcome column is unreadable and §2.3's schedule must not be run against
it.** Every count below is
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

### The frame was rebuilt from the RU vantage (operator, 2026-08-19)

Done rather than recommended. `ref/CANONICAL-FRAME` now declares `ru-mobile` as the frame the
study runs on, and the declaration carries the measurement that justifies it. At 102 domains
reached by both walks (ranks 13–2463):

| | admits | hole — the other vantage admits it, this one reached it and refused |
|---|---:|---|
| `nl-direct` | 8 | **3** — `rbc.ru` #1163, `gismeteo.ru` #1169, `hh.ru` #1754 |
| `ru-mobile` | 6 | **1** — `lenta.ru` #3128 |

**The first version of this section said the asymmetry was one-directional and that a
Moscow-built frame lost nothing. That was true of the data it was written on (ranks 13–2463)
and stopped being true fifteen minutes later**, when the RU walk reached rank 3128 and read
`lenta.ru` as `ad-serving-only` while Amsterdam admits it with two bidders.

So: **neither single-vantage frame is complete.** `hh.ru` is admitted only from Moscow,
`lenta.ru` only from Amsterdam, and *both are content differences rather than access* — each
vantage is served a page the other is not. The union (11 domains at ranks 13–3128) is
strictly better than either. The declaration stays at `ru-mobile` because that is the
operator's instruction and it is still the smaller hole, 1 against 3; moving to a union frame
is his call and carries a real cost, since a union is no longer a rank-ordered prefix of one
walk and its reproduction instructions change.

**`lenta.ru` is not a flake and not a config change, and the reason we can say so is the
interleaving** — the cheap trick for separating a per-vantage difference from a per-time one:

| time | vantage | verdict | `hb_settings` | window |
|---|---|---|---|---|
| 07:10:42Z | `nl-direct` | `admit` | true (2 bidders) | 8 s |
| 07:36:23Z | `ru-mobile` | `ad-serving-only` | false | 45 s |
| **07:38:47Z** | **`nl-direct`** | **`admit`** | **true (2 bidders)** | **45 s** |
| 07:42:45Z | `ru-mobile` | `ad-serving-only` | false | 45 s |

An NL read sits *between* two RU reads, and both re-probes ran at the same window. Time
cannot explain the split and neither can the window. The vantage can. Given the measured
per-probe miss rate of 15–25 % on known-HB sites (§3), a single disagreeing read would not
have been enough; four interleaved reads are.

**`hh.ru` is the one that widens the finding.** It is *not* blocked from Amsterdam — it
answers 200 and renders. From there it carries only the generic Yandex namespace and no
auction config; from Moscow it carries `YaHeaderBiddingSettings`. **Vantage-dependence is
therefore not only about access. A page you can load can still be the wrong page**, and no
HTTP status anywhere in the ledger would tell you.

`lenta.ru` #3128 and `iz.ru` #3322 are admitted by the deeper `nl-direct` walk and are
**pending** for the RU walk, which has not reached them. They are not added by hand: the
frame is a rank-ordered *prefix* of the universe, and admitting a domain out of order would
destroy the one property that lets a stranger reproduce it. They enter when the walk arrives.
`tools/adint-frame-compare` now separates a HOLE from a PENDING by construction, because the
two are indistinguishable in any disagreement total and only one of them is a defect.

### Re-measured at 13:19Z, on the full ledgers — and the reason it had to be re-measured

**The numbers in the table above were derived over a ledger that mixed two classifier
schemas, and the derivation trusted the verdict stored in each row.** `tools/adint-frame-stageb`
stamps `stageb_schema` on every row precisely because the rule can change under a running
walk — and it did, at `495148a`, when the frame rule stopped admitting Adfox *ad serving* as
header bidding. The 35 nl / 16 ru rows walked before that fix kept their old verdicts, and
`write_frame` read the field instead of re-deriving it. The effect was not subtle and it was
not at the margin: **7 domains re-entered the nl frame and 6 the ru frame as `admit` on
`Ya.adfoxCode+yaContextCb` alone — `mail.ru` #13, `dzen.ru` #18, `nic.ru` #93, `yandex.ru`
#98, `reg.ru` #148, `ok.ru` #310, `imgsmail.ru` #563 — the top of the rank order, ahead of
every genuine admission.** The header said `schema 3` over a table that was partly schema 1.

The sibling tool had it right the whole time: `adint-frame-compare` re-derives verdicts from
the stored probe rows by design, and says so in its docstring — *"the comparison is between
the pages, not between the days the tools were edited."* The discipline existed in the
comparison path and was missing from the derivation path, which is why the disagreement
tables were sound while the frames were not. `write_frame` now re-derives any row whose
`stageb_schema` is not current (an *unstamped* row counts as stale — assuming it current is
the assumption being removed) and prints how many it moved and in which direction, so a
correction can never be silent:

    # re-derived 35 row(s) written under an older stageb_schema: admit->ad-serving-only=7 …

**A stored verdict is a claim about the page AND about the code that read it.** The raw
fields beside it are facts that do not age; the verdict ages the moment the classifier is
fixed. Re-derive on read, or the fix reaches only the rows walked after it.

With that repaired and both walks carried to their current depth (ranks 13–5117 reached by
both):

| | admits | hole — the other vantage admits it, this one reached it and refused |
|---|---:|---|
| `nl-direct` | 13 | **5** — `rbc.ru` #1163, `gismeteo.ru` #1169, `hh.ru` #1754, `aviasales.ru` #4583, `interfax.ru` #5117 |
| `ru-mobile` | 14 | **2** — `lenta.ru` #3128, `liveinternet.ru` #4678 |
| union | **18** | |

Direction held, every number moved — including which vantage admits more, which has
reversed. **And the hole counts should be read split, not pooled:** 3 of `nl-direct`'s 5
(`rbc.ru`, `gismeteo.ru` blocked; `interfax.ru` unreachable) and 1 of `ru-mobile`'s 2
(`liveinternet.ru` unreachable) are *our own blindness at that moment*, which a later walk
from the same vantage may close. The rest are the site serving that vantage a different
page, and no retry closes those. Quoting **5 against 2** without the split overstates how
permanent the gap is; the content-difference count is **2 against 1**.

The pending line from 07:46Z closed exactly as a pending line should: the RU walk reached
`iz.ru` #3322 (admitted) and `lenta.ru` #3128 (`ad-serving-only`, the hole above). Now
pending for RU: `sportbox.ru` #5813 and `playground.ru` #6336.

**Consequences, and they bind the design rather than decorating it:**

1. **The frame is built from the RU vantage** — done, above, and it is the smaller hole
   rather than no hole. The union of both with every disagreement published is now strictly
   better than either single vantage, and is what a third vantage would require anyway. An `nl-direct` frame is a frame of *Russian sites Amsterdam is allowed to
   load*, which is a different population and never says so.
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

> **This cell was captured on the `nl-direct` frame, which is now superseded.** It ran on
> `pikabu.ru` and `magnit.ru` — both of which the canonical RU frame also admits, so nothing
> here is invalidated. What it could not include is `rbc.ru`, `gismeteo.ru` and `hh.ru`,
> which that frame did not contain. The next cell runs on the canonical frame.

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
   capture.** Built from the RU vantage as of 2026-08-19 (`ref/CANONICAL-FRAME`), and every
   rejection statistic carries the vantage it was taken from (see §4). The clause is stronger
   than "geo-blocking exists": `hh.ru` serves a *different page* to the two vantages without
   refusing either.
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

---

## 9. The second cell, and the instrument that refuses a working vantage

Added 2026-08-19 after the first cell run against the **canonical** frame
(`ref/CANONICAL-FRAME` → `ru-mobile`, 14 admitted sites). The cell above was run on two sites
of the *nl-direct* frame while its walk was still in progress; this is the first one on the
frame the study actually declares.

`run_id=2026-08-19-day-2026081914` · 14 sites × 1 replicate · window 45 s. Its own summary
line said **`done · pairs=14`**. On disk:

```
data/hb-paired-2026-08-19-day-2026081914.nl.jsonl   14 loads
data/hb-paired-2026-08-19-day-2026081914.ru.jsonl   10 loads
```

**Four of the fourteen attempts have no RU side at all**, so the cell holds ten pairs and four
unpaired NL loads — and the number the run printed was the one that could not be checked
against anything. Three separate defects, each of which survives every gate the tools had.

### 9.1 The vantage proof fails transiently on a vantage that is working

Each load verifies its own arm from inside the browser before the site is opened: read an IP
echo, compare it to the arm's expected exit. Unread ⇒ `vantage_verified: null` ⇒ the run
refuses (rc 2) rather than record a load it cannot label. That refusal is right, and it is not
what went wrong. What went wrong is that **one unread echo was the whole verdict.**

Measured on the RU tether during the cell (n = 8, `curl --max-time 30` to the same echo, from
inside the `ruvantage` netns):

| read the address | did not |
|---|---|
| 12.1 s · 12.5 s · 23.9 s · 26.3 s | 19.8 s · 19.8 s (connection failure) · 30 s · 30 s (timeout) |

Four of eight — **while the same tether carried eleven 45-second captures in the same hour.**
The echo fails transiently on a link that is up. One attempt turns that into `vantage
unproven`, and the study loses the load.

**The obvious diagnosis is wrong and worth writing down.** The module constant is
`ECHO_TIMEOUT_MS = 12000`, every successful read above took ≥ 12.1 s, and it is very tempting
to stop there. But that constant is overridden at startup from `--nav-timeout` — 120 s in a
paired cell — so the browser had ten times the patience of the slowest success and still read
nothing. Patience was never the binding constraint; raising it would have fixed nothing and
the fix would have been reported as working, because the failure is intermittent enough to
disappear on the next run either way.

The fix is `--echo-tries` (`tools/adint-hb-capture`), default **1** — a retry is a claim about
the link and has to be asked for — and **3** in a paired cell, *identical for both arms*. An
allowance given only to the flaky arm is an arm difference that is not the treatment: the pair
would then differ in how hard we tried to prove it, which is the confound pairing exists to
remove. The attempt count is recorded on every load (`browser_egress_attempts`) and printed at
preflight, because **a retry that hides its own count converts a degrading link into a clean
one** — the count is the only place the degradation is visible.

### 9.2 `pairs=N` counted attempts, and an attempt is not a pair

An arm's return code is a cause, not a boolean: rc 2 is the refusal above (no row reaches
disk), rc 1 is the two arms having collapsed onto one exit. Neither produced a pair and
neither is an error in the ordinary sense, so a summary counting attempts folds three facts
into the study's headline number. The runner now reports

```
done · attempted=14 · complete=10 · ru-mobile:vantage-unproven=4
```

`complete` counts only attempts where **every** arm recorded; the rest are attributed by arm
and by cause, with an unrecognised return code kept as `error-rc<n>` rather than absorbed into
a named one.

### 9.3 A cell that overruns its window keeps a zone label that is wrong for its tail

The MSK zone is the study's **treatment axis** and it is stamped once, at the cell's start.
This cell began 14:25 Z (MSK day) and its last four pairs landed after 15:00 Z, which is MSK
evening — under a `run_id` that says `day`. The id is not rewritten (it is the join key, and
the pairs before the boundary really are day); instead each pair re-derives its own zone and
the drift is named per pair and summarised, so a consumer can drop or re-label the tail rather
than trust the id.

### 9.4 Two contaminations, recorded because nothing else would show them

- **The tether measurement above perturbed the cell it measured.** Those eight curls ran
  14:52–14:57 Z, and the RU arm's duration for the two sites in that window jumps to 258 s and
  240 s against a ~140 s median. The vantage probe and the capture share one HSDPA link.
- **A concurrent `adint-frame-stageb` walk rewrote the canonical frame at 14:49 Z, mid-cell**
  (lock pid 3516205, since exited). The admitted set is unchanged — only the header tallies
  moved — so the cell stands, but nothing in this repository claims single-writer on a ledger
  a running cell is reading, and the next collision need not be this harmless.

### 9.5 What this does not change

Nothing here is a result about the Russian market. The frame holds fourteen sites, the study
holds ten complete pairs from one zone, and §2.3's design asks for 4480 loads per arm. Every
fix above makes the instrument's own failures countable; none of them makes the sample bigger.

---

## 10. The vantage cannot answer the question, and the auction data is how we know

§9 fixed the instrument. This section is about what the fixed instrument then showed, which
is a verdict on the **vantage** rather than on any tool: **a tethered HSDPA phone cannot
observe an auction whose deadline is the same order as its own round-trip time.**

### 10.1 The contrast that looks like a finding

Evening cell, 14 sites × 4 replicates, both arms launched simultaneously per site, identical
45 s window, read by `tools/adint-hb-report`:

| arm | rendered | dominant outcome | ever returned a price |
|---|---:|---|---:|
| `nl-direct` | 44/56 | **`no-bid`** — MyTarget/VK 34 vs 2, adriver 24 vs 2, roxot 16 vs 0 | 4 |
| `ru-mobile` | 30/56 | **`no-answer`** — adriver 3 vs 12, sape 0 vs 10, betweendigital 0 vs 9 | **0** |

Read carelessly this says Russian SSPs answer a Dutch datacenter and ignore a Russian phone.
It is one of the most publishable-looking things this study has produced and it is an
artifact of our own channel.

### 10.2 The latencies make the mechanism visible

Among the bidders that **did** answer, median response time:

```
ru-mobile   614 – 944 ms        the wrapper's own deadline: 700 – 1000 ms
nl-direct   346 – 524 ms
```

The RU responses sit **at the publisher's deadline**. The bidders recorded as answering are
the ones that just made it; everything slower is recorded as silence. And note why
`after-deadline` reads ≈0 on *both* arms — a late answer is not recorded as late, it is
recorded as **absent**. The column that would have named this confound cannot see it.

This confirms at n=56 per arm what `step0b` §"the confound this treatment carries" asserted
from a single pair, and it upgrades the claim: it is not merely that the channel *could*
explain the difference, it is that the RU arm's response times straddle the exact threshold
the outcome column is computed against.

### 10.3 What follows for §2.3's schedule

**Running the 14-day, 8960-load schedule against this vantage would buy 4480 RU loads whose
dominant outcome column is unreadable.** The fixes in §9 make the instrument honest; they do
not make the phone fast. Nothing in more replicates addresses a confound that lives in the
millisecond budget of every single load.

What the phone is still the only source for: a genuinely Russian *mobile subscriber* identity
— the SIM, the carrier resolver, the operator's address space. What it cannot do is meet a
1 s auction deadline. Those are separable, and the study should stop asking one device for
both:

1. **A Russian vantage on a fast link** (VPS in RU, or a residential fixed line) becomes the
   measurement arm. Its RTT must be small against the wrapper deadline, and that should be
   *measured and published beside every cell*, not assumed.
2. **The phone stays** for what genuinely needs a mobile subscriber, and for exactly the
   comparisons where a slow link is not disqualifying.
3. **Publish the RTT/deadline ratio as a coverage statistic.** A cell whose arm RTT is a
   large fraction of the wrapper deadline cannot report `no-answer` as a fact about a bidder,
   and the reader should be able to see that from the table rather than from this paragraph.

### 10.4 A second limit, smaller but worth knowing now

The NL arm returned **4 priced bids across ~280 candidate requests**. `adint-hb-report` says
it plainly — "ONE curve, not a set of shapes" — so **price-distribution analysis is
unsupported on either arm**, independently of the vantage problem. A study that promised
price shape per DSP would have discovered this after fourteen days instead of after two
cells.
