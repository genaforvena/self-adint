# Device capture disposition — 2026-09-07T08:58:43Z

Task: `owner-adint-self-adint-device-lane-conti`

Status: `BLOCKED_OPERATOR_EXPORT`

The board request asked for exactly one complete PCAPdroid CSV from the operator's Redmi,
using the explicit label `redmi10-daily`. The operator has not supplied the exported CSV's
exact reachable Termux SSH endpoint or remote path. No CSV is present in `data/`, so no read-only import, manifest, observations,
or bundle-entry artifact can honestly be produced.

Recorded board request: 2026-09-07T08:58:43Z, from `health@mesh-home`.

Dry-run result:

```json
{"command":["scp","-P","8022","<termux-user>@<phone-host>:<operator-export-path-required>","data/device-observations.jsonl.csv.part"],"out":"data/device-observations.jsonl"}
```

This was a preview only. No Android UI, VPN capture, network probe, SCP, or write to the
measurement lane was performed. Next action: operator supplies one complete export and its
exact remote path; then rerun this command with `--pull` and verify the manifest/import.
