# Step 0d first-cell live reconciliation — 2026-09-16

Status: **pilot only; acceptance incomplete**.

## Evidence personally inspected

The canonical declaration in `ref/CANONICAL-FRAME` names `ru-mobile`, with source frame
`ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv` and ledger
`public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl`.

`python3 tools/adint-status --include-private` regenerated the current report from disk:

- canonical ru-mobile: 209 domains, 14 admits, 3 duplicate rows collapsed;
- readable paired coverage: **239 / 4480 (5.3%)**;
- therefore the first cell is not a completed 14-day study cell.

The historical paired files were directly parsed as JSONL with no parse errors:

- `data/hb-paired-2026-08-19-day-2026081914.nl.jsonl`: 1,047 records;
- `data/hb-paired-2026-08-19-day-2026081914.ru.jsonl`: 72 records;
- `data/frame-stageb-nl-direct-2026-08-19.jsonl`: 604 records;
- `data/frame-stageb-ru-mobile-2026-08-19.jsonl`: 561 records.

SHA-256 observed in this reconciliation:

```text
1ecf156450b8210c370de0450c49f04972f410d0b6dc5fe532bad39c57d828fa  data/hb-paired-2026-08-19-day-2026081914.nl.jsonl
995df89d4bdd0557425e9ac2551392f56db64e8b59340f87e7d4ac40882ec004  data/hb-paired-2026-08-19-day-2026081914.ru.jsonl
4e24f7c3cd109b383dbd33f92aac0dcc83a175a042b40cb457aa7fa00c5dd055  data/frame-stageb-nl-direct-2026-08-19.jsonl
5f6d82344eaee95654e76987340ef4a6f8a4384ce93984195ac45a092d93fd72  data/frame-stageb-ru-mobile-2026-08-19.jsonl
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
```

`tools/adint-frame-compare` over the two source JSONL ledgers exited 0. The offline self-tests
for `adint-frame-stageb`, `adint-frame-compare`, `adint-paired-run`, `adint-hb-report`,
`adint-publish`, and `adint-status` all exited 0. `adint-publish --test` passed its privacy
allowlist gate. `adint-hb-capture --test` exited 2 because Playwright is absent from the system
interpreter; the documented browser virtualenv is required for that capture check.

The broad ad-hoc token scan is **not** used as a clean result: it matches an intentional
`ref/sdk-signatures.tsv` documentation row and tracking identifiers in historical public URLs.
The authoritative publish self-test passed; no new data was written by this reconciliation.

## Acceptance gap and next gate

The narrowest missing evidence is a clean post-fix paired run on the canonical ru-mobile frame,
with both-arm vantage verification, no concurrent frame/ledger writers, and completion of the
planned 4,480 readable pairs. No seat onboarding, payment, GAID reset, router write, or external
contact is authorized or performed here.

Delegation: `adint-step0d-audit` performed a read-only independent audit. I personally inspected
the cited files and reran the counts, hashes, parse, comparison, and self-tests above; its report
was not treated as evidence by itself.
