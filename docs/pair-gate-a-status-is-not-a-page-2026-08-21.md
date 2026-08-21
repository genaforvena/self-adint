# A status is not a page: the pair gate accepted a refusal as a rendered arm

*2026-08-21. Exploratory half. `tools/adint-holdout`, commit below.*

## What was wrong

`analysis_set()` decided whether an arm had rendered its page with

```python
ok = r.get("http_status") is not None and bool(r.get("vantage_verified"))
```

which reads as *something came back*, and was used to mean *the arm rendered*. Those are
different claims, and one live site separates them cleanly.

`rbc.ru` answers **HTTP 401 to both foreign exits and 200 to the domestic one** — the
vantage-conditional access this study exists to measure. Under the predicate above the
401 counted as a rendered arm, so the pair passed the gate. The served arm then
contributed its whole request set against a partner that structurally cannot contribute
any:

| arm | `rbc.ru` load rows | entries reaching the contrast |
|---|---|---|
| ru-mobile | 20 no-load, 5× **200** | **1111** |
| nl-direct | 16× **401** | 0 |
| us-exit | 10× **401** | 0 |

`ref/CANONICAL-FRAME` had already decided to keep `rbc.ru` and recorded it as *a
structural zero for the PAIR unit* — and at the pair unit it is exactly that, harmless.
The pooled Fisher contrast is at the **entry** unit, where the zero does not cancel: it
inflated one arm's denominator by **21.9 %**. A decision reasoned through at one unit of
analysis does not travel to another on its own.

Direction, stated so nobody has to guess: the inflation was **conservative** for the
effect actually reported. Removing it moves the exploratory one-sided Fisher p from
0.0563 → **0.0246** (vs nl-direct) and 0.0395 → **0.0131** (vs us-exit), on unchanged
priced counts (11 / 0 / 1). The paired permutation is untouched — it was already at the
pair unit. Being wrong in the direction that flatters nobody is still the wrong
statistic.

## The rule was already written down, in the tool next door

`tools/adint-status` has classified 4xx/5xx as `refused` — explicitly **not** rendered —
since it was written, and its docstring names this site's 401 as the case. Two readers of
one corpus, one correct filing, no gate between them, and the published per-site yield
table and the statistical analysis set quietly disagreed about which pairs exist. The fix
is not only the predicate: `--test` now loads `adint-status`'s own `rendered` from source
and fails if the two rules disagree on any of ten statuses, and fails *louder* if
`adint-status` cannot be read at all — an import guard that skips its own gate would
report a rule nobody could compare as a rule nobody disagreed with.

## The bucket was named after the rarest of its causes

Every pair the gate dropped landed in one exclusion bucket called `vantage-unproven`,
which sends a reader to fix the vantage proof. On the live corpus that name was wrong for
**92 %** of what it held. Splitting it:

| ru-mobile load rows | n | share |
|---|---|---|
| ok | 215 | 0.621 |
| **no-load, vantage proof PASSING** | 112 | 0.324 |
| no-load & unverified | 15 | 0.043 |
| loaded but vantage unverified | 4 | 0.012 |

The instrument was working. The tether was not fetching the page. Three named buckets now
replace the one — `pair-gate:refused` (a fact about the site), `pair-gate:no-load` (a fact
about our link), `vantage-unproven` (a fact about the instrument) — because they send you
to three different fixes.

## The gate is 8× more lossy on one arm, and that now travels with the numbers

| arm | rendered | cleared the gate | rate |
|---|---|---|---|
| ru-mobile | 346 | 215 | **0.621** |
| nl-direct | 223 | 210 | 0.942 |
| us-exit | 141 | 136 | 0.965 |

Every surviving pair is conditioned on the fragile tether having been up at that moment,
and the conditioning is eight times stronger on the arm whose behaviour is the headline.
It is also **site-structured, not random** — RU clears the gate on `iz.ru` 0.88 and
`sports.ru` 0.85, against `aif.ru` 0.20, `rbc.ru` 0.20, `matchtv.ru` 0.33 — so the
analysis set is not a random subsample of the frame but a re-weighted one that
under-represents exactly the sites the link struggles with.

That is a statement about the analysis set, not about the market, so `coverage()` now
publishes the per-arm table under `render_gate` in every report. It cannot be quoted away
because it arrives attached to the thing it conditions.

## What is still open

The whole point of the split is that none of this is confirmed. It is the exploratory
half; the holdout is unread and stays unread until there is one claim worth spending it
on. Twelve priced entries across three sites is not yet a finding — it is a reason to let
the fortnight's cadence run.
