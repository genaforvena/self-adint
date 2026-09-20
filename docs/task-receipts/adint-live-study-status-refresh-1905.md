# Live study status refresh — 2026-09-19 19:05 UTC

Task: `adint-goal-20260919-live-study-status-refresh-1905/live-study-status-refresh-1905`

This was a read-only local reconciliation. No device capture, external request, publication,
router write, or irreversible action was performed.

## Commands and results

```text
cd /home/mesh-home/self-adint
python3 tools/adint-study --test    # exit 0: selftest: PASS
python3 tools/adint-study --test    # live status command below was preceded by this check
python3 tools/adint-study            # exit 0
```

The live coordinator emitted:

| phase | status | next action |
|---|---|---|
| step0-device-entry-points | MEASURED | rederive from append-only records |
| step1-local-receiver | MEASURED | audit win and spend markers |
| step2-questions | DRAFT_ONLY | operator sends questions after choosing a candidate |
| step3-seat | GATED_OPERATOR_GO | operator explicitly authorises onboarding, agreement, and payment |
| step4-device | BLOCKED_OPERATOR_EXPORT | import an operator-exported PCAPdroid CSV |
| step4-web | MEASURED | inspect passive ledger and corpus |
| step5-oracle | NOT_ELIGIBLE | wait for measured payload-empty/hashed evidence |
| step6-gaid-reset | NO-BASIS | build passive baseline, then request explicit reset go |

## Input artifact evidence

```text
data/bundle-exchange.tsv | size=7195 | sha256=5272ec47853f9f267f7cd69cdb1a5c33692a950138fb0f54758a6249a1a7ddf1
data/device-observations.jsonl | ABSENT
data/passive-collect-ledger.jsonl | size=393811 | sha256=a33fe644aadd14d6e0bce1def79c6872e0bb1c08dd4d20a0b9bdf45ed2e125c3
data/step4-payload-verdict.json | ABSENT
data/step4-baseline.json | ABSENT
receiver/testdata/step1-2026-09-07-final2/summary.json | size=373 | sha256=12685729b78af0c486c0110af0bcc9df60889740b1d25e007924ffd583c08b63
```

Exact command output is retained for this turn at `/tmp/adint-study-test-1905.txt` and
`/tmp/adint-study-live-1905.txt`. The safe next action is to preserve the device row as
`BLOCKED_OPERATOR_EXPORT`, continue the passive lane, and leave seat onboarding and GAID reset
behind their explicit operator gates.

Observed on `pane:adint` after settlement.
