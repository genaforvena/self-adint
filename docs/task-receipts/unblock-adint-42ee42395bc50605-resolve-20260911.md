# adint blocker-resolution receipt — `42ee42395bc50605`

Result: `BLOCKED/dependency` (2026-09-11T22:40Z)

- Task: `unblock/adint/42ee42395bc50605/resolve`
- Parent named by the task: `unblock/haunt/f2571f5df1359758/resolve`
- Actor: `adint`
- Repository checked: `/home/mesh-home/tiny-fleet`

## Finding

The narrow safe prerequisite is not available locally. All five frozen specialist adapter
directories required by the real matrix are absent:

```text
adapters/study-pooled
adapters/study-toy_passage_ppl
adapters/study-executable_code
adapters/study-rated_style
adapters/study-adversarial_safety
```

The current GPU preflight is an NVIDIA GeForce RTX 3060 with 12,288 MiB total, 8,908 MiB used,
and 3,005 MiB free. Shared consumers remain active. No workload was stopped and no real study was
started. The real runner cannot silently substitute another arm: its adapter-refusal test passes.

## Verification artifacts

```text
.venv/bin/python scripts/test_run_study_matrix.py
exit=0; 7 tests passed

.venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root /tmp/adint-resolve-yYBPvA/plan --plan-only
exit=0; plan_sha256=a7cd86ab04675c127743b26ff0416c3cf56154320b254e5821226619eb3ed14c

nvidia-smi --query-gpu=name,memory.total,memory.used,memory.free --format=csv,noheader
NVIDIA GeForce RTX 3060, 12288 MiB, 8908 MiB, 3005 MiB
```

## Exact retry

After all five frozen adapter artifacts exist and a fresh preflight confirms sufficient
non-destructive headroom without stopping shared workloads, run:

```bash
cd /home/mesh-home/tiny-fleet
.venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

Until that command emits the complete validated raw matrix, the parent remains blocked. This
receipt is not a scientific result and does not authorize taking its independent verifier.
