# Live-state trace — 2026-09-12

## Source path

`mesh-dash`'s `adint` renderer reads `~/.mesh/adint-step` into `_astep` and prints
that text with its file age (`/home/mesh-home/.local/bin/mesh-dash`, around line 3588).
The live file currently says:

> step0d-hb-first-cell — frame built (2 admits/69, density 2.9%), first paired cell captured (6 pairs); walk continuing toward 20 admits

So `2/69` is hand-maintained step text, not a live count computed by the dashboard.

## Check against project evidence

The project documentation and tools do not define what the `69` denominator counts.
The checked-in `step0d` report has no such denominator; the prior read-only audit
(`live-state-audit-20260912.md`) reproduced current Stage B ledger sizes of 208 RU and
254 NL touched domains, which are separate counts. No relation between those counts
and `69` is established by the available artifacts.

## Disposition

Keep the dashboard text as an unverified historical note. Do not use `2/69` as a
current frame density or revise the study from it. The denominator remains UNKNOWN;
the source artifact or definition that produced `69` is needed before it can be used.
