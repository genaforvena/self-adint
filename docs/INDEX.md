# What is in `docs/`, and which thread it belongs to

The documents in `docs/` accumulated across **two different lines of work**, and their file
names do not say which. This index is the missing signpost. It is a map, not a summary: each
entry says what the document DECIDES, so you can tell from here whether you need it.

> **No count is written here on purpose.** This paragraph used to open "Eighteen documents
> accumulated in four days"; by the time anyone read that sentence it was twenty-five across
> seven, and nothing in the file could tell you so. A total maintained by hand beside the
> roster it counts rots the moment the roster grows, and it rots *silently* — the wrong
> number is as readable as the right one. The roster below is the count, and nothing but the
> roster is. There is no gate that a new document reaches this index — adding one is a hand
> step, and `step0f`/`step0g` sat here unlisted for a day, which is how this paragraph exists.

> **Known defect in the naming, stated rather than silently repaired.** `step0b` / `step0c`
> belong to the *measurement* thread; `step0d`, `step1`, `step2`, `step3` belong to the
> *access* thread. They are not a sequence — `step0c` does not lead to `step0d`. The files
> are **not** being renamed, because the board, the commit messages and the operator's own
> Telegram all cite these names, and a rename would break every existing reference to save a
> reader one paragraph. This paragraph is that paragraph.

---

## Thread A — MEASUREMENT: what can be read off the wire, from where

The question: *which exchanges actually sell this device's traffic, and what does the
demand side hand over about it for free?* Everything here is observation.

| document | what it decides |
|---|---|
| `step0-method-2026-08-15.md` | how `(app → exchange)` is observed from the device side at zero cost, and what was rejected (monitor-mode capture, outright) |
| `research-device-capture-2026-08-15.md` | the instrument survey behind that choice |
| `research-exchange-domains-2026-08-15.md` | hostname → org → does it run an OpenRTB exchange |
| `step0-prebid-vantage-2026-08-16.md` | the first header-bidding sweep, and what our VANTAGE forbids us to conclude from it |
| `step0b-ru-vantage-note3-2026-08-18.md` | building a Russian-resident mobile vantage from a rooted phone's tether; the vantage as a live cause of one bidder's silence |
| **`step0c-hb-ru-market-study-2026-08-18.md`** | **the study's design**: the sampling frame, the schedule, the pairing, the channel confound. Start here for the current work |
| `step0d-hb-first-cell-2026-08-19.md` | the first frame built and the first paired cell run under that design — and the three places the design did not survive contact. **§9**: the first cell on the *canonical* frame, and why its own `pairs=14` was four pairs short |
| **`step0e-bid-request-audience-2026-08-20.md`** | **what the bid request actually SAYS about the person.** The payload was never read before this: schemas 1–5 parsed it and threw it away, and schema 6 kept the body while cutting every GET exchange's query string at 500 chars. Answers the README's thesis with a refinement and a negative — 18 of 23 bid-request paths carry no audience field at all, and the identity that does travel is the wrapper operator's |
| `prebid-local-runbook.md` | standing up prebid locally, and what counts as the artifact |
| `REPRODUCE.md` | **the commands.** How to re-run any of the above and how to check the gates |
| **`step0f-adfox-cards-2026-08-21.md`** | **the operator's «AdFox раскрывает все карты» tested against the wire**, not against a draft. Prior art leads: Yandex documents the visibility as a publisher feature, so only what survives that citation is claimed as ours. The bidder ROSTER holds; the PRICE claim breaks — and a price turns out to be a function of the vantage it was asked from |
| **`step0g-loyalty-label-2026-08-21.md`** | **(ru) `userLoyaltyGroups` is a segment NAME and a first-visit DEFAULT, not a verdict about a person.** The operator read the finding as «на сайтах видно, насколько ты лоялен власти» and asked directly whether he had it wrong. He had — and the thing that shows it was already in our own capture |
| `DATASHEET-ru-hb-frame.md` | the datasheet for the published frame + rejection ledger, following *Datasheets for Datasets* (Gebru et al.) — what the dataset excludes, stated as content rather than caveat, because a dataset nobody can bound cannot be cited |

## Thread B — ACCESS: how a person or a small buyer gets in legally

The question: *what can be demanded, from whom, on what legal footing — and what does it
cost a small buyer to get a seat?* Everything here is procedure, correspondence and law.

| document | what it decides |
|---|---|
| `step0d-letter-templates-2026-08-15.md` | three data-subject request letters on three different legal footings |
| `step0d-recipients-eu-entity-2026-08-15-ru.md` | which exchanges have an EU entity to address, and what to lean on for each |
| `step1-permission-matrix-2026-08-17.md` | what the phone stops saying about itself once the user refuses |
| `step2-questions-exchange-2026-08-15-ru.md` | what to ask an exchange, and why each question |
| `step3-small-buyer-onboarding-2026-08-15-ru.md` | who onboards a SMALL buyer, read from primary sources |

## The fact base, and the operator's own instructions

| document | what it is |
|---|---|
| `brief-2026-08-15-operator.md` | **the fact base.** His brief, verbatim, never paraphrased. `⚠` marks a gap only he can close |
| `instruction-capture-2026-08-15-ru.md` | how HE runs the traffic capture on his own phone |
| `instruction-apk-install-2026-08-15-ru.md` | how HE installs the probe |
| `instruction-router-readonly-2026-08-15-ru.md` | what the home router already sees — **read-only, never reconfigured** |
| `method-planted-invariant-oracle-2026-08-15-operator.{txt,pdf}` | the planted-invariant oracle method, as given |

---

## Drafts written for readers outside this repository

| document | what it is |
|---|---|
| `draft-devto-2026-08-20-audience-profile.md` | a dev.to draft of the `step0e` result for a general audience — *"I went looking for my ad profile inside the bid request. It wasn't there."* A DRAFT: the `pub` channel owns dev.to, and nothing here is posted by the measurement thread |

## Reading order, if you are new

1. `README.md` in the root — the question and where it stands.
2. `step0c-hb-ru-market-study-2026-08-18.md` — the design. It is long because it is the
   place where every decision that could bias a result is written down before the result.
3. `step0d-hb-first-cell-2026-08-19.md` — what happened when that design met the network.
4. `REPRODUCE.md` — then do it yourself and disagree with us.

## Two conventions that run through every document here

**A number never travels without its `n` and its coverage.** Not as a footnote — in the same
line. The RU arm is small and saying so is not modesty, it is the only thing that stops a
chart from claiming more than the data.

**Absence is split by cause, always.** `no-bid` (they answered "no"), `no-answer` (they were
asked and said nothing), `blocked` (we never saw the page), `error` (no usable response) and
`no-wrapper` (the page carried nothing to ask with) are five different facts. Folding any of
them together converts our own blindness into a finding about the market, which is the one
mistake this project cannot afford to make quietly.
