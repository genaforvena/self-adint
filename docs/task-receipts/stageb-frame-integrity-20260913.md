# Stage B canonical frame integrity audit

Checked: 2026-09-13 22:03 UTC  
Task: `adint-stageb-integrity-20260913/audit-stageb-frame-integrity`

## Result

The current canonical RU Stage B ledger and frame agree with each other and with a fresh
`tools/adint-status` render. The published README status block is stale by one row: it reports
208 domains and 45 `no-web-apex`; the current deduplicated ledger reports 209 domains and 46
`no-web-apex`. The other seven verdict counts agree. The README block says it was generated at
2026-09-13 13:12 UTC, before the current source snapshot.

`report.html` is a separate bidder-request experiment report, generated 2026-09-10 12:55 UTC;
it does not publish the Stage B verdict table, so it cannot corroborate those frame counts.

## Measurement

- Read-only `tools/adint-status` at 22:01 UTC: 209 canonical domains after collapsing 3
  duplicate rows; 14 admits.
- Independent read via `adint-frame-stageb.load_rows(..., dedupe=True)`: 212 schema-3 ledger
  rows, 3 duplicate rows collapsed, 209 unique domains. Verdict counts: 14 `admit`, 18
  `ad-serving-only`, 24 `no-wrapper`, 63 `ya-generic-only`, 23 `blocked`, 1 `rate-limited`,
  20 `unreachable`, and 46 `no-web-apex`.
- The canonical TSV contains 14 rows; its domain set exactly matches the ledger's 14 admitted
  domains. No domain names or device identifiers are copied into this receipt.
- `tools/adint-status --test`: PASS (`selftest: ok`).

## Snapshot hashes

```text
ledger  8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92
frame   b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3
README  bf11414c7ee9c48dcfe6e5d179bfe4da0176bc469fbbc90c3ffe03d5f9cc760e
report  b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f
```

No source ledger, frame, README, or report was changed. The next publication action is to
refresh the generated README block with `python3 tools/adint-status --write`, then verify its
counts against a fresh status render.
