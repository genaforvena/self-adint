# Safe-study loop regression — 2026-09-20 07:10 UTC

Task: `adint-goal-20260920-safe-study-loop-regression/safe-study-loop-regression`
Owner: `adint`

## Verification

- `python3 tools/adint-study --test` — exit 0, `selftest: PASS`.
- `python3 tools/adint-study` — exit 0. States remain explicit: `MEASURED` (Step 0
  device, Step 1 receiver, Step 4 web), `DRAFT_ONLY` (Step 2),
  `GATED_OPERATOR_GO` (Step 3), `BLOCKED_OPERATOR_EXPORT` (Step 4 device),
  `NOT_ELIGIBLE` (Step 5), and `NO-BASIS` (Step 6).
- `python3 tools/adint-status --test` — exit 0, all checks passed (`selftest: ok`).
- `python3 tools/adint-status --include-private` — exit 0; regenerated the on-disk
  report. The canonical `ru-mobile` frame remains 209 domains touched, 14 admits,
  and 239/4480 readable pairs (5.3%). This is still a pilot, not a market estimate.
- `mesh-dash --once adint` — exit 0; observed on `pane:adint` at
  `2026-09-20T07:10:55Z`: `step0d-hb-first-cell — ru-mobile frame 14 admits/208
  domains; 239/4480`.
- `test -e /run/netns/ru-mobile` — exit 1 (`namespace=absent`); no route probe,
  capture, namespace provisioning, router write, or external action was performed.

## Inputs and worktree evidence

HEAD at verification: `097b40f750aac248b4631a6364fea31708ed886d`.
Pre-existing dirty worktree changes were preserved; this receipt does not claim them.

SHA-256:

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf3  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
```

## Verdict and retry edge

The safe local loop is green, but the paired RU-mobile gate remains `UNKNOWN` and is
held before capture. When `/run/netns/ru-mobile` exists, run:

```text
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Only a verified RU-mobile route permits the already-approved paired measurement.

## Delegation record

The existing `adint-step0d-audit` Codex worker was reused for a read-only independent
inspection. I personally inspected the cited frame, receipts, and checks above; the
worker report was advisory and did not settle this task.
