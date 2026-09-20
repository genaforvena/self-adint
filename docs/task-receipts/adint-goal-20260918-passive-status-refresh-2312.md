# Passive study status refresh — 2026-09-18T23:15Z

Task: `adint-goal-20260918-passive-status-refresh-2312/refresh-safe-passive-status`

## Commands and evidence

- `python3 tools/adint-study --test` → `selftest: PASS`.
- `python3 tools/adint-study` → fresh measured status below.
- Fresh passive artifact: `data/passive-2026-09-18.jsonl`, mtime `2026-09-18T23:15:22.985021468Z`, size `33,417,344` bytes.
- Current passive ledger: `data/passive-collect-ledger.jsonl`, mtime `2026-09-18T22:17:34.345774083Z`, size `370,021` bytes.
- No GAID reset, seat onboarding, payment, router change, or external routing-shadow action was performed.

## Measured result

| lane | status | next action |
|---|---|---|
| step0-device-entry-points | MEASURED | rederive from append-only records |
| step1-local-receiver | MEASURED | audit win and spend markers |
| step2-questions | DRAFT_ONLY | operator sends questions after choosing a candidate |
| step3-seat | GATED_OPERATOR_GO | operator explicitly authorises onboarding, agreement, and payment |
| step4-device | BLOCKED_OPERATOR_EXPORT | import an operator-exported PCAPdroid CSV |
| step4-web | MEASURED | inspect passive ledger and corpus |
| step5-oracle | NOT_ELIGIBLE | wait for measured payload-empty/hashed evidence |
| step6-gaid-reset | NO-BASIS | build passive baseline, then request explicit reset go |

Exact next safe action: inspect the passive ledger/corpus for the measured `step4-web` lane. The GAID reset remains closed because the status is `NO-BASIS`.
