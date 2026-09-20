# adint resolver receipt — study adapter producer is still missing

Result: `BLOCKED/dependency` (2026-09-12T03:26Z)

- Resolver: `unblock/bash/f1909875a2fb4a54/resolve` (owner `adint`)
- Parent: `unblock/adint/bc46beffbe51292e/resolve` (still blocked)
- S06 owner step: `tinyfleet-publication-science-20260908/run-paired-replications` (owner `haunt`)
- Study repository: `/home/mesh-home/tiny-fleet`

## Current evidence

The blocker is live and accurately names missing trained study adapters. The frozen registration
`runs/fleet-study-v1/registration.json` has five arms and seeds 17/29/43 (15 matrix rows), with
SHA-256 `2498cc3146b673751f15106d39c31b059188c9af60668dfa6998e00eb5e230ba`. The current runner
`scripts/run_study_matrix.py` has SHA-256
`d12292a9787162f50c94dcf86082d49690a2c4564936bf957a2a7360cc1ea6c8`.

All required directories are absent:

```text
adapters/study-pooled
adapters/study-toy_passage_ppl
adapters/study-executable_code
adapters/study-rated_style
adapters/study-adversarial_safety
```

`require_adapters()` in the current runner refuses missing pooled/specialist paths. `run_registered_arm()`
loads existing adapters with `TransformersBackend`; it does not train them. The CLI accepts
`--max-train-steps` only for compatibility and documents that no training is performed. The current
result field `training_executed: not args.verification_only` would label ordinary inference as
training, so that field cannot be treated as proof that training ran. The only existing
`scripts/train_eval.py` trainer hard-codes the legacy `guitar` and `sourdough` corpus domains; it is
not a producer for these five frozen study artifacts. The S06 plan assigns the complete matrix to
the still-blocked `run-paired-replications` step, and its current implementation files are untracked
or modified in the shared worktree. I left that owner's files untouched.

Fresh resource observation: GPU 12,288 MiB total / 7,688 MiB free / 0% utilization; RAM 24 GiB
available; swap 1.8 GiB free. This is only a preflight observation, not evidence that five training
runs fit the registered cap or a substitute for implementing the study's missing training path.

## Verification

```text
rtk proxy .venv/bin/python scripts/test_run_study_matrix.py
exit=0; 10 tests passed, including refusal of a missing adapter without fallback

rtk proxy .venv/bin/python scripts/run_study_matrix.py --registration runs/fleet-study-v1/registration.json --run-root runs/fleet-study-v1 --plan-only
exit=0; 15 registered arm/seed rows; plan-only, no execution or writes
```

No full-matrix prediction JSONL or smoke summary was found under `runs/fleet-study-v1` at audit
time. I did not invoke the real runner: its non-verification path writes a resource receipt before
refusing the absent adapters, and it cannot produce the missing trained artifacts.

## Exact next action

The S06 owner must implement and execute the preregistered training for the pooled union and four
specialist domains from the frozen `corpus/study-v1` inputs, recording per-run configuration,
seed, input hashes, optimizer steps, resource/time data, and adapter hashes under the exact five
paths above. The S06 runner must report training truthfully before the full matrix is run. After
those artifacts exist, perform a fresh non-destructive resource preflight, then run:

```bash
cd /home/mesh-home/tiny-fleet
rtk proxy .venv/bin/python scripts/run_study_matrix.py \
  --registration runs/fleet-study-v1/registration.json \
  --run-root runs/fleet-study-v1
```

No adapter was fabricated or copied, no training or matrix run was started, and the parent remains
blocked until the S06-owned artifacts and full raw matrix exist.
