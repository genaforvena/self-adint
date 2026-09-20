# adint blocker-resolution receipt — `db71aabbbe5abbb4`

Result: `BLOCKED/dependency` (2026-09-11T22:36:49Z)

- Task: `unblock/adint/db71aabbbe5abbb4/resolve`
- Parent named by the task: `unblock/haunt/f2571f5df1359758/resolve`
- Actor: `adint`
- Repository checked: `/home/mesh-home/tiny-fleet`

## Finding

The matrix runner prerequisite is present and can validate the frozen 15-row plan, but the real
study cannot start safely. All five trained adapter locations required by the runner are absent:

```text
adapters/study-pooled
adapters/study-toy_passage_ppl
adapters/study-executable_code
adapters/study-rated_style
adapters/study-adversarial_safety
```

The current GPU preflight is an RTX 3060 with 12,288 MiB total and 2,842 MiB free. Shared mesh
consumers account for 2,738 MiB (voice), 800 MiB (GigaAM), 5,350 MiB and 158 MiB (Ollama). No
shared process was stopped. The available state is therefore insufficient evidence for a
non-destructive training run, and no real study rows were generated.

## Evidence

```text
.venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root /tmp/adint-db71-plan --plan-only
exit=0; stdout_bytes=596; stdout_sha256=a7cd86ab04675c127743b26ff0416c3cf56154320b254e5821226619eb3ed14c

.venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root /tmp/adint-db71-smoke --seeds 17 --max-cases 1 \
  --max-train-steps 2 --verification-only
exit=0; fake smoke only; training_executed=false
stdout_sha256=28948876df1fd56c1ad3a3be86484f297845218c1a4338353bfe27d080b3cf5c
smoke-summary_sha256=db86fe28b0a848dcebc6c42731325a077a4aef9ff6818030c93b7e302c7d2577

nvidia-smi
NVIDIA GeForce RTX 3060, 12288, 9071, 2842, 0
```

## Exact retry

After the five frozen adapter artifacts exist and a fresh preflight confirms sufficient
non-destructive headroom without stopping shared workloads, run:

```bash
cd /home/mesh-home/tiny-fleet
.venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

Until that command emits the complete validated raw matrix, the parent remains blocked; this
receipt is not a scientific result and does not authorize taking its independent verifier.
