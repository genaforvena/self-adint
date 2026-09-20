# Step 0d pane/report reconciliation — 2026-09-18

Status: **pilot remains open; no capture started**.

## Live evidence

- `mesh-dash --once adint` at 2026-09-18 21:39 UTC rendered `step0d-hb-first-cell`, `14 admits/208 domains`, and `239/4480` readable pairs on `pane:adint`.
- `python3 tools/adint-status --include-private` exited 0 at 2026-09-18 21:39 UTC and regenerated `report.html`.
- The fresh report says `ru-mobile`: **208 domains walked**, **14 admits**, and **239/4480 (5.3%)** readable pairs; attempted-load coverage is **901/8960 (10.1%)**.
- The report explicitly classifies the result as a pilot, not a Russian-market sample.

## Reconciliation

The pane and report agree on the operative measurements: 14 admitted domains and 239/4480 readable pairs. The earlier receipt's 209-domain prose is stale relative to the regenerated report; the current on-disk canonical presentation is 208. This is presentation/history drift, not a change to the admitted count or readable-pair result.

Integrity hashes after regeneration:

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
```

Prerequisites: `/run/netns/ru-mobile` is absent; `timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true` exited 0. SSH reachability does not establish the missing approved namespace.

No capture, seat onboarding, payment, GAID reset, router write, external contact, or other irreversible action occurred.

Retry edge: when the approved `/run/netns/ru-mobile` exists with verified RU-mobile egress, rerun the canonical paired command with both-arm vantage verification and reconcile the 4,480-pair target.

Delegation decision: no subagent launched. This was one tightly coupled live-pane/read-only-report reconciliation; splitting it would duplicate the same source state. The owner personally inspected the pane, status output, hashes, namespace check, SSH check, and receipt.
