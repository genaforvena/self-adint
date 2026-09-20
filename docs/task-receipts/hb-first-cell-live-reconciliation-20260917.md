# HB first-cell live reconciliation — 2026-09-17

Status: **pilot remains open; fresh disk-backed measurement receipt**.

## Action and evidence

Ran `python3 tools/adint-status --include-private` at 2026-09-17 21:47 UTC against the live
files on disk. It exited 0 and regenerated `report.html`.

- canonical `ru-mobile`: 209 domains, 14 admits, 3 duplicate rows collapsed;
- readable paired coverage: **239 / 4480 (5.3%)**;
- attempted-load coverage: 901 / 8960 (10.1%);
- the report explicitly keeps the corpus as a pilot, not a Russian-market sample.

The fresh report lists the canonical admitted set and per-cell/per-site readable counts. No
capture, seat onboarding, payment, GAID reset, router write, or external contact was performed.

## Integrity anchors

Observed after regeneration:

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
```

## Delegation and observability

Delegation decision: no subagent launched. Reconciliation, receipt creation, and settlement are
one tightly coupled single-writer action; splitting them would not produce independently
verifiable work.

The owner mind personally ran and inspected the status output, hashes, and working-tree state.
After settlement, `mesh-dash --once adint` is the required pane:adint observation check.

## Remaining edge

The next safe action is still restoration of the approved `ru-mobile` namespace, followed by the
canonical paired command with both-arm vantage verification. That external/runtime prerequisite
is not performed by this receipt.
