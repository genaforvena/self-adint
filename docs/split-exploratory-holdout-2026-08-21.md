# The corpus is split in two before anyone looks at it

**Written 2026-08-21, before the collection window it governs opens**, and checkable:
`tools/adint-holdout --verify` hashes this file against the digest recorded when the split
was declared.

Declared by `adint@mesh-home` for task `adint-two-week-collection-2026-08-21`, replacing
an earlier draft that fixed an expected effect in advance. **That draft is withdrawn**, on
the operator's instruction and on the merits.

---

## 0. Why nothing is predicted here

The operator's position, 2026-08-21: *«гипотезы давайте не делать, я не вижу смысла в
гипотезах… мне интересны данные, и я не хочу, чтобы мы имели какие-то предубеждения
насчёт того, что мы обнаружим… без предубеждения было бы их собирать лучше»*.

This project's own best result is the argument for it. **Nobody predicted the vantage
effect.** It fell out of the paired design — we were not looking for it, we were simply
capturing two arms at once, and the contrast was there when we read the ledgers. An
expectation fixed in advance would not have found it, and worse, an expectation *corrupts
collection*: you look more often where you expect to find, and you are quicker to call an
inconvenient record an outlier.

## 1. What is still true, and it is arithmetic rather than philosophy

The corpus carries **220 distinct request-parameter names, 29 bidders, 14 sites, 2 arms**.
That is thousands of comparisons available to anyone who goes looking. At p < 0.05 roughly
**one in twenty comes back "significant" out of pure noise**. So a test chosen *after*
looking at the data will always find significance, and what it measures is **us** — the
breadth of our search — not the market. This mesh has the finding already in another
shape: *a binary search over a deferred release quota measures the searcher*.

The tension is not between "data" and "predictions". It is between **wanting unlimited
search** and **wanting a trustworthy number**, and those two do not come together.

## 2. The procedure

**The corpus is split into an EXPLORATORY half and a HOLDOUT half, by a rule declared
here, before any analysis.** Declaring the rule matters as much as splitting: an
undeclared split becomes the place the fitting happens.

### 2.1 The split rule

**A cell (one paired run, identified by its `run_id`) goes to the holdout when
`sha256("adint-split-2026-08-21|" + run_id)` has an even first byte, and to the
exploratory half otherwise.**

- **The whole cell moves together.** A pair must never straddle the halves — the two arms
  of a pair are one observation, and splitting them would leak the holdout into the
  exploratory half through the pairing itself.
- **By cell, not by time.** A time split would make the holdout a different period, and
  the whole point of the fortnight is to add hours, dayparts and weekdays we do not have
  yet. With a time split, ordinary temporal drift would read as a failed confirmation.
- **Deterministic and stable.** A cell's half is a function of its id alone, so it cannot
  change as the corpus grows, and it cannot be nudged.
- Cells already on disk before this file are assigned by the same rule. They were captured
  without knowledge of it, which is exactly the property that makes the assignment safe.

### 2.2 On the exploratory half: no rules at all

Turn it any way you like. Follow whatever is curious. Compute anything, as often as you
want, with no correction and no plan. This half exists to be mined, and this is precisely
what the operator asked for. Nothing found here needs to have been anticipated.

### 2.3 On the holdout: exactly one look

Anything the exploratory half suggests is checked against the holdout, **which has not
been looked at before**.

**The holdout is computed ONCE.** This is the rule that is easy to break by accident and
worthless once broken:

- A look — any look, including one that "was just a sanity check" — **spends** it.
- **A look is RECORDED, not quietly continued.** `tools/adint-holdout --confirm` appends
  every read to `ref/HOLDOUT-READS.jsonl`, which is append-only and committed. After a
  read, the tool reports the holdout as **spent** and every later read as what it now is:
  more exploration.
- Related failure this mesh already owns: *a null computed inside the selection window is
  not a holdout*. A number computed on data you used to choose what to compute is not a
  check on that choice.

### 2.4 The language of reports — enforced, not left to care

While an effect lives only on the exploratory half it is **"interesting, unconfirmed"**.
It is **never** called significant, and no p-value from that half is presented as one.
`tools/adint-holdout` labels every exploratory figure `EXPLORATORY — not confirmed` in its
own output, so the wording does not depend on who is writing the sentence that day.

## 3. What does not change

The three collection rules stand exactly as they were:

- **An arm that could not be reached renders `blocked`, NEVER `no-bid`.** The RU arm rides
  a fragile single tether. A fortnight that recorded our own blindness as the market's
  silence would produce a confident, publishable, false result — the one error this
  project cannot afford to make quietly. Gated in `tools/adint-holdout`'s analysis set and
  in `tools/adint-paired-run`'s vantage refusal.
- **Ledgers are append-only and kept whole for the whole window.** No sliding window, no
  pruning. A pruned ledger makes every `n` unreproducible and moves it DOWN as well as up.
  `tools/adint-collect-cell` fails the tick if the corpus shrinks.
- **Every run publishes its coverage, including the runs that collected nothing.** A night
  that collected nothing must look like zero rows with a reason, never like a clean run.

---

*Amendments after the window opens are recorded as dated amendments and must never
silently alter §2. `--verify` reports a hash mismatch, which is the point.*
