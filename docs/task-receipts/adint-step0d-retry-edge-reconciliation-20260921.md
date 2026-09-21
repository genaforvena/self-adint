# Step 0D retry-edge reconciliation — 2026-09-21

Task: `adint-goal-20260921-step0d-retry-edge-reconciliation/step0d-retry-edge-reconciliation`

## Live evidence

- `mesh-dash --once adint`: exit `0`; `pane:adint` rendered `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480`.
- `ip netns list`: exit `0`; no namespaces listed.
- `ip netns exec ru-mobile true`: exit `255`; `Cannot open network namespace "ru-mobile": No such file or directory`.
- `nc -vz -w 2 <redmi-tail-ip> 8022`: exit `0`; TCP connection succeeded from the host. This does not satisfy the namespace-scoped retry predicate by itself.
- `python3 tools/adint-status`: exit `0`; rendered the canonical `ru-mobile` frame with `14` admits and `239/4480` readable pairs.
- `python3 tools/adint-status --test`: exit `0`; `selftest: ok`.

## Disposition

The retry edge remains closed: the required `ru-mobile` namespace predicate is false, even though the host-level SSH port probe succeeds. No capture or routing change was attempted. Next retry: re-run the namespace probe and the canonical paired capture only after `ru-mobile` exists; preserve the SSH result as separate evidence.
