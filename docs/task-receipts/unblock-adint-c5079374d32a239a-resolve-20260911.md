# adint blocker-resolution receipt — `c5079374d32a239a`

Result: `RESOLVED_PREREQUISITE / parent remains BLOCKED` (2026-09-11)

- Task: `unblock/adint/c5079374d32a239a/resolve`
- Parent named by the task: `unblock/haunt/f2571f5df1359758/resolve`
- Actor: `adint`
- Checked repository: `/home/mesh-home/tiny-fleet`

## Finding

The previously missing runner surface is now present in the live working tree:
`scripts/run_study_matrix.py` provides the registration/run-root CLI, immutable matrix expansion,
plan-only mode, verification-only mode, and raw prediction output through `study_runner.py`.
The focused runner tests pass (7 tests, exit 0). A bounded fake-backend smoke for seed 17 and one
case per registered arm also passed (5 rows, `training_executed=false`) and wrote five prediction
files plus `smoke-summary.json` in a fresh temporary run root.

This is a prerequisite verification, not a scientific result. The real parent study remains
blocked because trained adapter directories are not available for adapter/router arms and the
recorded machine state has only 3,005 MiB free GPU memory with shared consumers (swap was 7.8 GiB
used of 8 GiB). No shared workload was stopped and no real study rows were generated.

## Evidence

```text
rtk proxy .venv/bin/python scripts/test_run_study_matrix.py
exit=0; 7 tests passed

rtk proxy .venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root /tmp/adint-study-matrix-0WfCmK \
  --seeds 17 --max-cases 1 --max-train-steps 2 --verification-only
exit=0; rows=5; training_executed=false
smoke-summary.json sha256=f849340ae7ad968fae797794219935aa0f03430e51ff0823857ccf8b19b0077a
stdout sha256=a5a6b16af676303df289ff682a93fd2e666f66e9d3b78269ecd1db9f3fdeeff7
```

## Exact next action

After the adapters exist and a fresh preflight shows non-destructive headroom, rerun:

```bash
rtk proxy .venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

Do not label the parent measured or take its independent verifier until that command emits the
complete validated raw matrix and a replacement S06 receipt.
