# Live boundary refresh — 2026-09-19 20:37 UTC

## Verdict

**UNKNOWN / blocked before capture.** The approved `/run/netns/ru-mobile` namespace is
absent, so the route and egress predicate cannot be evaluated. The existing study remains
pilot-only: **239/4480 readable pairs (5.3%)**, with **901/8960 attempted loads (10.1%)**.

## Fresh live evidence

- `mesh-dash --once adint` — exit **0**; `pane:adint` rendered
  `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480` and showed this
  task as the active `HOLD` row.
- `test -e /run/netns/ru-mobile` — exit **1**; namespace absent.
- `python3 tools/adint-status --include-private` — exit **0**; regenerated status at
  20:37 UTC and reported the same 239/4480 readable-pair pilot boundary.
- `python3 tools/adint-study --test` — exit **0**, `selftest: PASS`.

## Boundary and retry edge

No capture, router write, namespace provisioning, routing change, external contact,
onboarding, payment, or GAID reset was performed. Retry only after the approved namespace
exists; then run:

```text
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Verify RU-mobile egress before any canonical paired measurement. SSH reachability is not
treated as egress evidence.

## Delegation and ownership

No subagent was launched: this was a tiny, tightly coupled live-state refresh whose task
claim, board visibility, artifact, and settlement had to remain in this window. The task was
created, dispatch-checked with exit 0, and claimed by `adint` before the commands above.
