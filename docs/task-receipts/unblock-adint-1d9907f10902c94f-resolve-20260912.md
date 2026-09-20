# Resolver audit: `unblock/adint/1d9907f10902c94f/resolve`

- Date: 2026-09-12 UTC
- Owner: `adint`
- Disposition: **REJECTED — live task, stale and unsafe instructions**
- Related parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)

## Current state

The resolver was live and owner-assigned when checked: `mesh-task status
unblock/adint/1d9907f10902c94f` reported its only step `active`, owner `adint`.
The parent remains blocked, but its current typed blocker is
`experiment-contract`, not `operator-input`. Its retry text requires the haunt
owner to narrow the claim to the measured A08-only result or produce a
BbyWVY-compatible dictionary arm. `mesh-task check resume
unblock/haunt/fd5d75268b44e1cc/resolve haunt` exited 2, as expected while that
owner's criterion is unresolved.

The dispatch description is therefore stale in its asserted missing-input
reason and its proposed retry. The five dictionary fields are already frozen
in Tiny Fleet's `runs/operator-selected-dictionary-v1/input-manifest.json`
(SHA-256 `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`).
The measured A08 synthetic result is 120/120 exact on the heldout fixture, but
its own verdict is `INCONCLUSIVE` for broader correction behavior and it does
not measure BbyWVY. The detailed source, corpus, method, and output hashes are
in `tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`.

The prescribed retry, `scripts/bbywvy_test.py`, is also mis-specified as
dictionary evidence. Inspection at Tiny Fleet HEAD
`3d687ea664d9b638cd9f8755457d8239bb537421` shows six generic generative prompts
and no dictionary arm. Its own current experiment contract explicitly says not
to rerun it as a dictionary test. The matching contract is
`/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md`.

## Disposition

No source or experiment state was changed. Following the stale retry would
produce runtime-smoke evidence that cannot discharge the current blocker. The
safe next action belongs to `haunt`: amend the parent claim and closure gate to
the narrow A08-only statement with its `INCONCLUSIVE` limit, or implement and
measure a compatible BbyWVY dictionary arm against the frozen manifest and
corpus. Until then, the parent remains blocked and its resume gate should
continue to refuse.

## Verification

- `mesh-task status unblock/adint/1d9907f10902c94f`: resolver active at audit.
- `mesh-task status unblock/haunt/fd5d75268b44e1cc`: parent blocked on `experiment-contract`.
- `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`: exit 2.
- Frozen input manifest SHA-256: `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`.
- Tiny Fleet revision inspected: `3d687ea664d9b638cd9f8755457d8239bb537421`; its working tree had pre-existing unrelated modifications, preserved.
- `scripts/bbywvy_test.py` SHA-256: `377521b04612de877b18caf046157d772293e23539f8d11817ee771cad598a8a`; source inspection confirms generic chat prompts, no dictionary input or scoring arm.
