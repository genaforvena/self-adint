# RU-mobile gate preflight — 2026-09-20

UTC observation: `2026-09-20T20:32:27Z`.

## Commands and results

- `mesh-dash --once adint` exited `0`.
- `pane:adint` rendered the active Step-0D signal: `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480`.
- The live pilot boundary remains `239/4480` readable pairs; this is not a completed market sample.
- Read-only namespace check: `/run/netns/ru-mobile` is absent. `ip netns list` exited `0` and returned no namespace.
- No `nsenter` route lookup was attempted because the namespace file is absent.

## Disposition

The RU-mobile route and paired capture remain `UNKNOWN`; no capture, routing change, router write, external probe, publication, onboarding, payment, or GAID reset was performed.

Retry edge: after the approved `/run/netns/ru-mobile` namespace is restored, run `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` first; only a successful route read permits the canonical paired run.

Task: `adint-goal-20260920-ru-mobile-gate-preflight/ru-mobile-gate-preflight-20260920`.
