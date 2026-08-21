# What widening the frame would cost, and what it would buy

*2026-08-21. A costed decision for `ref/CANONICAL-FRAME`. **Not acted on** — the frame is a
study declaration, and changing it mid-collection is not a code change.*

## Why this is on the table

The access finding is **frame-limited**, not time-limited. Its unit is the site, because a
refusal is a site's policy toward a class of vantage and every extra cell re-observes the
same policies rather than drawing new ones. With 14 comparable sites and 4 refusing the
foreign class, the exact one-sided p is **0.0625**, and thirteen more days of cadence will
not move it. Only more sites will.

The price contrast is the opposite: its unit is the pair, and it genuinely wants the
fortnight.

## The measured budget

Median **109.8 s per pair**, over the six most recent cells (28 pairs each, 51.8–54.8 min
of pair time, ~58 min wall). Against the 2-hour cadence:

| frame | replicates | pairs/cell | pair time | verdict |
|---|---|---|---|---|
| 14 | 2 | 28 | 51.3 min | **current** |
| 18 | 2 | 36 | 65.9 min | fits |
| 21 | 2 | 42 | 76.9 min | fits, ~35 min margin on wall time |
| 28 | 1 | 28 | 51.3 min | fits — doubles the frame at today's cost |
| 28 | 2 | 56 | 102.5 min | **too tight** — collides with the next tick once overhead is counted |

## The actual trade

**Sites buy access evidence; replicates buy price evidence.** They compete for the same
two-hour budget, and the two findings want opposite things:

- `28 × 1` doubles the frame for free in time — and halves the pairs per cell, which is
  exactly what the price contrast is short of.
- `21 × 2` is the only option that improves *both*: +50 % frame and +50 % pairs, at 77 min
  of pair time.

If the observed 4/14 propensity holds among added sites — an assumption, and the added
sites would come from the same `ref/ru-universe-tranco` pool the current ones did — the
access p goes:

| frame | expected discordant | p |
|---|---|---|
| 14 | 4 | 0.0625 |
| 18 | 5 | 0.031 |
| 21 | 6 | 0.016 |
| 28 | 8 | 0.0039 |

18 sites is where it crosses 0.05.

## What it costs that is not time

Changing the frame mid-collection makes the corpus **heterogeneous**: early cells carry 14
sites and later ones more, so every aggregate over sites spans a changing frame, and the
exploratory/holdout split — which moves whole cells — puts a mix on both sides. That is a
real cost and it is not paid back by a smaller p.

Three ways to handle it, in decreasing order of rigour: restart the fortnight on the wider
frame; run the wider frame as a **separate declared study** with its own window; or widen
now and report every site-level aggregate twice, once on the 14-site core and once on the
full frame, with the core as the headline.

## The recommendation, which is not a decision

`21 × 2` if the frame changes at all, and **only** with the third mitigation above — the
14-site core stays the headline and the wider frame is reported beside it, so nothing that
has already been collected is retroactively reinterpreted.

Doing nothing is a defensible answer. The access result is real, consistent across two
unrelated exits, and honestly reported at p = 0.0625; a fortnight's price contrast may be
the more valuable thing to protect. **The decision belongs in `ref/CANONICAL-FRAME`, where
the frame is declared, and to the operator — not to whichever mind happened to notice the
constraint.**
