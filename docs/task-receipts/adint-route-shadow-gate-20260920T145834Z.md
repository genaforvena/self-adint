# Route-shadow gate preflight — 2026-09-20T14:58:34Z

Task: `adint-goal-20260920-route-shadow-gate-preflight/route-shadow-gate-preflight`

## Live predicate

`test -e /run/netns/ru-mobile` returned exit `1`: the approved `ru-mobile` network
namespace is absent. `ip netns list` returned no entries.

Verdict: `UNKNOWN` — route-shadow verification cannot be performed from this node
until the namespace exists. The paired heartbeat capture was not run.

## Retry edge

After the namespace is present, run:

```bash
test -e /run/netns/ru-mobile
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Only if that route shows the approved RU-mobile egress, run the canonical paired
capture:

```bash
python3 tools/adint-paired-run \
  --frame ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv \
  --replicates 1 --window 45 --seed 20260819
```

## Input identity

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
```

The receipt itself is the bounded artifact for this preflight; it does not claim
that Step 0d is measured or that the route gate passed.
