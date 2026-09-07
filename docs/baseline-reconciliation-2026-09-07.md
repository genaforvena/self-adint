# Baseline reconciliation — 2026-09-07

This note records the read-only baseline before the Self-ADINT expansion work.

## Repository state

- Repository: `/home/mesh-home/self-adint`
- Revision: `12ecf70` (`adint step4: wire the passive collection lane, and unblock the guard that refused it`)
- Dirty paths at inspection:
  - modified: `docs/INDEX.md`
  - modified: `docs/step4-passive-lane-2026-08-30.md`
  - modified: `tools/adint-collect-passive`
  - untracked: `docs/device-capture-runbook-2026-09-07.md`
  - untracked: `docs/full-study-review-2026-09-07.md`
  - untracked: `docs/superpowers/`
  - untracked: `tools/adint-device-capture`
  - untracked: `tools/adint-study`

The 2026-09-07 changes are a coherent draft: redundant passive IP-echo attempts, the explicit
read-only device handoff, and artifact-derived study status. They are retained. No unrelated path
was changed or reverted. The existing receiver corpus and prior study artifacts remain untouched.

## Offline checks

The targeted expansion checks passed:

```text
python3 tools/adint-collect-passive --test                 PASS
python3 tools/adint-device-capture --test                 PASS
python3 tools/adint-study --test                          PASS
python3 tools/adint-import-pcapdroid --test               0 checks failed
python3 tools/adint-aggregate --test                      0 checks failed
python3 tools/adint-status --test                         selftest: ok
(cd receiver && go test ./...)                            PASS
python3 -m py_compile [changed Python tools]              PASS
```

The documented repository-wide command was also attempted. All tools before
`tools/adint-conn-report` passed; that tool has no `--test` mode and instead interpreted the
flag as a CSV path, exiting 1 with `FileNotFoundError`. This is a test-harness compatibility
issue, not a measurement result, and is left unchanged because it is outside this expansion's
immediate action.

## Status interpretation

`data/study-status.json` currently reports:

- device entry-point: `MEASURED` from existing bundle-entry artifacts;
- local receiver and passive web lane: `MEASURED`;
- device capture lane: `BLOCKED_OPERATOR_EXPORT`;
- active oracle: `NOT_ELIGIBLE`;
- GAID reset: `NO-BASIS`;
- seat: `GATED_OPERATOR_GO`.

The device status is correctly blocked even though older Step 0 entry-point artifacts exist: the
required current PCAPdroid export has not been imported. No absent export is treated as no demand.

## Reconciliation result and next action

Retain the dirty 2026-09-07 draft and proceed to the plan's immediate device-lane action. The
remaining blocker is one complete PCAPdroid CSV export from the operator's device plus a reachable
host and explicit device label. Until that is supplied, only `--dry-run`, fixture tests, and
read-only documentation/status work are permitted.

## Receiver expansion reconciliation

The local receiver was unit-green; the missing piece was a fresh process-level artifact.
`tools/adint-receiver-run` supplies that run in an explicit dated directory, with a SHA-256
manifest and the red-before-green proof in `docs/receiver-red-before-green-2026-09-07.md`.

The requested device-observations JSONL remains `BLOCKED_OPERATOR_EXPORT`. No complete PCAPdroid
CSV is present, so no empty or synthetic device artifact was written. Provide one complete CSV
export and its device label, then run `tools/adint-device-capture --pull` as documented.
