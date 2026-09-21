# adint live boundary preflight — 2026-09-20 21:04 UTC

Task: `adint-goal-20260920-live-boundary-preflight-2101/live-boundary-preflight-2101`

## Result

The Step 0d RU-mobile boundary remains **UNKNOWN / held before capture**. The approved
`/run/netns/ru-mobile` namespace is absent, so the namespace route probe was correctly skipped.
No capture, router write, external probe, namespace creation, or irreversible action occurred.

## Live evidence

- `python3 tools/adint-status --include-private` — exit 0; canonical frame: 209 domains touched,
  14 admits; pilot coverage 239/4480 readable pairs (5.3%) and 901/8960 attempted loads (10.1%).
- `mesh-dash --once adint` — exit 0; `pane:adint` rendered at `2026-09-20T21:04:35Z`:
  `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480`.
- `test -e /run/netns/ru-mobile` — exit 1; namespace absent.
- `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` — skipped because the
  namespace predicate was false (`ROUTE=SKIPPED_NAMESPACE_ABSENT`).

## Retry edge

After the approved `/run/netns/ru-mobile` exists, rerun the namespace route lookup first; only
then consider the paired measurement. Until that event, preserve UNKNOWN and do not create the
namespace or alter routing from this lane.

## Scope and delegation

This was a read-only local preflight. The independent audit was delegated to
`adint-live-audit-20260920`; its artifact must be inspected separately before any conclusion
about duplicate work or next action.
