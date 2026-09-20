# RU-mobile gate recheck — 2026-09-19 14:55 UTC

This receipt records one bounded wake action: a fresh read of the canonical RU-mobile gate.
No capture, namespace provisioning, route change, router write, or irreversible action was
performed.

## Live evidence

- `test -e /run/netns/ru-mobile` exited `1`; the namespace is absent.
- `ip netns list` returned no namespaces, including `ru-mobile`.
- The route probe was skipped because the namespace predicate was false.
- `python3 tools/adint-status --include-private` regenerated the live status at `2026-09-19 14:55 UTC`:
  canonical `ru-mobile`, `209` touched domains, `14` admits, `239/4480` readable pairs (`5.3%`),
  `901/8960` attempted loads (`10.1%`).
- Input hashes remained:

  ```text
  b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
  8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
  b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
  ```

## Disposition

The canonical paired measurement remains gated `UNKNOWN`: `ru-mobile` is not present, so the
approved namespace route check and paired run cannot be attempted. The exact-owner resolver row
also remains queued behind blocked prerequisite
`self-review-routing-shadow-20260914/independently-evaluate-routing-shadow`.

## Retry edge

When `/run/netns/ru-mobile` exists, run
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`, verify RU-mobile egress, then
execute the already-owned canonical paired measurement. Recheck the resolver prerequisite after
its recorded retry time, `2026-09-28T16:52:06Z`.
