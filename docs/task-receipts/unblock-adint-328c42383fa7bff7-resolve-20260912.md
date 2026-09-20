# Resolver receipt: `unblock/adint/328c42383fa7bff7/resolve`

- Checked: `2026-09-12T03:08:27Z` UTC on `mesh-home`
- Parent: `unblock/haunt/62c0662129fa8ee9/resolve` (owner `haunt`)
- Dispatch: exact-owner dispatch check exited `0`; claimed by `adint`
- Result: **operator-input diagnosis is stale; parent needs its owner's experiment-contract reconciliation**

## Live gate and evidence

The parent remains blocked, owner `haunt`, with retry text requesting frozen dictionary inputs.
The live owner resume gate, `mesh-task check resume
unblock/haunt/62c0662129fa8ee9/resolve haunt`, exited `2`.

Those inputs have since been selected and measured under the operator's instruction to choose
from his repositories. The frozen selection and its result are recorded in
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`
(SHA-256 `c441c0bf969cbe50a177cdf04b842c97193f30c3a6746ae0fd2f32dc604fe31a`). The result is
limited to the pinned synthetic A08 fixture: 120/120 exact outputs, with an explicit
`INCONCLUSIVE` verdict for broader correction behavior. It does not establish BbyWVY dictionary
behavior; the six-case BbyWVY smoke has no dictionary arm.

The current contract analysis is
`/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md`
(SHA-256 `805703b915a0b30f7d8afe83a71495a2d917ff28b5c74f457da95a975523e338`). It specifies the
two valid owner paths: narrow the parent claim and closure criterion to the A08-only result and
retain its `INCONCLUSIVE` limit, or provide/implement a BbyWVY-compatible dictionary arm against
the frozen inputs and measure it before resuming.

## Exact owner action

`haunt` should update its own parent receipt and task criterion to choose one of those two paths,
replace the stale `operator-input` blocker/retry with the remaining experiment-contract condition,
and run the owner resume check. Do not resume or edit the `haunt`-owned row from `adint`. Until
that reconciliation, the exact gate above correctly refuses resume; no further operator datum is
needed.
