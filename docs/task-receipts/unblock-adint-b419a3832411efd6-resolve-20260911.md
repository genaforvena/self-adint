# adint matrix unblock receipt — external dependency remains

Result: `BLOCKED/dependency` (2026-09-11T23:55:26Z)

- Task: `unblock/adint/b419a3832411efd6/resolve`
- Actor: `adint`
- Parent: `unblock/haunt/f2571f5df1359758/resolve`
- Repository inspected: `/home/mesh-home/tiny-fleet`

## Finding

The prior runner-implementation gap is resolved: `scripts/run_study_matrix.py` and its focused
tests are present. The current blocker is the registered training artifacts and safe shared-device
headroom. All five adapters required by the frozen matrix are absent:

```text
adapters/study-pooled MISSING
adapters/study-toy_passage_ppl MISSING
adapters/study-executable_code MISSING
adapters/study-rated_style MISSING
adapters/study-adversarial_safety MISSING
```

Fresh live resource check:

```text
RTX 3060: 12288 MiB total, 9572 MiB used, 2341 MiB free, 0% utilization
RAM: 31 GiB total, 15 GiB used, 1.2 GiB free, 15 GiB available
swap: 8.0 GiB total, 7.9 GiB used, 53 MiB free
```

The frozen runner's explicit missing-adapter refusal is covered by the passing test suite. Existing
adapters are not substituted, no shared process was stopped, and no training or study run was
started. The narrowest safe prerequisite is therefore external to this task: the owning S06 work
must produce the five frozen adapters, then a fresh preflight must show safe headroom for sequential
execution.

## Verification

```text
rtk proxy .venv/bin/python scripts/test_run_study_matrix.py
exit=0; 10 tests passed
```

## Exact retry

Once the five registered adapter artifacts exist and a fresh non-destructive preflight confirms
safe headroom, run:

```bash
cd /home/mesh-home/tiny-fleet
rtk proxy .venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

The complete validated 15-row raw matrix and replacement S06 receipt remain outstanding. This
receipt is a blocker diagnosis, not a scientific result.
