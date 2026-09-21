# Report freshness reconciliation — 2026-09-21

## Result

The checked-in `report.html` is stale: it remains mtime `2026-09-10 12:55:22 UTC` and carries
the marker `generated 2026-09-10 12:55 UTC`. A read-only status regeneration completed at
`2026-09-21 09:04 UTC` and printed the current disk-backed corpus to stdout, but did not rewrite
`report.html`. The measured Step 0d boundary remains a pilot: `239/4480` readable pairs
(`5.3%`), `901/8960` attempted loads (`10.1%`). No capture, routing, router, or external
mutation ran.

## Commands and evidence

| UTC | command/result |
|---|---|
| 09:04:15 | `stat -c 'mtime=%y size=%s path=%n' report.html`; before refresh: mtime `2026-09-10 12:55:22 UTC`, size `47442`; HTML marker `generated 2026-09-10 12:55 UTC` |
| 09:04:15 | `python3 tools/adint-status --test`; exit `0`, `selftest: ok` |
| 09:04:15 | `python3 tools/adint-status --include-private`; exit `0`, stdout status generated `2026-09-21 09:04 UTC`; `report.html` remained unchanged |
| 09:04:15 | `ip netns list`; exit `0`, no namespaces; `test -e /run/netns/ru-mobile`; exit `1` (absent) |
| 09:04:24 | `mesh-dash --once adint`; exit `0`; `pane:adint` rendered `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480` |
| 09:04:26 | `mesh-task queue --dispatch --owner adint`; exit `0`, no eligible owner row |

## Duplicate and blocked-task check

The exact queued rows `routing-shadow-resolver-followup-adint-20260914/review-frozen-gate-after-parent-completion`
and `routing-shadow-11fcb-followup-adint-20260914/reconcile-11fcb-after-existing-review` remain
waiting for `self-review-routing-shadow-20260914/independently-evaluate-routing-shadow`; the witness
gate is frozen until `2026-09-28T16:52:06Z` or terminal `INCONCLUSIVE` by `2026-10-14T16:52:06Z`.
The delegated read-only audit at
`/home/mesh-home/.mesh/evidence/adint/delegation-audit-20260921.md` was personally inspected and
confirmed that paired capture, namespace recovery, and prior status-refresh actions are already
represented or blocked. This task is limited to the distinct stale-report freshness gap.

## Retry edge

The next substantive collection action remains blocked/UNKNOWN until `/run/netns/ru-mobile` exists
and its RU-mobile route/egress is verified. At that point, run the canonical paired capture and
4480-pair reconciliation; do not infer a negative result from namespace absence.
