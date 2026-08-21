# The holdout leaked, and the ledger stayed empty

*2026-08-21. Recorded in `ref/HOLDOUT-LEAKS.jsonl`. The holdout is not spent; it is
partially compromised, and every future confirmation must be reported carrying this.*

## What happened

While measuring why the RU arm loses 32 % of its renders, an ad-hoc analysis script
tallied priced rows by arm and load-outcome over **every cell on disk**, with no half
filter. The table it printed:

| arm | priced rows, all halves |
|---|---|
| ru-mobile | 26 |
| nl-direct | 13 |
| us-exit | 5 |

The exploratory half's own counts were already known — 11 surviving plus 2 gate-dropped
for ru-mobile, 2 (both gate-dropped) for nl-direct, 1 plus 2 for us-exit. Subtraction is
not a tool you can be denied. The holdout holds roughly **ru 15 / nl 11 / us 4**.

That is not a neutral fact. The exploratory half shows nl-direct with **zero** priced rows
surviving the pair gate and only two in existence; the leaked arithmetic says the holdout
holds about eleven. It points, before any confirmation, at the gap being weaker there —
which is precisely the pressure to quietly re-specify the test that the split exists to
remove.

What was **not** seen, stated so the damage is bounded rather than assumed total: no test
statistic, no p-value, no pairing, no per-site or per-cell breakdown of the holdout, and
no indication whether those rows survive the pair gate there.

## The guard was on the door, and the wall was missing

`--confirm` is well defended: it refuses without a `--claim`, it refuses against a changed
declaration, and it writes its ledger row **before** it prints anything, so a crash cannot
swallow the read. All of that held. `ref/HOLDOUT-READS.jsonl` is still empty and the
holdout is formally unread.

None of it was ever in the path. `cell_rows()` iterated every cell on disk by default, and
a five-line script read straight past the protocol without touching it. **A protocol
enforced only at the entrance a careful person uses is enforced against care, not against
accident** — and the accident here was mine, in the same session that wrote the gate for
the last one.

## The fix is in the corpus, not the door

`cell_rows()` now has **no default half**. It takes `"exploratory"`, `"holdout"`, or an
explicit `None` that says in the source that straddling both halves is what the author
meant, and raises `TypeError` otherwise. Two callers legitimately pass `None`:
`split_table()`, and `analysis_set()`, which must see the other half in order to *count*
it as a `wrong-half` exclusion — an exclusion count is a bias term and deleting it would
be its own defect.

Three gates, each seen red before green: an unfiltered call refuses; a half-scoped read
returns that half and nothing else (the two run-id sets are disjoint); and `None` still
sees both. The redundant inner half-check in `render_gate_table` was **deleted** rather
than kept as belt-and-braces — two guards for one fact means the live one is never the one
under test.

## What this costs

The holdout is worth less than it was this morning, and saying so is the only way any of
it survives. A confirmation reported as a clean first look would be worth nothing at all,
because it would be false. When there is finally one claim worth spending it on, the
report goes out with this file attached and the reader decides what the number is worth.

The exploratory half is untouched by this and remains free to mine.
