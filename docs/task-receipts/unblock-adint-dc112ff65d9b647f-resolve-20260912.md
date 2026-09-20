# Resolver receipt: `unblock/adint/dc112ff65d9b647f/resolve`

- Checked: `2026-09-12T03:30Z` UTC on `mesh-home`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **queued operator-input diagnosis is stale; parent is now blocked on owner-side experiment-contract reconciliation**

## Revalidation

The frozen dictionary manifest and measurement artifacts remain present. I recomputed the listed
SHA-256 values:

| Artifact | SHA-256 |
|---|---|
| `runs/operator-selected-dictionary-v1/input-manifest.json` | `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08` |
| `runs/operator-selected-dictionary-v1/corpus.jsonl` | `81a41340c0c01555a536e9aa30b3b2b72b1dd9ae48908014fd0f58192515dcd3` |
| `runs/operator-selected-dictionary-v1/result/raw.jsonl` | `555f1ab7808d5d01324c770e58fc594b48ea17ff55364d6a4214e4aed8b988fd` |
| `runs/operator-selected-dictionary-v1/bbywvy-smoke.txt` | `f64847a15bc68114b0d369200a0173d3f37e8a76c13fe499d93f9854eb4cffc6` |

The manifest pins source, languages, normalization, immutable source commit, and a 152-row corpus.
The measured A08 fixture result is synthetic-only and `INCONCLUSIVE` for broader behavior. The
six-case BbyWVY smoke is runtime evidence only and has no dictionary arm. Neither justifies
relabeling the A08 result as BbyWVY dictionary behavior.

`mesh-task status unblock/haunt/fd5d75268b44e1cc` reports the parent as `blocked` on
`experiment-contract`. `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`
exited `2`, so adint cannot resume that haunt-owned task.

## Exact next action

The `haunt` owner must reconcile the parent criterion: either narrow the claim and closure condition
to the frozen synthetic A08 result with its `INCONCLUSIVE` limit, or identify/implement and measure a
BbyWVY-compatible dictionary arm against the frozen manifest and corpus. Then `haunt` can update its
own task receipt and re-run its owner-scoped resume check. No additional operator input is missing;
adint did not modify or resume the parent task.

## Verification commands

```text
sha256sum /home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/input-manifest.json /home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/corpus.jsonl /home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/result/raw.jsonl /home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/bbywvy-smoke.txt
mesh-task status unblock/haunt/fd5d75268b44e1cc
mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt
```
