# Datasheet — the RU header-bidding frame and its rejection ledger

Following *Datasheets for Datasets* (Gebru et al.). A dataset without a datasheet cannot be
cited: a reader has no way to know what it excludes, and exclusions are where the wrong
conclusions live. **In this dataset the exclusions are not a caveat — they are half the
content.** The rejection ledger is published on purpose.

Files this datasheet describes, all under `public-data/`:

    frame-stageb-<vantage>-<date>.jsonl              every domain the inclusion probe touched
    frame-stageb-<vantage>-<date>-nowrapper-audit.jsonl   the rejections, re-examined later
    ref/ru-frame-stageB-<vantage>-<date>.tsv         the admitted set alone

Current counts are **not written here**, because a hand-copied count rots. Run
`python3 tools/adint-status`, or read the header of each file, which carries its own row
count and the fields dropped at publication.

---

## Motivation

**Why was it created?** To make a claim of the form *"on the Russian header-bidding market,
…"* possible at all. The 31 sites captured before it were assembled by hand with no written
rule, which supports only *"on these 31 sites, …"*. A frame that a stranger can re-run is the
difference between the two sentences.

**Who created it and who funded it?** One operator and the agent working his node, for his own
research. No funding, no client, no commercial interest in the outcome. Licence CC0.

**What gap does it fill?** The literature on header bidding covers US and EU inventory. The
Russian SSPs — `Yandex`, `adriver`, `buzzoola`, `betweendigital`, `sape`, `smi2`, `MyTarget/VK`,
`astralab`, `gnezdo`, `hybrid`, `bidvol`, `alfasense` — appear by name in our own capture and
not in the papers.

## Composition

**What does an instance represent?** One **domain** and one **verdict about it, from one
vantage, at one moment**. Not a site's nature: a verdict is a joint claim about the domain and
about our ability to see it, and the two are separated by the verdict vocabulary rather than
merged into a boolean.

**How many instances?** See the file headers. The universe they are drawn from is fixed:
Tranco list `V3JQN` (created 2026-08-01, 30-day Dowdall combination over crux, farsight,
majestic, radar, umbrella), every registrable `.ru` / `.su` / `.xn--p1ai` domain in the top
50 000 — **n = 1944**, committed as `ref/ru-universe-tranco-V3JQN.tsv` with the source CSV's
sha256 recorded in `docs/step0c-hb-ru-market-study-2026-08-18.md`.

**Is it a sample or the whole?** A sample, drawn by a **declared, mechanical, resumable rule**
walking the universe in rank order until a target number of domains is admitted. The walk is
prefix-complete: every domain above the stopping rank was touched and appears with a verdict.
Nothing below it was looked at, and nothing was hand-picked or hand-removed.

**What does each instance consist of?** The domain and its Tranco rank; the HTTP status and
content type; the wrapper globals found in the page **in every frame, not only the top one**;
the verdict and the named evidence for it; the observation window that produced the verdict;
and the vantage — the exit address as read **by the browser**, with a flag for whether that
read succeeded and how old it is.

**Is any information missing?** Yes, and structurally:

- **The server-side auction is absent entirely.** This is a passive read of what a wrapper
  hands the browser. A domain recorded `no-wrapper` may be selling inventory server-to-server.
- **`blocked` rows carry no information about the site.** We never saw the page. They are
  published so that the frame's coverage is auditable, *not* as evidence of absent ad-tech.
- **Categories are absent.** The design stratifies by category; no citable Russian category
  source has been committed yet, so the walk runs **unstratified** and every table says so.
  This is the largest known weakness — see *Known problems*.

**Does it contain confidential or personal data?** No. Every row is an observation of a
publicly reachable web server, made by a throwaway browser profile with no history, no login
and no prior state. The only addresses recorded are **our own** exit addresses. The operator's
device data lives in a separate, gitignored tree and never enters this one; the boundary is
enforced by `tools/adint-publish`, which copies through an explicit field **allowlist** so
that a column added by a future schema is dropped by default rather than leaked by omission.

**Could it cause harm?** The rejection ledger names sites that did not admit to the frame. It
must not be read as "these sites run no header bidding" — for `blocked` and `unreachable` rows
that reading is simply false, and the file says so in its own header and in every gloss.

## Collection process

**How was the data acquired?** Directly observed: a headless Chromium (Playwright) loaded each
domain's apex URL, waited out a settle window, scrolled to trigger lazily-mounted ad slots, and
read the wrapper globals out of **every frame** of the page.

**What was the sampling strategy?** Two stages, fixed before collection in
`docs/step0c-…` §1.4. Stage A is the TLD-and-rank filter over the pinned Tranco list. Stage B
is the inclusion probe, admitting only on HTTP 200 + an HTML document + an ad wrapper present.

**Over what timeframe?** Each file's rows carry a UTC timestamp; the collection date is in the
filename. A frame is dated because the web is not stable: re-running it later produces a diff,
which is the intended way to use it.

**Were the subjects notified, and did they consent?** No, and no. These are HTTP requests to
publicly reachable servers, of the kind any browser makes, at a volume far below ordinary
traffic — one to two page loads per domain. Where a server answered 428/429 we recorded
`rate-limited` and stopped rather than pressing. No robots directive was circumvented, no
authentication was attempted, no non-public endpoint was touched.

