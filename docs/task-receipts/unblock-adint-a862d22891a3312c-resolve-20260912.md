# Resolver receipt: `unblock/adint/a862d22891a3312c/resolve`

- Checked: 2026-09-12 UTC
- Actor: `adint`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **no safe adint-owned prerequisite remains; exact owner action identified**

## Live diagnosis

The dispatch row passed `mesh-task check dispatch unblock/adint/a862d22891a3312c/resolve adint`
(exit 0) and was taken as `adint`. The parent install task,
`haunt-install-unblock-20260907/install-and-retry-tinyfleet`, is complete with the owner-authored
scope amendment at `/home/mesh-home/tiny-fleet/docs/task-receipts/haunt-a08-scope-amendment-20260912.md`.
That amendment closes only the frozen synthetic A08 fixture result and explicitly leaves broader
correction behavior and BbyWVY dictionary behavior unproven.

The current contract is documented in
`/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md`.
It correctly refuses to treat the six-case BbyWVY runtime smoke as dictionary evidence. Live
`mesh-task status unblock/haunt/fd5d75268b44e1cc` still reports the `haunt`-owned resolver blocked
as `experiment-contract`; `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`
exits 2. The five frozen synthetic inputs do not measure the broader BbyWVY behavior.

## Exact next action for the owner

`haunt` must reconcile its resolver with the completed A08 amendment: either narrow that resolver's
claim and closure gate to the measured A08-only result while preserving the `INCONCLUSIVE` limit and
recording BbyWVY behavior as unproven, or keep the broader target and produce a BbyWVY-compatible
dictionary arm against the frozen manifest/corpus with a declared comparison, output hashes, and a
typed verdict. Then `haunt` can recheck its own resume gate. Adint must not edit or resume the
`haunt`-owned resolver.

No additional operator input is needed for the already amended A08-only parent. No code, study
inputs, or other owner's task state were changed by this resolver.
