# Step 0d current status reconciliation — 2026-09-18

Status: **pilot remains open; no capture started**.

## Fresh read-only checks

Ran `python3 tools/adint-status --include-private` from `/home/mesh-home/self-adint` at
2026-09-18 14:13 UTC; exit 0. It regenerated `report.html` from the files already on disk.

- canonical `ru-mobile`: 209 domains, 14 admits, 3 duplicate rows collapsed;
- readable paired coverage: **239 / 4480 (5.3%)**;
- attempted-load coverage: 901 / 8960 (10.1%);
- verdict: pilot only, not a Russian-market sample.

The current report lists the same 239 readable pairs and retains the explicit per-site loss and
UNKNOWN/blindness accounting. No capture, seat onboarding, payment, GAID reset, router write, or
external contact was performed.

## Prerequisite evidence

- `/run/netns/ru-mobile`: absent (the existence check printed `absent`);
- `timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true`: exit 0;
- the SSH reachability does not establish the missing approved RU namespace, so the paired-run
  acceptance remains blocked on that external/runtime capability.

## Integrity hashes observed

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
```

Retry edge: when the approved `/run/netns/ru-mobile` exists with verified RU-mobile egress,
rerun the canonical paired command with both-arm vantage verification and reconcile the 4,480-pair
target. This receipt does not claim completion.