**Any ethical review?** No institutional review. The project's standing constraint — nobody
but the operator is observed as a person — is what governs it, and this dataset contains no
person at all.

## Preprocessing / cleaning / labelling

**Was anything cleaned?** Nothing was removed. Labelling is the verdict, applied by a **pure
function of one probe row** so that it can be tested without a browser and re-derived without
re-collecting.

**Two labelling decisions that materially shape the data, stated plainly:**

1. **Bare `window.Ya` does not admit a domain.** The design named the Adfox signal as
   "`Ya` / `YaHeaderBiddingSettings`", but `window.Ya` is Yandex's *generic* namespace and is
   set by Yandex.Metrika, an analytics product installed across much of the Russian web. A
   rule keyed on it would have filled the frame with sites running no auction at all, and
   every later number would silently describe *"sites using some Yandex product"*. Admission
   requires an advertising global — `Ya.adfoxCode`, `YaHeaderBiddingSettings`, or
   `yaContextCb`. Domains with only the generic namespace get their own verdict,
   `ya-generic-only`, so the size of this correction is a published number.
2. **A navigation failure is split by whose fact it is.** `no-web-apex` (DNS, certificate or
   TLS failure — the domain serves no page at its apex) is a fact about the domain;
   `unreachable` (timeout, refused connection) is our own blindness. They were pooled in the
   first schema and separated in the second after the first walk showed six of seven such rows
   were CDN and nameserver apexes rather than failures to reach a real site.

**Is the raw data kept?** Yes. The probe row is retained verbatim in the ledger, which is why
re-labelling under a new schema costs no page loads (`--reclassify`) and why schemas are never
pooled: a row is re-derived into a new file, and the old one stays as it was.

## Uses

**What has it been used for?** Building the site list for a paired RU/NL header-bidding
capture, and — unexpectedly — as evidence about the frame rule itself (see *Known problems*).

**What should it not be used for?**

- **As a census of Russian header bidding.** It is a rank-ordered prefix of one TLD-based
  universe, unstratified, from one or two vantages, on one date.
- **To claim any named site "does not run header bidding".** Only `no-wrapper` even gestures
  at that, and only within our detector's coverage and our observation window.
- **To compare against any dataset collected with a different observation window.** The window
  is what separates "no answer" from "no bid"; comparing across windows manufactures a
  difference. Every row carries the window that produced it.

## Known problems — the ones that would change a conclusion

1. **The two-stage rule does not solve the problem it was written for.** The design rejects the
   naive "top-N `.ru`" rule because it selects CDNs, registrars and infrastructure rather than
   ad-supported content. Stage B admits on *"does this page run an ad wrapper"* — and
   registrars and portals do. The first admissions included exactly that class of domain. The
   category source deferred as "open" in §1.6 is therefore **load-bearing, not cosmetic**:
   without stratification the frame reproduces the very failure the two stages were introduced
   to prevent. This is measured in the ledger rather than argued.
2. **The detector's coverage bounds every `no-wrapper` row.** It looks for Adfox and Prebid,
   the two wrappers the study is about. Google's `googletag` and Amazon's `apstag` were added
   later as *recorded evidence* (never as admission criteria), because a `no-wrapper` verdict
   that cannot see them is a claim about the detector wearing the clothes of a claim about the
   site. Rows written before that addition cannot make the distinction at all.
3. **The frame is vantage-dependent — measured, not feared, and the study now runs on the
   RU-built frame** (`ref/CANONICAL-FRAME`). At 102 domains reached by both walks, an
   Amsterdam-built frame is missing three domains a Moscow-built one admits: `rbc.ru` and
   `gismeteo.ru` (refused 401/403) and `hh.ru` — which is **not refused from Amsterdam at
   all**. It answers 200 and serves a page carrying no auction config, while the same URL
   from Moscow carries `YaHeaderBiddingSettings`. So this is not only geo-blocking: a page
   you can load can still be the wrong page, and no HTTP status in the ledger reveals it.
   The reverse hole does not exist — every domain Amsterdam admits and Moscow has reached,
   Moscow admits too. Any user of this dataset must read the `vantage` column as part of the
   verdict, never as metadata about it.
4. **Unstratified means rank-ordered, and rank correlates with everything.** Until categories
   are committed, the admitted set is biased toward whatever the top of the Russian resolver
   traffic distribution happens to be.

## Distribution and maintenance

**How is it distributed?** In this public repository under **CC0 1.0**, as newline-delimited
JSON with a header naming its source, its rule, its licence and the fields dropped at
publication.

**Who maintains it, and will it be updated?** The operator's `adint` window. New walks are
added as new dated files. **Files are not edited in place** — a re-labelling writes a new file
at the new schema, so a citation of an older file keeps meaning what it meant.

**What happens to older versions?** They stay. Git history is the archive.

**Can others contribute or extend it?** Yes, without asking. The universe is pinned and
re-fetchable, the rule is a runnable tool, and `docs/REPRODUCE.md` is the whole procedure. A
walk from a third vantage would be the single most useful contribution — it would measure how
much of the frame is a property of where the observer stands.
