# RU-mobile vantage preflight — 2026-09-19 20:07 UTC

## Verdict

**UNKNOWN / blocked before capture.** The canonical `/run/netns/ru-mobile` namespace is
absent, so the approved route and egress predicate cannot be evaluated. No namespace,
route, router, capture, or irreversible state was changed.

## Live evidence

- `ip netns list` returned no namespaces.
- `test -e /run/netns/ru-mobile` failed; the route probe was therefore skipped rather than
  treated as a success.
- `python3 tools/adint-status --include-private` succeeded at 20:07 UTC and rendered the
  current pilot boundary: **239/4480 readable pairs (5.3%)**, **901/8960 attempted loads
  (10.1%)**; canonical frame `ru-mobile`.
- Canonical self-tests passed: `adint-frame-stageb --test`, `adint-frame-compare --test`,
  `adint-paired-run --test`, and `adint-hb-report --test`.
- Current input hashes were read and recorded:

  ```text
  b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
  8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
  b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
  ```

## Retry edge

When `/run/netns/ru-mobile` exists, run:

```text
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Verify RU-mobile egress, then run the already-approved canonical paired measurement. Until
then, the study remains pilot-only and no market conclusion is valid.
