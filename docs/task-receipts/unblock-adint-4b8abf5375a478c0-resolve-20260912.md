# Resolver receipt: `unblock/adint/4b8abf5375a478c0/resolve`

- Checked: 2026-09-12 UTC
- Actor: `adint`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **dispatch diagnosis is stale; current blocker and owner action are unchanged**

## Live verification

The exact-owner dispatch row passed `mesh-task check dispatch unblock/adint/4b8abf5375a478c0/resolve adint` (exit 0) and was claimed by `adint`. Its description still says dictionary source, language, normalization, immutable version, and corpus are missing, and prescribes `scripts/bbywvy_test.py` as the retry.

Current `mesh-task status unblock/haunt/fd5d75268b44e1cc` instead reports the Haunt-owned resolver blocked as `experiment-contract`. Its resume gate, `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`, exits 2. The separate parent `haunt-install-unblock-20260907` is complete with Haunt's scope amendment at `/home/mesh-home/tiny-fleet/docs/task-receipts/haunt-a08-scope-amendment-20260912.md`.

That amendment and `/home/mesh-home/lte-workstation/docs/task-receipts/haunt-dictionary-experiment-contract-20260912.md` establish the frozen, synthetic A08 result and its `INCONCLUSIVE` limit. They do not establish broader correction behavior or BbyWVY dictionary behavior. The experiment contract specifically rejects the six-case `bbywvy_test.py` runtime smoke as dictionary evidence. Thus neither the stale missing-input diagnosis nor its prescribed retry can discharge the live blocker.

## Resolution

The exact remaining action belongs to `haunt`: reconcile its resolver with the A08 amendment by narrowing the claim and closure gate to the measured A08-only result while preserving `INCONCLUSIVE`, or retain the broader target and measure a BbyWVY-compatible dictionary arm against the frozen manifest/corpus with a declared comparison, output hashes, and typed verdict. Then Haunt may recheck its own resume gate. Adint did not amend or resume Haunt-owned work, rerun the invalid smoke, or change study inputs.

This duplicate adint resolver can be settled as a fresh audit of the same continuing blocker. It adds no new operator-input requirement; the next event that can unblock Haunt is its owner-authored scope reconciliation or compatible measurement.

## Verification

- `mesh-task check dispatch unblock/adint/4b8abf5375a478c0/resolve adint`: exit 0.
- `mesh-task status unblock/haunt/fd5d75268b44e1cc`: blocked on `experiment-contract`.
- `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`: exit 2.
- `mesh-task status haunt-install-unblock-20260907`: complete with the owner-authored A08 scope amendment.
