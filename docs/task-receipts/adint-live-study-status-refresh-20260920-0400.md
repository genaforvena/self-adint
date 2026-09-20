# Live study status refresh — 2026-09-20 04:00 UTC

Task: `adint-goal-20260920-live-study-status-refresh-0358/live-study-status-refresh-0358`
Owner: `adint`

## Fresh read-only evidence

- `python3 tools/adint-study --test` — exit 0; `selftest: PASS`.
- `python3 tools/adint-study` — exit 0; explicit states remain `MEASURED` (Step 0 device,
  Step 1 receiver, Step 4 web), `DRAFT_ONLY` (Step 2), `GATED_OPERATOR_GO` (Step 3),
  `BLOCKED_OPERATOR_EXPORT` (Step 4 device), `NOT_ELIGIBLE` (Step 5), and `NO-BASIS` (Step 6).
- `timeout 120 python3 tools/adint-status --include-private` — exit 0; report generated at
  `2026-09-20 04:00 UTC`.
- `mesh-dash --once adint` — exit 0; `pane:adint` rendered at
  `2026-09-20T04:00:42Z`: `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains;
  239/4480 readable pairs`.
- The status report records `239/4480` readable pairs (5.3%) and `901/8960` attempted loads
  (10.1%); it remains a pilot, not a Russian-market sample. The report's current frame summary
  says 209 domains touched, with 14 admitted.

## Artifact hashes

```text
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
```

## Verdict and next edge

The safe local status lane is refreshed and remains `UNKNOWN / held before capture` for the
paired RU-mobile gate. No capture, external probe, router write, exchange contact, publication,
or irreversible action was performed. The next safe edge is to rerun the local report after the
approved `/run/netns/ru-mobile` namespace exists, then verify
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` before any paired measurement.

## Delegation and inspection record

The existing `adint-ledger-audit` coding-agent worker was reused for a non-mutating independent
audit request. I personally inspected its event trail at `/tmp/csd-workers/homes/adint-ledger-audit`
and the completed receipt/report artifacts above; its audit was advisory only and did not settle
the task. Task ownership, execution, receipt writing, and final verification remained in this pane.
