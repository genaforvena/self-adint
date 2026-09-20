# Live study status reconciliation — 2026-09-18

Task: `adint-goal-20260918-live-study-status-reconciliation/live-study-status-reconciliation`
Owner: `adint`
Observed: `2026-09-18T22:43:38Z`

## Checks

From `/home/mesh-home/self-adint`:

```text
python3 tools/adint-study --test                         exit 0; selftest: PASS
python3 tools/adint-study                              exit 0
runbook status coverage check                            exit 0; 8 statuses, missing=[]
git diff --check                                         exit 0
```

The live coordinator emitted these statuses and next signals:

| phase | status | next signal |
|---|---|---|
| step0-device-entry-points | `MEASURED` | rederive from the append-only records |
| step1-local-receiver | `MEASURED` | audit win and spend markers |
| step2-questions | `DRAFT_ONLY` | operator sends questions after choosing a candidate |
| step3-seat | `GATED_OPERATOR_GO` | operator explicitly authorises onboarding, agreement, and payment |
| step4-device | `BLOCKED_OPERATOR_EXPORT` | import an operator-exported PCAPdroid CSV |
| step4-web | `MEASURED` | inspect passive ledger and corpus |
| step5-oracle | `NOT_ELIGIBLE` | wait for measured payload-empty/hashed evidence |
| step6-gaid-reset | `NO-BASIS` | build passive baseline, then request explicit reset go |

The runbook covers all eight status literals emitted by `tools/adint-study`, including
`UNKNOWN`, `READY_OPERATOR_GO`, and the live statuses not emitted in this read. The named
local command and input-path inventory was checked: `tools/adint-study`, `tools/adint-status`,
`tools/adint-step1-run`, the Step 1 summary, `data/bundle-exchange.tsv`, and the passive ledger
exist; `data/device-observations.jsonl`, `data/step4-payload-verdict.json`, and
`data/step4-baseline.json` are absent. Absence is preserved as a gate/unknown condition, not a
negative measurement.

## Verdict

`PASS` for status-literal coverage and read-only status rendering. The live result preserves the
operator gates: no device export is fabricated, no active targeting is enabled, and no GAID reset,
seat onboarding, payment, external contact, capture, publication, or router write was performed.

The next safe edge is to inspect the phase-specific artifact for a `MEASURED` row; for
`BLOCKED_OPERATOR_EXPORT`, retry after an operator-owned PCAPdroid export with its time window and
source application; for `NO-BASIS`, build the passive baseline before asking for reset authorization.

## Reproduction hashes

```text
c9053c43371d8a6b9db33d3ca3f2e3f0168839bc62e6223683bd6e132a876ce0  tools/adint-study
268c7d35f400c6b8907a4f32d65f432d408d8f72f73d0fb5872bc34ffc8fe2f8  docs/status-next-action-runbook-20260918.md
8aeb5cc5f4be65d6bec48566e1a41c9fd432636edd158154bc3aa3d48739f24c  docs/task-plans/adint-goal-20260918-live-study-status-reconciliation.tsv
```

`pane:adint` observed the task as `HOLD adint-goal-20260918-live-study-status-reconciliation`
and the current Step 0d summary (`14 admits/208 domains; 239/4480 readable`) at
`2026-09-18T22:43:40Z`.
