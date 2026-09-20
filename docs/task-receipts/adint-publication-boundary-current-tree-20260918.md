# Current publication-boundary audit — 2026-09-18

Status: **BLOCKED — private-looking identifiers remain in the tracked public tree**.

## Scope

- Task: `adint-publication-boundary-current-tree-20260918/current-worktree-publication-boundary-audit`
- Delegation: no subagent; this was a tightly coupled single-snapshot filesystem audit. Splitting the status, privacy scan, and tests would risk mixing different worktree states.
- No device capture, exchange contact, router write, VM change, or publication was performed.

## Evidence personally inspected

- `git status --short`: 9 modified tracked files and untracked publication/task artifacts, including the current Stage B RU frame, a lock file, and the offline test.
- `git diff --check`: exit 0.
- `git grep` over tracked non-`data/` files: matched browser/shell egress IP fields in `public-data/frame-stageb-nl-direct-2026-08-19-schema3.jsonl`, the RU egress summary in `README.md`, and additional egress-IP rows in `docs/step0c-hb-ru-market-study-2026-08-18.md`, `docs/step0d-hb-first-cell-2026-08-19.md`, and tracked Stage B data. The scan also matched `UserID` fields in tracked public capture data. Values are intentionally not reproduced here.
- `mesh-dash --once adint`: rendered the active row and current `step0d-hb-first-cell` signal on `pane:adint`.

## Verification

- `python3 tools/test-step1-egress-offline.py`: exit 0, 2 tests passed.
- `(cd receiver && go test ./...)`: exit 0, `selfadint/receiver` passed.
- `python3 tools/adint-status --test`: exit 0, `selftest: ok`.
- An initial combined probe had a shell quoting error before tests; it was corrected and rerun. No test result is claimed from the failed probe.

## Disposition

The executable gates pass, but the publication boundary is not clear: tracked public artifacts contain egress-IP and `UserID` fields. The exact next action is a separate owner decision to redact/remove those fields from the public corpus and regenerate dependent reports, or explicitly establish that the fields are approved public method data. This audit does not make that irreversible publication decision.

