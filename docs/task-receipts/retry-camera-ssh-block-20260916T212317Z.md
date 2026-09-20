# Camera SSH retry — typed external-event block — 2026-09-16T21:23:17Z

Task: `witness-64073-g7-camera-ssh/retry-camera-ssh` (owner `adint`).

## Live predicates personally checked

- `test -e /run/netns/ru-mobile`: **false** (`rc=1`); `ip netns list` returned no
  `ru-mobile` namespace.
- `ssh -o BatchMode=yes -o ConnectTimeout=5 -p 8022 u0_a380@100.103.99.16
  'termux-battery-status'`: **connection refused** (`rc=255`).
- `tailscale status --json` at 2026-09-16T21:20:56Z reported `Redmi 10
  100.103.99.16 online:false`.

These checks do not produce a camera artifact. No capture, router write, VPN change,
or repository data mutation was performed.

## Typed block and exact retry edge

Type: `external-event`.

Needs: the approved `/run/netns/ru-mobile` namespace with working RU-mobile egress,
and the Redmi SSH service reachable on `100.103.99.16:8022`.

Retry: after both predicates are true, run the canonical camera/paired retry, verify
the fresh reachable-frame receipt and both-arm vantage, then reconcile the expected
4,480 pairs. Until then the prior 239/4,480 state remains pilot-only, not complete.

Delegation: none; this was a tightly coupled single-owner live retry. I personally
inspected the prior namespace receipt, live Tailscale/namespace state, and this SSH
attempt; no subagent report was used as evidence.
