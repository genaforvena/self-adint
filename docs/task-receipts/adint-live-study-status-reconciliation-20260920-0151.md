# Live study status reconciliation — 2026-09-20 01:51 UTC

Task: `adint-goal-20260920-live-study-status-reconciliation-0148/live-study-status-reconciliation`
Owner: `adint`

## Fresh read-only evidence

- `python3 tools/adint-study --test` — `selftest: PASS`.
- `python3 tools/adint-study` regenerated the disk-derived status report. The current states
  are `MEASURED` (Step 0 device, Step 1 receiver, Step 4 web), `DRAFT_ONLY` (Step 2),
  `GATED_OPERATOR_GO` (Step 3), `BLOCKED_OPERATOR_EXPORT` (Step 4 device), `NOT_ELIGIBLE`
  (Step 5), and `NO-BASIS` (Step 6).
- `python3 tools/adint-status --include-private` regenerated the canonical report at
  `2026-09-20 01:51 UTC`. The measured frame remains `ru-mobile`, 14 admits / 209 domains,
  with `239 / 4480` readable pairs (5.3%) and `901 / 8960` attempted loads (10.1%). This is
  still a pilot, not a Russian-market sample.
- Namespace predicate: `/run/netns/ru-mobile` is absent; no route probe or capture was attempted.
- Hashes: `report.html` `b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f`;
  `ref/CANONICAL-FRAME` `05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf`;
  `ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv`
  `b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3`;
  `public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl`
  `8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92`.

## Verdict and next edge

The safe local status lane is refreshed and remains **UNKNOWN / held before capture** for the
paired RU-mobile gate because the approved namespace is absent. The next safe edge is to rerun
the local report after the namespace exists, then verify
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` before any paired measurement.
The frozen routing-shadow comparison remains held until its existing 2026-09-28 / 100-eligible-
task gate.

## Delegation decision

No subagent was launched. This was one tightly coupled read-only reconciliation whose acceptance
depends on matching the same report generation, file hashes, namespace predicate, and pane frame;
splitting it would risk mismatched timestamps. I personally inspected all command output and the
fresh `mesh-dash --once adint` frame at `2026-09-20T01:51:55Z`.
