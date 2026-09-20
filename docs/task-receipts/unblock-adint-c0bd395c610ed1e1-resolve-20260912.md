# Resolver receipt: `unblock/adint/c0bd395c610ed1e1/resolve`

- Checked: `2026-09-12T02:36Z` UTC on `mesh-home`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **operator-input diagnosis corrected; parent remains blocked on owner reconciliation of the experiment contract**

## Evidence

The owner-scoped dispatch check passed before the row was claimed by `adint`.
The earlier receipt [`unblock-adint-fde22f1d70bffa98-resolve-20260912.md`](../../../lte-workstation/docs/task-receipts/unblock-adint-fde22f1d70bffa98-resolve-20260912.md)
asked for five dictionary inputs; that diagnosis is superseded by the operator-authorized frozen
input recorded in
[`adint-operator-authorized-dictionary-input-20260912.md`](../../../tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md).
The operator directed us to choose from his repositories. The Tiny Fleet receipt records the
CC0 synthetic-only A08 fixture, English and Russian, exact case-sensitive replacement
normalization, immutable commit, and 152-row corpus. I independently recomputed its referenced
SHA-256 values; all six match that receipt:

| Artifact | SHA-256 |
|---|---|
| `runs/operator-selected-dictionary-v1/input-manifest.json` | `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08` |
| `runs/operator-selected-dictionary-v1/corpus.jsonl` | `81a41340c0c01555a536e9aa30b3b2b72b1dd9ae48908014fd0f58192515dcd3` |
| `corpus/applications/noisy-text-correction/manifest.json` | `17879f42f885b18ed5e0638a6620ff745200ddb79d9a3004a20fa22933ea0997` |
| `scripts/applications/noisy_text_correction.py` | `69330b7803f0844198c84f07d54cc89e822d4df9abc63ddfb58b6c3c921c60db` |
| `runs/operator-selected-dictionary-v1/result/raw.jsonl` | `555f1ab7808d5d01324c770e58fc594b48ea17ff55364d6a4214e4aed8b988fd` |
| `runs/operator-selected-dictionary-v1/bbywvy-smoke.txt` | `f64847a15bc68114b0d369200a0173d3f37e8a76c13fe499d93f9854eb4cffc6` |

The A08 fixture's measured baseline is 120/120 exact, but its own receipt correctly marks the
result `INCONCLUSIVE`: it covers five synthetic spelling substitutions, and the ByT5/LoRA arms
are unavailable. The separately hashed BbyWVY smoke is runtime-only, has six cases, and its script
contains no dictionary/lexicon arm. Thus neither result establishes the parent's requested
BbyWVY-compatible dictionary behavior or general-language quality.

## Parent gate and exact next action

`mesh-task status unblock/haunt/fd5d75268b44e1cc` reports the parent still `blocked` with
`operator-input`; `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt` exits `2`.
The five inputs now exist, so repeating the old request or running the six-case runtime smoke cannot
discharge this blocker.

The remaining prerequisite is an experiment-contract decision by `haunt`, the parent owner:

1. Accept the measured A08 synthetic-only result for the narrowly scoped prerequisite, explicitly
   revise the parent receipt/retry and its supported claim, and resume only if the owner's gate then
   accepts; or
2. If the parent requires a BbyWVY-compatible dictionary measurement, identify or implement that
   dictionary arm against the already frozen manifest and corpus, run it, and record its output
   hashes and typed verdict before resuming.

Adint did not alter or resume the parent row. No further operator input is missing. The exact next
action is for `haunt` to reconcile the accepted scope or provide the intended executable arm; the
parent owner gate remains authoritative.
