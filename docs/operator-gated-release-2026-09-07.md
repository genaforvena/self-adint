# Operator-gated release bundle — 2026-09-07

The reproducible handoff is generated at `data/operator-release-2026-09-07/` by
`tools/adint-operator-release`. It keeps the unfinished operator branches in one append-only
decision ledger and gives each branch an evidence block plus an explicit witness disposition.

The bundle is deliberately redacted: it does not copy the raw device capture, private passive
ledger, full advertising ID, request body, or any foreign-device observation. The passive baseline
is bound by `baseline-lock.json`; its SHA-256 is the comparison anchor, and no reset was performed.

## Measured disposition

| branch | status | witness disposition |
|---|---|---|
| device capture | `BLOCKED_OPERATOR_EXPORT` | unfinished; operator export absent |
| seat onboarding | `GATED_OPERATOR_GO` | unfinished; explicit operator decision required |
| active oracle | `NOT_ELIGIBLE` | unfinished; measured device/seat payload gate absent |
| GAID reset | `NO-BASIS` | unfinished; device baseline and explicit reset go absent |

No onboarding, agreement, payment, active bid, outbound letter, or GAID reset was performed.

## Verification record

The release tool's self-test checks clean two-directory reproduction, baseline hash equality, and a
witness disposition on every unfinished ledger row. The generated bundle is then checked with
`(cd data/operator-release-2026-09-07 && sha256sum -c SHA256SUMS)`.

Exact next action: obtain one complete operator-exported PCAPdroid CSV plus reachable Termux
SSH host/device label, run `tools/adint-device-capture --pull`, and regenerate the bundle.
