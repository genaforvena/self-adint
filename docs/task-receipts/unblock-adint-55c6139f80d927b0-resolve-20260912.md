# Resolver receipt: `unblock/adint/55c6139f80d927b0/resolve`

- Checked: `2026-09-12T02:51Z` UTC on `mesh-home`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **the stale operator-input gate is now open; parent is active under its owner**

## Verification

The exact-owner dispatch check passed (`mesh-task check dispatch ... adint`, exit 0), and the row
was claimed with `MESH_TASK_ACTOR=adint`. The live parent status is `active`, owned by `haunt`;
`mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt` now exits 0. The prior
operator-input diagnosis is superseded by the authorized frozen fixture and measurement recorded in
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`.

I rechecked all six artifact hashes against that receipt; each matched:

| Artifact | SHA-256 |
|---|---|
| `runs/operator-selected-dictionary-v1/input-manifest.json` | `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08` |
| `runs/operator-selected-dictionary-v1/corpus.jsonl` | `81a41340c0c01555a536e9aa30b3b2b72b1dd9ae48908014fd0f58192515dcd3` |
| `corpus/applications/noisy-text-correction/manifest.json` | `17879f42f885b18ed5e0638a6620ff745200ddb79d9a3004a20fa22933ea0997` |
| `scripts/applications/noisy_text_correction.py` | `69330b7803f0844198c84f07d54cc89e822d4df9abc63ddfb58b6c3c921c60db` |
| `runs/operator-selected-dictionary-v1/result/raw.jsonl` | `555f1ab7808d5d01324c770e58fc594b48ea17ff55364d6a4214e4aed8b988fd` |
| `runs/operator-selected-dictionary-v1/bbywvy-smoke.txt` | `f64847a15bc68114b0d369200a0173d3f37e8a76c13fe499d93f9854eb4cffc6` |

The A08 fixture remains synthetic-only and its measured result is explicitly inconclusive for a
general-language or BbyWVY-compatible dictionary claim. The six-case BbyWVY smoke is runtime-only;
it has no dictionary arm. No retry or change to the parent task is warranted from this duplicate
resolver, and only `haunt` may reconcile the parent's experiment contract. The exact next action is
for `haunt` to continue its now-eligible parent task and record its own result.
