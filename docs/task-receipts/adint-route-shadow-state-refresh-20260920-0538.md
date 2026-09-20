# Route-shadow / RU-mobile state refresh — 2026-09-20 05:38 UTC

Task: `adint-goal-20260920-route-shadow-state-refresh-0534/route-shadow-state-refresh-0534`

## Live evidence

- `mesh-dash --once adint` exited 0 and rendered on `pane:adint`:
  `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480 readable pairs`.
- `python3 tools/adint-status --include-private` exited 0 and regenerated the status from disk:
  canonical `ru-mobile`, 209 touched domains, 14 admits, 239/4480 readable pairs (5.3%),
  and 901/8960 attempted loads (10.1%).
- `ip netns list` returned no namespaces.
- `test -e /run/netns/ru-mobile` exited 1; the approved namespace is absent.
- The route probe was not run because its namespace predicate was false.
- No capture, namespace provisioning, router write, route change, or irreversible action was
  performed. Existing unrelated worktree changes were preserved.

## Verdict and retry edge

The canonical paired measurement remains `UNKNOWN`, held before capture. When
`/run/netns/ru-mobile` exists, run:

```text
nsenter --net=/run/netns/ru-mobile ip -o route get 1.1.1.1
```

Only after that route confirms the RU-mobile egress may the already-owned paired measurement
run. The receipt is observed on `pane:adint`; the live pane still shows the same held gate.

## Delegation record

Carver (`01a0bd4f-a372-7683-984d-0ed13e432958`) performed a read-only inspection. I personally
inspected its cited `docs/step0d-hb-first-cell-2026-08-19.md`, `ref/CANONICAL-FRAME`, and prior
gate receipt, then reran the live commands above. The delegated report was not used as standalone
evidence.
