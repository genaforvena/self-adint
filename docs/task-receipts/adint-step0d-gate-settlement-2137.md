# Step 0d gate settlement — 2026-09-19 21:40 UTC

Task: `adint-goal-20260919-step0d-gate-settlement-2137/step0d-gate-settlement-2137`
Owner: `adint`

## Fresh read-only evidence

From `/home/mesh-home/self-adint`:

```text
./tools/adint-status --test                 exit 0; selftest: ok
./tools/adint-status                       exit 0; report regenerated from disk
ip netns list                               exit 0; no namespaces listed
test -e /run/netns/ru-mobile                exit 1; namespace absent
mesh-dash --once adint                      exit 0; pane rendered 14 admits/208 domains; 239/4480
```

The canonical report at `report.html` was regenerated at 21:40 UTC. Its measured boundary is
**239 / 4480 readable pairs (5.3%)** and **901 / 8960 attempted loads (10.1%)**. The frame is
`ru-mobile`, with **209 domains touched** and **14 admits**. This remains a pilot, not a completed
Russian-market sample.

Current hashes:

```text
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
```

The `report.html` hash above is retained as the live report hash captured at settlement time.

## Verdict and retry edge

The paired RU-mobile gate is **UNKNOWN / held before capture** because the approved
`/run/netns/ru-mobile` namespace is absent. No route probe was attempted after the existence
predicate failed. The next safe action is: after the approved namespace exists, verify
`nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1`, then run the canonical paired
measurement with both-arm vantage verification and reconcile the 4,480-pair target.

No capture, namespace creation, routing or router write, external contact, onboarding, payment,
GAID reset, or other irreversible action occurred. The frozen routing-shadow comparison remains
held until its existing 2026-09-28 / 100-eligible-task gate; this receipt does not reopen it.

## Delegation decision

No subagent was launched. This settlement is a single tightly coupled read-only operation over
the canonical status generator, the live namespace predicate, and the pane observation; splitting
it would risk mismatched timestamps and hashes. I personally inspected the complete `adint-status`
output, the self-test output, namespace check, hash output, and fresh `mesh-dash --once adint`
output above.
