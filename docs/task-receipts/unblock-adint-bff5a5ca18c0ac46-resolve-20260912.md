# Resolver receipt: `unblock/adint/bff5a5ca18c0ac46/resolve`

- Checked: 2026-09-12 UTC on `mesh-home`
- Parent: `haunt-install-unblock-20260907/install-and-retry-tinyfleet` (owner `haunt`)
- Result: **the five dictionary input fields now have an operator-authorized frozen selection and a measured A08 synthetic arm; the parent remains blocked on experiment-contract interpretation, not missing operator input**

## Evidence

The previously missing source, language, normalization, version, and corpus are frozen in
`/home/mesh-home/tiny-fleet/runs/operator-selected-dictionary-v1/input-manifest.json`
(SHA-256 `9081dd78b28ed94b501016698cd6d5347095a5ea079730bf888123c3ec2e1e08`). The associated
operator authorization, exact field definitions, verification, and limits are recorded in
`/home/mesh-home/tiny-fleet/docs/task-receipts/adint-operator-authorized-dictionary-input-20260912.md`.
That receipt verifies the 152-row corpus and artifact hashes, a five-substitution A08 result of
120/120 exact with identity 0/120, and the verdict `INCONCLUSIVE` for broader language behavior.
It separately verifies the six-case BbyWVY smoke as runtime-only; that script has no dictionary
arm.

The original blocker evidence is
`/home/mesh-home/tiny-fleet/docs/task-receipts/haunt-install-unblock-20260908.md`, which records
the six-case BbyWVY retry and `BLOCKED_DICTIONARY_INPUT`. Thus the old missing-input diagnosis is
superseded, but the new A08 result does not establish a general or BbyWVY-compatible dictionary
measurement.

The live parent remains `blocked`, owned by `haunt`. I checked
`rtk mesh-task check resume haunt-install-unblock-20260907/install-and-retry-tinyfleet haunt`;
it exited 2. I did not claim or resume the parent from another mind's owner lane.

## Exact next action

`haunt` should inspect the frozen manifest and A08 result, then update its own receipt and resume
only if the synthetic-only A08 scope satisfies the parent's intended criterion. If the parent
requires a BbyWVY-compatible dictionary measurement, `haunt` must name and implement that dictionary
arm against the frozen inputs/corpus, then capture its result. No further operator input or
authorization is currently missing.
