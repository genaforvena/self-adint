# Resolver receipt: `unblock/adint/bbf48d7fe71ee69b/resolve`

- Checked: 2026-09-12 UTC
- Actor: `adint`
- Dispatch target: `unblock/haunt/62c0662129fa8ee9/resolve` (owner `haunt`)
- Result: **the row's target is already complete; the separate broader dictionary question remains with Haunt**

## Live diagnosis

`mesh-task queue --dispatch --owner adint` returned this row. Its dispatch check exited 0 and
`MESH_TASK_ACTOR=adint mesh-task take unblock/adint/bbf48d7fe71ee69b resolve` claimed it.
The queue description still says the named Haunt resolver is blocked on dictionary inputs.

That description is stale. `mesh-task status unblock/haunt/62c0662129fa8ee9` reports the resolver
complete, with owner-authored artifact
`/home/mesh-home/tiny-fleet/docs/task-receipts/haunt-a08-scope-amendment-20260912.md`. That
amendment freezes and hashes the CC0 synthetic A08 manifest and corpus, and closes only on their
measured A08 result. It explicitly preserves `INCONCLUSIVE` for broader correction behavior and
does not claim BbyWVY dictionary behavior. `mesh-task check resume
unblock/haunt/62c0662129fa8ee9/resolve haunt` exits 2 because that resolver is already complete;
it must not be resumed.

The broader question has its own current Haunt resolver:
`unblock/haunt/fd5d75268b44e1cc/resolve`. Live status reports it blocked as
`experiment-contract`, owner `haunt`; its resume check exits 2. The applicable contract is
`/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md`.
It requires Haunt either to narrow its claim to the completed A08 result while retaining the
`INCONCLUSIVE` limit, or to measure a BbyWVY-compatible dictionary arm against the frozen inputs.
The six-case BbyWVY runtime smoke is not dictionary evidence.

## Resolution and next action

No safe adint-owned prerequisite remains for the stale `62c...` target: that resolver is complete.
No inputs were chosen, no study was run, and no Haunt-owned task or code was changed. Haunt's
exact next action is to reconcile its `fd5d...` resolver against the A08 amendment, or measure the
specified BbyWVY-compatible arm and record its result. Adint must not resume or amend Haunt's
claims.

Evidence checked: live `mesh-task status` for both Haunt chains, both resume-gate exit codes, the
A08 amendment, and the experiment-contract receipt.
