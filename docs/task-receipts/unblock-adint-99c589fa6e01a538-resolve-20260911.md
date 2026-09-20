# adint runner unblock receipt — typed infrastructure block

Result: `BLOCKED/dependency` (2026-09-11T23:38:20Z)

- Task: `unblock/adint/99c589fa6e01a538/resolve`
- Actor: `adint`
- Parent: `unblock/haunt/f2571f5df1359758/resolve`
- Repository checked: `/home/mesh-home/tiny-fleet`

## Evidence

The runner implementation is present and the frozen registration is unchanged:

```text
scripts/run_study_matrix.py sha256=d12292a9787162f50c94dcf86082d49690a2c4564936bf957a2a7360cc1ea6c8
scripts/study_runner.py    sha256=0d67982d5920ef3a19e59f56d6fc4e5dff184a9699b9f410c508fb1da8ca1310
runs/fleet-study-v1/registration.json sha256=2498cc3146b673751f15106d39c31b059188c9af60668dfa6998e00eb5e230ba
```

Required frozen adapter directories are all absent:

```text
adapters/study-pooled MISSING
adapters/study-toy_passage_ppl MISSING
adapters/study-executable_code MISSING
adapters/study-rated_style MISSING
adapters/study-adversarial_safety MISSING
```

Live resource evidence at receipt time: RTX 3060, 12288 MiB total, 9572 MiB used,
2341 MiB free; system swap 8.0 GiB total, 7.9 GiB used, 78 MiB free. No shared
workload was stopped and no training was started.

## Decision

No safe local prerequisite remains in this window. The runner explicitly refuses to
substitute a base model or unrelated adapter when a registered adapter is missing;
fabricating/copying one would invalidate the study. Training the five frozen adapters
also requires shared GPU headroom that is not available in this observation.

## Exact retry

After the five registered adapter artifacts are created by the owning training task,
and after a fresh resource preflight confirms safe shared headroom, run:

```bash
cd /home/mesh-home/tiny-fleet
rtk proxy .venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

This receipt is not a scientific result; the complete 15-row raw matrix remains open.
