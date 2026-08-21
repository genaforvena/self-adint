# The refusals were the finding, and the pair gate was throwing them away

*2026-08-21. Exploratory half. `tools/adint-holdout --explore`, field `access_contrast`.*

## Two findings, and the robust one was an exclusion

The study's headline contrast is about **price**: does a bidder answer with a number more
often from a domestic vantage. On the exploratory half that rests on **twelve** priced
rows across three or four sites. It is the rarest thing in the corpus.

Underneath it sits a second contrast that is recorded on **every load**: does the
publisher's server answer the vantage at all. Until today that lived entirely in an
exclusion bucket. `pair-gate:refused` is the correct exclusion for the price contrast —
there is no auction to compare on a page one arm never saw — and being right about that is
exactly how it stayed invisible as a measurement.

| arm | refused / attempted | rate | sites refusing it |
|---|---|---|---|
| ru-mobile | **0 / 82** | 0.000 | — |
| nl-direct | 9 / 53 | 0.170 | gismeteo.ru, hh.ru, iz.ru, rbc.ru |
| us-exit | 10 / 82 | 0.122 | iz.ru, rbc.ru |

The domestic arm is never refused, by any of the fourteen sites, on any load. Two
independent foreign exits — different countries, different ASNs — are refused by four.
`rbc.ru` refuses both foreign arms on every single attempt (401) while serving the
domestic one; `iz.ru` refuses both intermittently (403), which is a different mechanism
from a blanket geo-block and worth keeping distinct.

## The p that looks decisive is the wrong unit

Per load, one-sided Fisher gives 1.4 × 10⁻⁴ against nl-direct and 7.3 × 10⁻⁴ against
us-exit. **Do not quote those.** The same fourteen sites are drawn again every replicate,
within a cell and often within one browser generation; the loads are not independent
draws, and a test over them treats correlated repetition as fresh evidence. It will look
decisive no matter what.

The unit the design supports is the **site**. Concordant sites carry no information about
the contrast, so this is a sign test over discordant ones:

- 14 sites, **4 discordant**, all 4 refusing the foreign class only, 0 refusing domestic only
- one-sided exact p = 2⁻⁴ = **0.0625**

Not significant, and total-looking asymmetry cannot rescue it: with four discordant sites
0.0625 is the smallest p the test can produce. The pattern is clean, it is consistent
across two unrelated exits, and it is **unconfirmed**. Both figures are published together
with the inflation named, because reporting only the site level would hide how the count
was reached, and reporting only the load level would be a lie about the evidence.

## And a second holdout leak, from the fix for the first

Adding this contrast leaked the holdout again — the same day, through the mechanism that
had just been hardened against exactly this.

Leak #1 was fixed by making `cell_rows()` refuse a call with no half. That bought nothing
here, because `_report()` **defaulted `half=None`**, and `None` means *both halves* one
layer down. `access_contrast` straddled the split on the very commit that hardened the
accessor, and printed all-cells refusal counts (nl 40/238, us 19/136, ru 0/231) beside the
exploratory ones, from which the holdout's follow by subtraction.

**A permissive default in a CALLER re-opens the door the callee just locked.** Locking one
function is not locking a path. `_report()` now refuses too, and the gate asserts both the
refusal and that a half-scoped contrast actually counts fewer loads than a straddling one —
the second half matters, because a refusal that is never reached proves nothing about the
numbers that got printed.

Recorded as the second entry in `ref/HOLDOUT-LEAKS.jsonl`. The holdout is still not spent,
and is now partially compromised on **both** contrasts. It was the access contrast that
leak #1 had left clean; that is no longer true, and it is worth saying plainly rather than
discovering later.
