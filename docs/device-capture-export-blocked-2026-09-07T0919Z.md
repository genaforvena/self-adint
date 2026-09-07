# Device capture export — blocked disposition — 2026-09-07T09:19Z

Task: `self-adint-device-capture-export`

Status: `BLOCKED_OPERATOR_EXPORT`

## Current audit

The exact board claim was live at 2026-09-07T09:18Z. The current implementation remains
read-only: `tools/adint-device-capture` only pulls an operator-exported PCAPdroid CSV when
`--pull` is explicitly supplied; it does not control Android UI, start or stop a VPN, or guess a
remote path.

No complete operator-exported PCAPdroid CSV is present in `data/`, and the operator has not
provided its exact reachable Termux SSH path. Existing `.pcap` files and older device-entry
artifacts are not the requested CSV export and cannot be substituted without changing the task.

## Safe action performed

The exact read-only preview was run; no SCP, Android UI, VPN, or measurement-lane write occurred:

```json
{"command": ["scp", "-P", "8022", "<termux-user>@<phone-host>:<operator-export-path-required>", "data/device-observations.jsonl.csv.part"], "out": "data/device-observations.jsonl"}
```

The next action is specific: the operator supplies one complete PCAPdroid CSV export and its
exact remote path; then rerun the documented `--pull` command and verify the JSONL plus SHA-256
manifest. Until then, no device observation or bundle-entry artifact may be claimed.

## Verification

```text
python3 tools/adint-device-capture --test  PASS
python3 tools/adint-study --test            PASS
```
