# Receipt: review-adint-readme-stale-20260916/refresh-stageb-readme

Run: 2026-09-16T18:23:24Z–18:23:45Z UTC

## Eligibility and before state

- `mesh-task check dispatch review-adint-readme-stale-20260916/refresh-stageb-readme adint` passed.
- The task ledger row was already active for owner `adint` (`mesh-task status review-adint-readme-stale-20260916`); `mesh-task take review-adint-readme-stale-20260916 refresh-stageb-readme` confirmed `already active`.
- Before README hash: `bf11414c7ee9c48dcfe6e5d179bfe4da0176bc469fbbc90c3ffe03d5f9cc760e`.
- Before canonical inputs: `public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl` = `8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92`; `data/frame-stageb-ru-mobile-2026-08-19.jsonl` = `5f6d82344eaee95654e76987340ef4a6f8a4384ce93984195ac45a092d93fd72`; `ref/CANONICAL-FRAME` = `05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf`.
- Writer lock metadata before/after remained unchanged: `data/frame-stageb-ru-mobile-2026-08-19.jsonl.lock` size 8, inode 7473614, mtime `2026-08-19 13:56:53.288388373 +0000`; `public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl.lock` size 8, inode 7511246, mtime `2026-09-13 13:45:08.075057967 +0000`.

## Mandated writer output

Command:

```text
python3 tools/adint-status --write
```

Output:

```text
README.md numbers block rewritten (123 lines)
```

After README hash: `c3730e1e16ee637cfda88b661de9b6862733084cf00cd0f0862760c339473d1b`.

The generated canonical Stage B block reports 209 domains and these ru-mobile verdict counts:

```text
admit=14  ad-serving-only=18  no-wrapper=24  ya-generic-only=63
blocked=23  rate-limited=1  unreachable=20  no-web-apex=46
```

## Independent verification

Commands run after the write:

```text
python3 tools/adint-status --test
python3 tools/adint-status
```

Result: `selftest: ok`; all individual checks printed `ok`.

The independent fresh render's normalized hash was
`35685dd1f08f27fa56b9359aa0cae7871d02f488606675c8a2bc03ad49bed4f4`, matching the README block's
normalized hash exactly. The normalization removes only the intentionally changing generated UTC
timestamp line. Exact comparison result: `normalized README block == fresh render: True`.

Final unchanged input hashes: public ledger `8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92`, private ledger `5f6d82344eaee95654e76987340ef4a6f8a4384ce93984195ac45a092d93fd72`, canonical frame `05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf`.

Observed top pane: `pane:adint` is live and passed `mesh-pane-check` (`PASS: adint`; final output `pane-check: clean`).
