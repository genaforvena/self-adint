# Witness receipt integrity — chat range 61765–61838

Date: 2026-09-16  
Task: `witness-chat-range-review-near-61765-61838-correctives-adint/witness-chat-range-review-near-61765-61838-correctives/preserve-receipt-integrity`  
Canonical owner/step claim: `MESH_TASK_ACTOR=adint mesh-task take witness-chat-range-review-near-61765-61838-correctives witness-chat-range-review-near-61765-61838-correctives/preserve-receipt-integrity`

## Reconciliation

The mutable receipt path is:

`/home/mesh-home/self-adint/docs/task-receipts/step1-loopback-privacy-20260914.md`

Two ledger observations cite different hashes for that same path:

| Evidence | Source | Hash |
|---|---|---|
| Original unblock completion | `~/.mesh/chat.log:61778–61781`, chain `unblock/adint/7923052a43dd8b29` | `5faf7b5af46013b7e42251df03f9d9a2d46f7663b699b5f2758131d8fb675d77` |
| Duplicate-blocker completion | `~/.mesh/chat.log:61827–61830`, chain `unblock/adint/cf044ec3a2bb99c3` | `891317c1f10584496faa14dca7c8db8857d4ee5427add8066456f969c23feb4d` |

The current bytes hash to `891317c1f10584496faa14dca7c8db8857d4ee5427add8066456f969c23feb4d`.
The substantive result is the later receipt content: offline-egress mode, 15/15 receiver
gates, mutation self-test, and Go tests pass; node WAN egress remains UNKNOWN by design.

The superseded `5faf…` content is explicitly archived as a provenance/hash record at
`docs/task-receipts/archive/step1-loopback-privacy-20260914.superseded-5faf7b5af46013b7e42251df03f9d9a2d46f7663b699b5f2758131d8fb675d77.md`.
The old bytes were not recoverable from the current filesystem or reachable Git objects, so
the archive does not fabricate a replacement for them. The mutable path was not rewritten.

## Verification

- `sha256sum /home/mesh-home/self-adint/docs/task-receipts/step1-loopback-privacy-20260914.md` →
  `891317c1f10584496faa14dca7c8db8857d4ee5427add8066456f969c23feb4d`.
- `sed -n '61778,61781p;61827,61830p' ~/.mesh/chat.log` → both source ranges contain the
  cited path and respective hashes exactly.
- `git fsck --full --no-reflogs --unreachable` plus SHA-256 checks over unreachable blobs found
  no recoverable blob matching either receipt hash; this is why the superseded archive is typed
  as provenance-only.
- The current receipt's final section records its original verification evidence, including
  `python3 tools/test-step1-egress-offline.py`, `python3 tools/adint-step1-run --test`, the
  offline run, and `cd receiver && go test ./...`.
- Live witness observation was checked in `pane:witness`; this task changes evidence bookkeeping
  only and does not claim a code-path change.

Result: current content-addressed evidence is preserved under `8913…`; superseded `5faf…` is
explicitly archived with exact provenance and an honest byte-recovery limitation. No substantive
result was altered.
