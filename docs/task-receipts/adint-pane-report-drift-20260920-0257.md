# Step 0d pane/report drift reconciliation — 2026-09-20 03:00 UTC

Task: `adint-goal-20260920-pane-report-drift-0257/pane-report-drift-0257`
Owner: `adint`

## Delegation

No subagent was launched. This is one tightly coupled read-only comparison: the acceptance
depends on one report regeneration, the same canonical input hashes, and the already-rendered
`pane:adint` frame. Splitting those reads could compare different snapshots. I personally
inspected the command outputs and the pane observation.

## Commands and results

- `python3 tools/adint-study --test` — exit `0`, `selftest: PASS`.
- `python3 tools/adint-study` — exit `0`; emitted `MEASURED` for Step 0 device, Step 1
  receiver, and Step 4 web; `DRAFT_ONLY` for Step 2; `GATED_OPERATOR_GO` for Step 3;
  `BLOCKED_OPERATOR_EXPORT` for Step 4 device; `NOT_ELIGIBLE` for Step 5; `NO-BASIS` for
  Step 6.
- `python3 tools/adint-status --include-private` — exit `0`; regenerated `report.html`.
  The current report says `ru-mobile`: **209 domains touched**, **14 admits**, and
  **239/4480 readable pairs (5.3%)**; attempted-load coverage is **901/8960 (10.1%)**.
- Namespace predicate: `[ -e /run/netns/ru-mobile ]` — `absent`; no route probe or capture
  was attempted.
- The requested `mesh-dash --once adint` read at `2026-09-20T02:57:16Z` rendered on
  `pane:adint`: `step0d-hb-first-cell — 14 admits/208 domains; 239/4480 readable pairs`.
  A later repeat timed out under concurrent mesh-ledger load, so it is not treated as a fresh
  pane reading.

## Drift classification

The `209` in the freshly regenerated private report and the `208` rendered on `pane:adint`
are **presentation/history drift**, not a measured-pair change: admits and readable pairs
agree exactly (`14`, `239/4480`). The current canonical report is the on-disk source of truth;
the pane's domain total is stale until its next successful refresh.

Canonical hashes personally verified after regeneration:

```
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
```

## Verdict and retry edge

The safe Step 0d boundary remains **UNKNOWN / held before capture** because the approved
`ru-mobile` namespace is absent. The next safe action is to rerun the status/pane reconciliation
after that namespace exists, then verify
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` before any paired measurement.
The frozen routing-shadow comparison remains held until `2026-09-28T16:52:06Z` and at least
100 eligible tasks. No capture, namespace, routing, router, external-contact, publication,
onboarding, payment, or GAID action was performed.
