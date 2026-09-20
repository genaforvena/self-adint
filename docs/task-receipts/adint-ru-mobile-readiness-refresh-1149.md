# RU-mobile readiness refresh — 2026-09-19 11:52 UTC

Task: `adint-goal-20260919-ru-mobile-readiness-refresh-1149/ru-mobile-readiness-refresh-1149`

This was a read-only reconciliation from live disk. No traffic capture, namespace provisioning,
routing change, router write, external contact, seat onboarding, payment, or GAID reset occurred.

## Live observations

- `python3 tools/adint-status --include-private`: exit 0.
- Canonical `ru-mobile`: 209 domains touched, 14 admits, 3 duplicate rows collapsed.
- Coverage: 239 / 4480 readable pairs (5.3%); attempted-load coverage 901 / 8960 (10.1%).
- `mesh-dash --once adint`: exit 0; `pane:adint` rendered `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480` and `-- pane live 2026-09-19T11:52:49Z`.
- `test -e /run/netns/ru-mobile`: exit 1; namespace absent.
- `nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`: exit 1; `cannot open /run/netns/ru-mobile: No such file or directory`.
- `timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true`: exit 255; `No route to host`.

## Canonical input hashes

```text
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
5f6d82344eaee95654e76987340ef4a6f8a4384ce93984195ac45a092d93fd72  data/frame-stageb-ru-mobile-2026-08-19.jsonl
```

## Verdict and retry edge

Readiness is **UNKNOWN / not ready**. The exact retry edge is: after the approved
`/run/netns/ru-mobile` exists, verify its RU-mobile egress with
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`, then run the canonical paired
measurement with both-arm vantage verification and reconcile the 4,480-pair target.

Artifact written by the `adint` mind and observed on `pane:adint`.
