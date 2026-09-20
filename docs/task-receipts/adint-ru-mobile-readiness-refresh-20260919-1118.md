# RU-mobile readiness refresh — 2026-09-19 11:20 UTC

Task: `adint-goal-20260919-ru-mobile-readiness-refresh-1118/ru-mobile-readiness-refresh-1118`

## Read-only probe

Commands run from `/home/mesh-home/lte-workstation`:

```text
date -u +%Y-%m-%dT%H:%M:%SZ
2026-09-19T11:20:48Z
ip netns list
(no rows; exit 0)
test -e /run/netns/ru-mobile
false: absent
```

The namespace was absent, so the route probe was correctly not attempted. No namespace,
routing, router, capture, external-contact, onboarding, payment, or GAID-reset action was
performed.

## Current pane observation

`mesh-dash --once adint` at `2026-09-19T11:20:48Z` rendered `pane:adint` as working on
`step0d-hb-first-cell`, with `ru-mobile frame 14 admits/208 domains; 239/4480` and the
undischarged `witness-62079-udev-followthrough/reap-orphan-listene…` hold.

## Verdict and retry edge

Current readiness: **UNKNOWN / not ready**. The approved `/run/netns/ru-mobile` prerequisite
is still absent. Exact retry edge: after that namespace exists, run
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`, then perform the canonical
paired measurement only if the RU-mobile egress is verified. Keep the active collection and
all irreversible branches closed.
