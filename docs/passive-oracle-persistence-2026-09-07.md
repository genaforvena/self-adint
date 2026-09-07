# Passive baseline and oracle persistence — 2026-09-07

The authoritative artifact is `data/passive-oracle-2026-09-07/`. It is generated offline by
`tools/adint-passive-oracle` from the append-only passive ledger and the 2026-09-07 passive corpus.

## Baseline

- request corpus: schema 7, 3,987 request rows;
- load rows: 48;
- priced request rows: 2;
- passive ledger: 116 rows — 92 `collected`, 24 `collected-nothing`;
- egress agreement in the ledger: 92 `true`, 24 `null` (the latter are not treated as agreement);
- source digest: recorded in `schema-egress-report.json` and both run rows.

This is browser passive evidence. It is not a phone/device capture and not a seat-side payload.

## Oracle eligibility

`oracle-verdict.json` is `NOT_ELIGIBLE`, with `active_result: null`. The passive browser corpus does
not establish the required device/seat payload-empty-or-hashed gate. No active bid, targeting
query, injection, payment, or GAID operation was attempted.

## Persistence check

`events.jsonl` and `runs.jsonl` contain one passive treatment replay and one identical control
replay. Both carry the same source SHA-256; each event carries a SHA-256 signature over its
canonical fields. `persistence-report.json` records equal inputs, valid signatures, zero network
calls, and no active result. `SHA256SUMS` verifies every artifact in the directory.

The exact next gate is a measured device/seat payload verdict. Until it exists, the active oracle
must remain held and any `NOT`/`IN` claim would be fabricated.
