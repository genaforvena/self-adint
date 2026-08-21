# The two arms do not share an uplink, and that is why a link fault cannot fake a refusal

*2026-08-21. Measured on mesh-home while the node's wifi dongle was mid-fault.*

## The measurement

The paired design runs both arms from one machine, which invites the assumption that they
share a network path. They do not. Measured directly:

| arm | namespace | interface | default gateway |
|---|---|---|---|
| `ru-mobile` | `ruvantage` netns | `enx0257356d3666` (Note 3 USB tether) | 192.168.42.129 |
| foreign (`nl-direct` / `us-exit`) | host | `wlxbcec43434a22` (`rtw_8822bu` USB wifi) | 192.168.8.1 |

Two USB devices, two carriers, two independent failure processes. "Our own link failed" has
never been one claim in this study; it has always been two, and nothing said so.

## Why it matters, in both directions

**It strengthens the access finding.** The refusal result rests on the claim that a 4xx
answered by a publisher is a decision about the vantage and not an artefact of our
instrument. If the arms shared a link, a fault on it could plausibly produce asymmetric
outcomes and a shared-cause confound would be live. They do not share one — so a fault on
either uplink can only produce a load that never returned, which `render_verdict` classifies
as `pair-gate:no-load` and never as `pair-gate:refused`. **A link fault cannot manufacture a
refusal**, by construction rather than by luck.

**It explains the arm-asymmetric blindness** that the gate table reports. `ru-mobile` clears
the render gate on 62 % of its loads against the foreign arms' 94–97 %, and that is not a
mystery about vantages: it is a mobile tether against a wifi dongle, two different devices
with two different fault shapes.

**And it makes the losses time-correlated within an arm, not across arms.** This node's sole
wifi uplink is a `rtw_8822bu` dongle with a documented fault history — it wedged for 55
minutes on 2026-08-17, and while writing this it was logging `timed out to flush queue 1`
and `failed to get tx report from firmware`. During such an episode the foreign arm loses
loads while the domestic arm, on a different device entirely, keeps working. That costs
**power** in the access contrast, by shrinking the count of comparable `(cell, site)` units.
It does not bias the direction.

## What this does not license

It does not make either arm's blindness harmless, and it does not retire the coverage
caveat: the analysis set is still conditioned on both fragile links having been up at the
same moment, and that conditioning is still eight times stronger on one side. What it
retires is one specific alternative explanation — *the instrument's own network produced the
asymmetry* — which is now ruled out by the topology rather than argued away.

---

## A live case, the same day: a tailscale outage took one arm and not the other

The claim above stopped being theoretical within hours. Timeline, from the journal and the
cell's own run log:

| time (UTC) | event |
|---|---|
| 11:40:45 | `tailscaled` loses the control plane — `PollNetMap: use of closed network connection`, then continuous `all connection attempts failed` to `controlplane.tailscale.com` |
| 12:01–12:47 | `nl-direct` refuses **23 consecutive loads**: `preflight vantage: via=direct exit=None echo_attempts=3/3` |
| 12:47:39 | `linkChange: in state Running; updating LAN routes`; `logtail: upload succeeded after 62 failures` |
| 12:48 | `nl-direct` returns `rc0` |

This node's egress **is** phaedra's tailscale exit node, so while tailscale could not reach
its control plane the foreign arm could not read an echo and correctly refused to record a
load it could not place. The domestic arm, on the Note 3 tether and outside tailscale
entirely, ran **28 / 28** through the whole window. The cell ended `ru: 28, nl: 5`.

Two things worth taking from it.

**The independence is real and it cuts both ways.** It is why a link fault cannot fake a
refusal — and also why one arm can lose 47 minutes while the other notices nothing. A
single-uplink design would have lost the cell entirely and made the outage obvious; this
one loses half a cell quietly.

**The loss is CONTIGUOUS, which is worse than random.** An outage removes a solid block of
wall-clock from one arm, not a scatter of loads. If anything this study measures varies by
time of day — and a bidder's willingness to price plainly might — then losing 12:01–12:47
from the foreign arm is not the same as losing 23 loads spread evenly. It is a small
time-of-day hole in one arm only. The exclusion counts show the magnitude; they do not show
that shape, and no aggregate over surviving pairs will reveal it either.
