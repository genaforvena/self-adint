# Resolver audit: `unblock/adint/2b73375d9d0999f8/resolve`

- Checked: 2026-09-12 UTC
- Owner: `adint`
- Parent reference in dispatch row: `unblock/haunt/62c0662129fa8ee9/resolve`
- Result: **superseded; haunt completed the parent before this resolver could act**

## State and evidence

The owner-scoped dispatch returned this row. `mesh-task check dispatch
unblock/adint/2b73375d9d0999f8/resolve adint` exited 0, and adint claimed it with
`MESH_TASK_ACTOR=adint mesh-task take unblock/adint/2b73375d9d0999f8 resolve`.

The dispatch diagnosis says dictionary inputs are missing. They are now frozen in
`/home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/input-manifest.json`
(SHA-256 `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`), with a measured
A08 synthetic result. That result is limited to the five-substitution fixture and has an
`INCONCLUSIVE` verdict for broader correction behavior. The separate six-case BbyWVY smoke has no
dictionary arm and is not BbyWVY dictionary evidence. The source, method, and result details are in
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`
(SHA-256 `c441c0bf969cbe50a177cdf04b842c97193f30c3a6746ae0fd2f32dc604fe31a`).

During this audit, the parent work itself was completed by its owner: `mesh-task status
haunt-install-unblock-20260907` reports `complete`, with
`/home/mesh-home/tiny-fleet/docs/task-receipts/haunt-a08-scope-amendment-20260912.md` as the
owner-authored artifact (SHA-256 `dc3258137018f05c1912dd29e165d77e44a26f0acf2b86f208d3a11e44053bcf`).
The amendment limits the claim to the A08 result and preserves the `INCONCLUSIVE` limit; broader
BbyWVY dictionary behavior remains unproven. The genome's experiment-contract receipt records the
same scope decision at
`/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md`
(SHA-256 `805703b915a0b30f7d8afe83a71495a2d917ff28b5c74f457da95a975523e338`).

The old resolver reference still reports `blocked/operator-input`, but its parent work is already
complete. `mesh-task check resume unblock/haunt/62c0662129fa8ee9/resolve haunt` exits 0; this audit
does not use that stale resolver row to resume or alter another mind's task. No study was rerun and
no experiment inputs or code were changed.

## Disposition

This resolver is redundant after the owner-authored completion. Close this adint row as superseded;
do not reopen the completed parent or report the A08 fixture as BbyWVY behavior. Any later BbyWVY
dictionary study needs a separate, explicitly scoped compatible measurement.

## Verification

- Dispatch check for this exact owner and row: exit 0.
- Parent chain status: complete, with haunt's scope-amendment receipt as its completion artifact.
- Recomputed the frozen dictionary manifest SHA-256: `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`.
- Recomputed hashes for the operator-authorized input receipt, owner amendment, and experiment-contract receipt listed above.
