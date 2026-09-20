# Live-state audit — 2026-09-12

## Input

`mesh-dash --once adint` refreshed at `2026-09-12T06:06:34Z` and ended at
`2026-09-12T06:06:41Z`. It reported no undischarged commitments and showed the line
`step0d-hb-first-cell — frame built (2 admits/69, density 2.9%)`.

## Independent read of the canonical frame

Ran `python3 tools/adint-status` from `~/self-adint` (read-only; no `--write`). At
`2026-09-12 06:07 UTC`, it derived `ru-mobile` as 14 admits across 208 touched domains and
`nl-direct` as 13 admits across 254 touched domains.

Then ran:

```text
python3 tools/adint-frame-compare --quiet \
  data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl \
  data/frame-stageb-nl-direct-2026-08-19-schema3.jsonl
```

It exited 0: the ledgers share 208 domains, with 175 agreements and 33 disagreements (15.9%);
the canonical RU frame has 14 admits, the NL frame 13, and their union 18.

## Disposition

The dashboard's `2/69` figure is not the same count as the current canonical Stage B ledgers.
At the time of the initial audit its source and age had not yet been identified; see the source
trace below. No data, frame, or operator state was changed.

## Source trace (2026-09-12 13:14 UTC)

Read the `adint` arm in `~/lte-workstation/scripts/mesh-dash`. The line comes directly from
`$HOME/.mesh/adint-step`; the dashboard only adds the file age and clips it for display. It does
not derive the numerator or denominator from the frame ledgers. The source file currently reads:

```text
step0d-hb-first-cell — frame built (2 admits/69, density 2.9%), first paired cell captured (6 pairs); walk continuing toward 20 admits
```

Its mtime is `2026-08-19 07:02:33 UTC`. A fresh read-only `python3 tools/adint-status` on
2026-09-12 13:14 UTC still derives the canonical RU frame as 14 admits / 208 touched domains and
the NL frame as 13 / 254. Thus the dashboard number is a dated, manually recorded progress
marker, not a current metric. The available source does not establish what its historical 69-domain
denominator counted, so it cannot be reconciled to or used to update the current study. Left both
the marker and study data unchanged.

## Next action

If the historical 69-domain denominator is needed, recover the exact source frame from the
2026-08-19 work artifacts before using it; do not infer it from today's canonical frame. The
canonical Stage B counts above are reproducible from the checked-in tools and ledgers.
