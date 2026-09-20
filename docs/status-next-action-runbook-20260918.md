# Study status → next action

This runbook is the operator-facing continuation of `tools/adint-study`. It is deliberately
read-only: status rendering and local checks do not capture a device, contact a seat, write the
router, publish data, spend money, or reset an identifier.

## Before interpreting a row

Run the deterministic local checks from the repository root:

```sh
python3 tools/adint-study --test
python3 tools/adint-study
```

The second command reads only artifacts already on disk. A missing artifact is not a negative
result; follow the row below and preserve `UNKNOWN`, `NO-BASIS`, and `NOT_ELIGIBLE` as written.

## Status dictionary

| emitted status | exact next local action | operator question, when required |
| --- | --- | --- |
| `MEASURED` | Re-run the phase-specific read-only report named in the phase table, then inspect its artifact and hashes. Do not infer a stronger claim than that artifact records. | None unless the phase table says the measured artifact exposes a decision. |
| `UNKNOWN` | Run `python3 tools/adint-study` and inspect the phase's input paths; for Step 0 use `python3 tools/adint-status --include-private`, and for Step 1 use `go test ./receiver` plus `tools/adint-step1-run --test`. | If the required local artifact is absent after those checks: what local export or capture should be supplied, and when? |
| `DRAFT_ONLY` | Read `PLAN.md` Step 2 and prepare the questions; do not send them. | Which candidate exchange may receive the drafted questions? |
| `GATED_OPERATOR_GO` | No local command advances this state. Verify the relevant evidence with `python3 tools/adint-study`, then stop at the gate. | Does the operator explicitly authorise this irreversible step—seat onboarding, agreement, payment, or the GAID reset? |
| `BLOCKED_OPERATOR_EXPORT` | Run `python3 tools/adint-study` to confirm the block and inspect the local `data/` boundary. Do not fabricate device observations. | Can the operator provide an export from his own device, with its time window and source application identified? |
| `READY_OPERATOR_GO` | Run `python3 tools/adint-study` and inspect the payload-empty/hashed evidence. Do not perform active targeting. | Does the operator explicitly authorise the active targeting step now? |
| `NOT_ELIGIBLE` | Run `python3 tools/adint-study` and inspect the passive payload verdict. Keep the active branch closed until the measured payload-empty/hashed evidence exists. | If the evidence is still absent, no question is needed; the next local action is to continue the safe passive lane. |
| `NO-BASIS` | Run `python3 tools/adint-study` and build the passive baseline using only the safe, already-authorised lane. Do not reset the GAID. | After a baseline artifact exists, does the operator explicitly authorise the destructive GAID reset? |

## Phase rows

| phase row | current status is derived from | exact next local command or question |
| --- | --- | --- |
| `step0-device-entry-points` | `data/bundle-exchange.tsv` | Run `python3 tools/adint-status --include-private` to inspect the device-side entry-point report. If absent, ask for the operator's own device export; an unseen exchange remains `UNKNOWN`. |
| `step1-local-receiver` | `receiver/testdata/step1-2026-09-07-final2/summary.json` | Run `go test ./receiver` and `tools/adint-step1-run --test`. If the fixture is valid, inspect win and spend markers; otherwise preserve `UNKNOWN`. |
| `step2-questions` | fixed `DRAFT_ONLY` state | Read `PLAN.md` Step 2 and ask which candidate exchange the operator wants to address. Sending remains outside this read-only runbook. |
| `step3-seat` | fixed `GATED_OPERATOR_GO` state | Verify with `python3 tools/adint-study`; ask for explicit authorisation before onboarding, signing, or payment. |
| `step4-device` | `data/device-observations.jsonl` | Run `python3 tools/adint-study`. If absent, ask for an operator-exported PCAPdroid CSV; do not create a substitute measurement. |
| `step4-web` | `data/passive-collect-ledger.jsonl` | Run `python3 tools/adint-study`. If measured, inspect the passive ledger; if unknown, start only the safe passive lane described by the existing local tooling. |
| `step5-oracle` | `data/step4-payload-verdict.json` with `verdict=empty_or_hashed` | Run `python3 tools/adint-study`. `NOT_ELIGIBLE` stays closed; `READY_OPERATOR_GO` requires the explicit active-targeting question above. |
| `step6-gaid-reset` | `data/step4-baseline.json` | Run `python3 tools/adint-study`. `NO-BASIS` means build the passive baseline first; never reset without the explicit operator go. |

## Coverage check

The following repository-local check extracts every status literal used by the coordinator and
asserts that this runbook names it. It must exit 0 before this task is closed:

```sh
python3 - <<'PY'
import re
from pathlib import Path

source = Path("tools/adint-study").read_text(encoding="utf-8")
runbook = Path("docs/status-next-action-runbook-20260918.md").read_text(encoding="utf-8")
statuses = set(re.findall(r'"([A-Z][A-Z_]+(?:-[A-Z_]+)*)"', source))
missing = sorted(status for status in statuses if f"`{status}`" not in runbook)
if missing:
    raise SystemExit(f"missing statuses: {', '.join(missing)}")
print(f"status coverage: PASS ({len(statuses)} statuses)")
PY
```

