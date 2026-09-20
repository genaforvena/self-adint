# Unblock receipt — `unblock/adint/746349c3a536c954/resolve`

Recorded: 2026-09-16T09:08:16Z

## Result

The mesh-internal GPU prerequisite is satisfied. Three separate fresh
`nvidia-smi` readings all show 9004 MiB free of 12288 MiB, above the required
5600 MiB threshold:

```text
2026-09-16T09:08:12Z  free=9004 MiB total=12288 MiB
2026-09-16T09:08:14Z  free=9004 MiB total=12288 MiB
2026-09-16T09:08:16Z  free=9004 MiB total=12288 MiB
threshold: free >= 5600 MiB on all three reads — PASS
```

The remaining prerequisite is external and still unsatisfied: an independent
human must review and provenance all 8 records in
`/home/mesh-home/finnegans-fake/wake/multistream-target.jsonl`. The existing
structural review explicitly says `independent_human_annotation=false` and
`H1 STILL UNMEASURED`; no annotation was fabricated and training was not run.

## Exact retry edge

After an independent human adds reviewer identity, review method, review date,
and a per-record judgment/provenance for all 8 records, with
`independent_human_annotation=true` only when that review occurred, rerun:

```text
cd /home/mesh-home/finnegans-fake
python3 wake/train_multistream.py --seed 1
```

The GPU gate must be rechecked with three separate readings immediately before
the retry.

## Evidence inspected

- `/home/mesh-home/self-adint/docs/task-receipts/unblock-adint-2ee3384cb9e14277-resolve-20260916.md`
  — prior receipt with the same external gate.
- `/home/mesh-home/finnegans-fake/docs/wake-multistream-annotation-review-2026-09-16.md`
  — structural PASS, but independent human annotation absent.
- Fresh live `nvidia-smi` readings recorded above.
