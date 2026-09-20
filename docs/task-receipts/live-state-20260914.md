# Live-state receipt — 2026-09-14

## Read

`mesh-dash --once adint` showed the active study step `step0d-hb-first-cell`, with the
`ru-mobile` frame and a progress marker of 14 admits / 208 domains and 239 / 4480 readable
pairs. It also showed the unresolved resolver `unblock/adint/3dd6562eb2e7cc86/resolve`.

`mesh-task queue --dispatch --owner adint` exited 0 with no eligible owned rows. The resolver
ledger remains blocked on the room-camera capability; its retry condition is successful bounded
SSH reachability, followed by a real read from `mesh-imac-cam --test`.

## Fresh derivations and retry

`python3 tools/adint-status` (read-only) derived 14 canonical `ru-mobile` admits across 209
touched domains, with 3 duplicate rows collapsed. Its paired-cell section reports 239 / 4480
readable pairs (5.3%). These are current corpus derivations; they do not rewrite the older
dashboard progress marker or establish its historical 69-domain denominator.

On 2026-09-14, the exact bounded SSH retry to the configured iMac endpoint exited 255 with
`Connection timed out`. The retry condition therefore remains unsatisfied and the real camera
read was not attempted. No camera service or mesh substrate was changed.

## Disposition

Keep `unblock/adint/3dd6562eb2e7cc86/resolve` blocked as a capability dependency. Retry when
`timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true` exits 0, then
require `mesh-imac-cam --test` to complete a real read before resuming. The study step remains
active; no frame, study data, or progress marker was modified.

## Consume refresh — 2026-09-14 06:29 UTC

Ran `mesh-dash --once adint`. It still shows `step0d-hb-first-cell` active, the dashboard marker
at 14 admits / 208 domains and 239 / 4480 readable pairs, plus the unresolved resolver above.
The fresh read-only `python3 tools/adint-status` derivation at 06:28 UTC confirms 14 canonical
admits / 209 touched domains (3 duplicate rows collapsed) and 239 / 4480 readable pairs.

Retried the resolver's bounded SSH condition at 06:27 UTC; it again exited 255 with
`Connection timed out`. A fresh `mesh-task queue --dispatch --owner 'adint'` exited 0 and
returned no eligible owned rows. No `mesh-task check` or `take` was applicable. The dependency
remains blocked. No frame or study rows, camera, or substrate settings were changed; this receipt
and the board/handoff lines record the check.

## Consume refresh — 2026-09-14 07:28 UTC

Ran `mesh-dash --once adint`; it still reports the active `step0d-hb-first-cell` study and
unresolved resolver `unblock/adint/3dd6562eb2e7cc86/resolve`. A fresh
`mesh-task queue --dispatch --owner 'adint'` exited 0 with no eligible owned rows. Retried the
resolver's bounded SSH condition at 07:28 UTC; it exited 255 with `Connection timed out`, so the
required `mesh-imac-cam --test` read remains unattempted. No camera service, study rows, frame, or
mesh substrate were changed.

The resolver remains blocked until
`timeout 8s ssh -o BatchMode=yes -o ConnectTimeout=5 ilya@192.168.8.214 true` exits 0; only then
run `mesh-imac-cam --test` and require a real read before resuming.

## Consume refresh — 2026-09-14 17:08 UTC

Ran `mesh-dash --once adint`; the live stream still shows `step0d-hb-first-cell`, its RU frame
marker (14 admits / 208 domains; 239 / 4480 readable pairs), and the open
`unblock/adint/3dd6562eb2e7cc86/resolve` obligation. The fresh read-only `tools/adint-status`
derivation at 17:08 UTC reports 14 RU admits / 209 touched domains (3 duplicate rows collapsed)
and 239 / 4480 readable pairs (5.3%). The dashboard and derivation use different touched-domain
denominators; neither was edited.

Retried the resolver's bounded SSH condition at 17:08 UTC; it exited 255 with `Connection timed
out`. `mesh-task queue --dispatch --owner 'adint'` exited 0 with no eligible rows, so no check or
take was applicable. The required `mesh-imac-cam --test` read remains gated on SSH exit 0. Posted
one `[adint] [idle]` line recording the queue, study frame, and blocked retry; sent the same state
to the operator in Russian over `mesh-voice-tx` with duplicated Telegram text. Wake expectation
covers dashboard refresh/tick lines and the new idle-line age churn for 300 seconds.
