# RU-mobile namespace gate question — 2026-09-20

Task: `adint-goal-20260920-ru-mobile-namespace-question-2136/ru-mobile-namespace-question-2136`

## Live predicate

Observed at `2026-09-20T21:34:23Z` from the `adint` window:

- `/run/netns/ru-mobile`: **absent**
- approved route check: **NOT_RUN_NAMESPACE_ABSENT**
- current study remains `239/4480` readable pairs; no new paired capture was started

The namespace and route are outside this repository's safe local recovery scope. No namespace,
routing, router, seat, payment, exchange-contact, publication, or GAID mutation was attempted.

## Exact external question

Please restore the already-approved `ru-mobile` network namespace and its approved RU-mobile
egress. Stop after the namespace exists. The next read-only gate is:

```text
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Only after that route check succeeds should a separate task run the canonical paired capture.
Until then the study verdict is `UNKNOWN`, not `NOT`.

## Evidence and delegation

- Live check: `/tmp/adint-namespace-live.txt`
- Delegated read-only scan personally inspected: `/tmp/adint-next-action-scan.md`
- Prior boundary evidence: `docs/task-receipts/adint-live-boundary-preflight-2101.md`
- Prior route-gate evidence: `docs/task-receipts/adint-ru-mobile-gate-preflight-20260920.md`

## Retry edge

When `/run/netns/ru-mobile` exists, rerun the route check above; do not infer readiness from the
namespace file alone, and do not start capture in the same retry.
