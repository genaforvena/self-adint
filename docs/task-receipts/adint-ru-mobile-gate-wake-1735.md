# RU-mobile gate wake — 2026-09-19 17:36 UTC

Status: **UNKNOWN / held before capture**.

## Fresh live check

- `test -e /run/netns/ru-mobile` — exit **1**: the approved RU-mobile namespace is absent.
- `mesh-dash --once adint` — exit **0**; `pane:adint` rendered `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480`.
- Because the namespace predicate failed, `nsenter`, route probing, capture, paired-run, and all routing mutation were not attempted.

## Boundary and next edge

The existing pilot remains **239/4480 readable pairs**, not a completed market sample. SSH reachability is not treated as RU-mobile egress evidence. Retry this gate only after the approved `/run/netns/ru-mobile` namespace is restored; then verify `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` before the canonical paired run.

No capture, router write, namespace provisioning, routing change, external contact, onboarding, payment, or GAID reset occurred.

## Delegation evidence

The read-only worker review at `/tmp/adint-hb-review.md` was personally inspected. It independently confirmed the absent namespace, the `239/4480` pilot boundary, the canonical input hashes, and the same retry edge; it made no repository or mesh mutation.
