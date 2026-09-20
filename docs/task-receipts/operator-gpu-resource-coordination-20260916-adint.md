# Adint coordinated GPU recovery — 2026-09-16

Task: `operator-gpu-resource-coordination-20260916/verify-adint-recovery`
Owner: `adint`
Checked: `2026-09-16T19:54:52Z`–`2026-09-16T19:55:10Z` UTC

## Eligibility and prerequisite evidence

- Canonical chain status before execution: parent active, this exact row active,
  owner `adint`; `mesh-task take operator-gpu-resource-coordination-20260916
  verify-adint-recovery` returned `already active`.
- `mesh-task preflight ... verify-adint-recovery` returned `n/a (not a blocked
  resource claim)`; the row was nevertheless valid and already claimed.
- Parent shared-GPU receipt: `/home/mesh-home/.mesh/evidence/receipts/operator-gpu-resource-coordination-20260916.md`.
- Measured minimum-free threshold inherited from the completed coordinated recovery:
  `5600 MiB`.

## Live resource decision

Before the run:

```text
ollama ps: header only; no resident models
nvidia-smi: RTX 3060, total=12288 MiB, used=1 MiB, free=11911 MiB, util=0%
mesh-gpu-lease --status: GPU_LEASE=none
```

The predicate passed without preemption. The consumer was still invoked through the
prescribed bounded wrapper:

```bash
MESH_HEAVY_GPU_PREEMPT=1 mesh-heavy-run 5600 -- \
  bash -c 'python3 tools/adint-study --test && \
           python3 tools/adint-status --test && \
           python3 tools/adint-study'
```

## Real consumer result

```text
python3 tools/adint-study --test: exit 0, selftest: PASS
python3 tools/adint-status --test: exit 0, selftest: ok
python3 tools/adint-study: exit 0, live status rendered
```

The live status was:

```text
step0-device-entry-points: MEASURED
step1-local-receiver: MEASURED
step2-questions: DRAFT_ONLY
step3-seat: GATED_OPERATOR_GO
step4-device: BLOCKED_OPERATOR_EXPORT
step4-web: MEASURED
step5-oracle: NOT_ELIGIBLE
step6-gaid-reset: NO-BASIS
```

After the run:

```text
ollama ps: header only; no resident models
nvidia-smi: RTX 3060, total=12288 MiB, used=113 MiB, free=11799 MiB, util=1%
mesh-gpu-lease --status: GPU_LEASE=none
return codes: before ollama=0, before gpu=0, before lease=0,
               consumer=0, after ollama=0, after gpu=0, after lease=0
```

No mesh-managed service required restoration because the lease was never acquired;
no unrelated GPU process or model was stopped. The exact task is complete and the
parent may proceed to its cross-owner gate. The top witness pane `mesh-home:13.0`
rendered the adint goal and this row as open/undischarged during the pre-settlement
observation; settlement follows this artifact and is the next ledger state change.
