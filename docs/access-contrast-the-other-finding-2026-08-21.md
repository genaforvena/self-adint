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

## The p that looks decisive is the wrong unit — twice over

Per load, one-sided Fisher gives 1.4 × 10⁻⁴ against nl-direct and 7.3 × 10⁻⁴ against
us-exit. **Do not quote those.** The same fourteen sites are drawn again every replicate,
within a cell and often within one browser generation; a test over them treats correlated
repetition as fresh evidence and will look decisive no matter what.

The first attempt to fix that was **also wrong, and worse** — it asked, per site, *has this
site ever refused a foreign arm*, and ran a sign test over the sites that differed. That is
not a test. **"Ever refused" latches:** one transient 4xx converts a site to discordant
permanently, so k can only rise as loads accumulate and p = 2⁻ᵏ can only fall. A statistic
that moves one direction with more data reaches any threshold by waiting. On this corpus it
went 2 → 2 → 2 → 2 → 3 → 3 → 4 across seven cells and would have kept going whether or not
the effect were real. It also made the *frame* look like the binding constraint — fourteen
sites, floor p = 2⁻¹⁴ — which was an artefact of the wrong unit.

The unit that works is **(cell, site)**, compared with McNemar. A cell is an independent
browser generation, and the domestic-only count `c` can rise, so the p is free to move both
ways. And **each foreign arm is compared alone**: pooling them asks whether the foreign
*class* was refused, which is an OR over two arms against the domestic arm's one — two
chances versus one, biased toward the finding.

| contrast | units | concordant | foreign-only | domestic-only | one-sided exact p |
|---|---|---|---|---|---|
| nl-direct vs ru-mobile | 20 | 17 | 3 | 0 | **0.125** |
| us-exit vs ru-mobile | 42 | 38 | 4 | 0 | **0.0625** |

Neither is significant. Pooling the two foreign arms would have given 7 discordant units and
p = 0.0078 — three-quarters of that was the two-chances-versus-one asymmetry.

So the honest state of the access finding: the pattern is clean, consistent across two
unrelated exits, and **unconfirmed**. What it needs is more *cells*, not a wider frame —
the corrected unit accrues twelve times a day, and each additional one-sided discordant unit
halves the p.

A note on how the corrected test was found: writing the McNemar p as a `min(b, c)` tail — the
smaller of the two — would have reported **p = 0.0625 for a corpus in which the DOMESTIC arm
was the refused one**, reading as confirmation of the opposite claim. The gate that catches
it feeds the test exactly that inverted corpus and requires the discordance to land on the
domestic side. It was written before the bug and it caught it.

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
