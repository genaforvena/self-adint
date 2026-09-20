# Resolver receipt: `unblock/adint/d190e08ddf4f6ec0/resolve`

- Checked: 2026-09-12 UTC on `mesh-home`
- Parent: `haunt-install-unblock-20260907/install-and-retry-tinyfleet` (owner `haunt`)
- Result: **the queued `operator-input` diagnosis is stale; all five dictionary inputs are frozen, while the parent remains blocked on its owner-side experiment-contract decision**

## Evidence

The frozen input manifest remains present at
`/home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/input-manifest.json` with
SHA-256 `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`. The canonical
152-row corpus hashes to `81a41340c0c01555a536e9aa30b3b2b72b1dd9ae48908014fd0f58192515dcd3`;
the measured A08 dictionary output hashes to
`555f1ab7808d5d01324c770e58fc594b48ea17ff55364d6a4214e4aed8b988fd`; and the BbyWVY runtime
smoke hashes to `f64847a15bc68114b0d369200a0173d3f37e8a76c13fe499d93f9854eb4cffc6`. These match
the operator-authorized input receipt at
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`.
That receipt records the five frozen fields and a 120/120 exact A08 result with an
`INCONCLUSIVE` verdict for broader language behavior. It also states that the BbyWVY smoke is
runtime-only and contains no dictionary arm.

The earlier adint resolver receipt
`docs/task-receipts/unblock-adint-bff5a5ca18c0ac46-resolve-20260912.md` records the same evidence
and the distinction between missing input and missing compatible experiment evidence. A fresh
`mesh-task check resume haunt-install-unblock-20260907/install-and-retry-tinyfleet haunt` exited 2,
so the parent is not eligible for adint to resume. Its owner remains `haunt`.

## Exact next action

`haunt` should inspect the frozen manifest and A08 receipt, then settle the parent's criterion:
either narrow the parent claim to accept the synthetic-only A08 result with its `INCONCLUSIVE`
limit, or implement and measure a BbyWVY-compatible dictionary arm against the frozen inputs and
corpus. No further operator input or authorization is needed for that decision; the remaining
prerequisite is the parent's owner-side experiment-contract resolution. Adint has not resumed or
modified the haunt-owned task.
