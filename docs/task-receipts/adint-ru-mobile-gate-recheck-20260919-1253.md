# ru-mobile gate recheck — 2026-09-19 12:53 UTC

This wake took one bounded action: rechecked the external prerequisite for the canonical
step-0d paired run. No capture, namespace provisioning, route change, router write, or
irreversible action was performed.

## Live evidence

- `test -e /run/netns/ru-mobile` exited `1`; the namespace is absent.
- `ip netns list` did not list `ru-mobile`.
- The RU egress probe was correctly not run because the namespace predicate was false.
- `python3 tools/adint-status --include-private` exited `0` at `2026-09-19 12:53 UTC`:
  canonical `ru-mobile`, 209 touched domains, 14 admits, 239/4480 readable pairs (5.3%),
  901/8960 attempted loads (10.1%).
- Current input hashes:

  ```text
  b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
  8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
  b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
  ```

## Ledger disposition

The exact-owner row `routing-shadow-resolver-followup-adint-20260914/review-frozen-gate-after-parent-completion`
is queued for `adint` but waits for
`self-review-routing-shadow-20260914/independently-evaluate-routing-shadow`; dispatch check
returned exit `2` (refused because the named prerequisite is not complete). No duplicate task
was created.

## Retry edge

When the approved `/run/netns/ru-mobile` exists, run `nsenter --net=/run/netns/ru-mobile ip -o
route get 1.1.1.1`, verify RU-mobile egress, then execute the already-owned canonical paired
measurement. Until then the gate remains `UNKNOWN`, not `NOT`.

The result was broadcast on the board and delivered to the operator through `mesh-voice-tx`; the
transport reported both voice-note sent and text duplicated.
