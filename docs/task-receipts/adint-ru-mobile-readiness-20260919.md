# RU-mobile readiness — 2026-09-19

Task: `adint-goal-20260919-ru-mobile-readiness/verify-ru-mobile-readiness`

At 2026-09-19 00:19 UTC, the live readiness check reported:

- `ip netns list`: no `ru-mobile` namespace;
- `/run/netns/ru-mobile`: absent;
- egress probe: not run because the namespace prerequisite was absent.

Verdict: **NOT READY**. Do not start the paired collection cycle.

Exact retry edge: after the approved `/run/netns/ru-mobile` exists, verify its RU-mobile egress with `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`, then run the canonical paired measurement with both-arm vantage verification. No collection, external contact, router write, seat onboarding, payment, or GAID reset was performed.

Fresh verification at 2026-09-19 00:23:02 UTC still reports `ip netns list` without `ru-mobile` and `/run/netns/ru-mobile` absent. The task dispatch check exited 0 and the exact-owner claim subsequently succeeded at 00:22:51 UTC after the contention cleared; the ledger row is active.

Delegation: none; this was one tightly coupled read-only prerequisite measurement. Personally inspected the live `mesh-dash --once adint` stream and the prior receipt `adint-next-cycle-gate-20260918.md` before acting.
