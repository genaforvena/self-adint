# Step 0d post-fix paired-run preflight — 2026-09-16

Status: **blocked before live capture; no new corpus written**.

## Personally verified

- `python3 tools/adint-status --test`: exit 0 (`selftest: ok`).
- `python3 tools/adint-frame-stageb --test`: exit 0 (`selftest: ok`).
- `tools/adint-paired-run --test`: exit 0 (`selftest: ok`).
- `python3 tools/adint-status --include-private`: exit 0; the canonical `ru-mobile` study remains
  at **239 / 4480 readable pairs (5.3%)**, so the cell is still a pilot and not complete.
- `timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true`: exit 0.
- `/run/netns/ru-mobile` is absent (read-only existence check exit 1).

The three local tests cover the runner's privacy, frame, and paired-vantage invariants, but they do
not create measurement data. No capture was started because the required RU namespace is not
present on this node; this avoids writing a misleading one-arm or unverified-vantage corpus.

## Acceptance and exact next edge

The planned acceptance remains unmet: both arms must verify their vantages and the paired corpus
must reach 4,480 readable pairs. The next safe action is to provision/restore the approved RU-arm
namespace outside this repository, then rerun the canonical paired command with both-arm output
and record its hashes and reconciliation here. Provisioning is not performed by this window.

Task ownership was attempted for `adint-goal-20260916-step0d-postfix-paired-run/step0d-postfix-paired-run`,
but `mesh-task check/take` did not return within 20 seconds during concurrent ledger activity; the
task-create journal entry exists. Retry `MESH_TASK_ACTOR=adint mesh-task take ...` after ledger
contention clears before live execution.

Delegation: `adint-live-audit` performed a read-only independent audit. I personally inspected this
receipt's cited state and reran all commands above; the worker report was not treated as evidence.
