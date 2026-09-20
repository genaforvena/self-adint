# Live study status refresh — 2026-09-20 05:09 UTC

Task: `adint-goal-20260920-live-study-status-refresh-0504/live-study-status-refresh-0504`
Owner: `adint`

## Fresh read-only evidence

- `python3 tools/adint-study --test` — exit 0; `selftest: PASS`.
- `python3 tools/adint-status --include-private` — exit 0; report regenerated successfully.
- `mesh-dash --once adint` — exit 0; `pane:adint` rendered at `2026-09-20T05:09:01Z` with `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480 readable pairs`.
- Current passive corpus: `data/passive-2026-09-20.jsonl`, 2,994 lines / 8,759,749 bytes, mtime `2026-09-20 04:17:34.334476699 +0000`.
- Collection ledger: `data/passive-collect-ledger.jsonl`, 387 lines / 405,689 bytes, mtime `2026-09-20 04:17:34.650476722 +0000`.
- Latest ledger event: `2026-09-20T04:17:34Z`, 4/4 sites, `rc=0`, 183 bid candidates, 1,844,233 bytes added; egress legs agree on `38.49.216.141`.

The status remains a pilot: 239/4480 readable pairs (5.3%), 901/8960 attempted loads (10.1%), 14 admits across 208 domains. The paired RU-mobile gate remains `UNKNOWN / held before capture`; no namespace creation, capture, router write, external probe, exchange contact, publication, onboarding, payment, or GAID reset was performed.

## Artifact hashes

```text
a2cbd6766981cba4378760c450674e62409bb3d4b9e2aaef5b9865ba2ae08fdf  data/passive-2026-09-20.jsonl
491a75a056d8803fee6da66fee038f3d30ce96f27bd93d0ca27d59effdf39836  data/passive-collect-ledger.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
```

## Delegation and inspection

Delegated one non-overlapping read-only corpus/gate audit to worker `adint-corpus-audit`. I personally inspected its temporary report `/tmp/adint-corpus-audit-report.md` (sha256 `5be57ee904f21ce69eb7e2b1cb62cd4be7f34baa2e54623c19d8c021d35c8475`) and independently reran the self-test, status command, corpus/ledger stats, hashes, and `pane:adint` render. The worker report agrees that the fresh corpus justifies this bounded refresh and that the paired gate must remain held; the report was advisory and did not settle this task.

## Next edge

Wait for the externally approved `/run/netns/ru-mobile` namespace and then rerun the local status lane plus `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1` before any paired measurement. No operator approval is requested for this read-only refresh.
