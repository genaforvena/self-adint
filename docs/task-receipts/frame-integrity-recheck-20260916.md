# Frame integrity recheck — 2026-09-16

Goal action: independently re-derive the checked-in canonical `ru-mobile` Stage-B frame for the
active `step0d-hb-first-cell` study. No source or data files were changed.

Commands and results:

```text
python3 tools/adint-status
  Generated 2026-09-16 00:41 UTC
  Vantage ru-mobile: 209 domains touched; 3 duplicate rows collapsed
  admit: 14
  Coverage: 239 of 4480 readable pairs (5.3%)

python3 tools/adint-status --test
  selftest: ok

python3 tools/adint-frame-stageb --test
  selftest: ok

JSONL schema check on public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
  rows=212; unique_domains=209; admits=14; schema=3; vantage=ru-mobile
```

The independently counted 14 admitted rows and 209 unique domains agree with the rendered status;
the study marker's 239/4480 readable-pair coverage is also reproduced. Acceptance passed.
