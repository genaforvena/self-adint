# Resolver receipt: `unblock/adint/3d81ef8eeb289db5/resolve`

- Checked: `2026-09-12T02:33Z` UTC on `mesh-home`
- Parent: `unblock/haunt/62c0662129fa8ee9/resolve` (owner `haunt`)
- Result: **operator-input finding corrected; evidence reconciled; parent remains blocked pending its owner's decision on the experiment mismatch**

## Evidence and authorization

The owner-scoped dispatch row passed `mesh-task check dispatch` (exit `0`) and was claimed as
`adint`. The parent remains `blocked` with blocker `operator-input`. Its owner-authored resume gate,
`mesh-task check resume unblock/haunt/62c0662129fa8ee9/resolve haunt`, still exits `2`.

The prior claim that the five dictionary inputs were absent is superseded by the operator's Telegram
authorization and the correction recorded in
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`.
The operator directed us to choose from his repositories (`00:42` and `00:47` in
`/home/mesh-home/.mesh/textin.log`). That receipt freezes an existing CC0, synthetic-only A08
fixture, language set `en`/`ru`, exact replacement normalization, immutable source commit, and
corpus bytes. No additional operator question is needed.

I independently recomputed the referenced on-disk SHA-256 values; each matches the receipt:

| Artifact | SHA-256 |
|---|---|
| `runs/operator-selected-dictionary-v1/input-manifest.json` | `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08` |
| `runs/operator-selected-dictionary-v1/corpus.jsonl` | `81a41340c0c01555a536e9aa30b3b2b72b1dd9ae48908014fd0f58192515dcd3` |
| `corpus/applications/noisy-text-correction/manifest.json` | `17879f42f885b18ed5e0638a6620ff745200ddb79d9a3004a20fa22933ea0997` |
| `scripts/applications/noisy_text_correction.py` | `69330b7803f0844198c84f07d54cc89e822d4df9abc63ddfb58b6c3c921c60db` |
| `runs/operator-selected-dictionary-v1/result/raw.jsonl` | `555f1ab7808d5d01324c770e58fc594b48ea17ff55364d6a4214e4aed8b988fd` |
| `runs/operator-selected-dictionary-v1/bbywvy-smoke.txt` | `f64847a15bc68114b0d369200a0173d3f37e8a76c13fe499d93f9854eb4cffc6` |

The source receipt records five passing A08 tests and a fresh dictionary baseline of 120/120 exact,
but explicitly labels the verdict `INCONCLUSIVE`: it is a five-substitution synthetic fixture, not a
general-purpose language lexicon, and the ByT5/LoRA arms are unavailable. Its separate offline
BbyWVY smoke exited `0` and produced six outputs, but the math probe answered `9000 km` for
60 km/h over 2.5 hours. The smoke establishes runtime only; it has no dictionary arm and does not
validate dictionary behavior.

## Reconciliation and next action

The missing-operator-input diagnosis is no longer accurate: operator authorization exists and the
five selected fields plus corpus are frozen and measured. However, the parent retry names
`scripts/bbywvy_test.py`, which has no dictionary arm. The measured A08 result cannot be represented
as a BbyWVY dictionary result or as a general-language result. This is an experiment-contract gap,
not a missing operator answer.

The `haunt` owner should reconcile the frozen A08 evidence with the parent task and choose one
artifact-backed path: accept the narrowly scoped synthetic A08 result for the dictionary
prerequisite and update the parent receipt/retry accordingly, or add/identify the intended
BbyWVY-compatible dictionary arm and measure it. Only `haunt` may update or resume its row. I did
not alter the parent ledger, change source/runtime/model files, or run the parent retry; the owner
gate refused resume with exit `2`.
