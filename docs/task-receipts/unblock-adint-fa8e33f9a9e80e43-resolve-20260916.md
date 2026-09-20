# Unblock receipt: Ollama exclusivity remains unavailable

Task: `unblock/adint/fa8e33f9a9e80e43/resolve`
Parent: `unblock/haunt/3dfba144a4fbc247/resolve`
Observed: `2026-09-16T12:10Z`

## Live evidence

`ollama ps` exited 0 but reported the unrelated resident model
`qwen3-vl:4b-instruct` (`ee4b975b58c1`, 4.2 GB, 100% GPU, context 8192, about 2 minutes
remaining). Provider exclusivity is therefore false. I personally inspected the live
output and `/home/mesh-home/src/hyperhauntology_for_kids/protocol/alternate-induction-registration-20260916.json`.

The protocol requires a live empty Ollama table and explicitly forbids evicting unrelated
models. The frozen registration and tape replay were not run, and no establishment verdict
is claimed.

Delegation: a read-only Codex worker was launched for an independent diagnosis. It produced
no artifact before this receipt; its report was not used as evidence. The live provider
output and protocol were checked directly in this mind.

## Typed block and retry

Block type: capability. Need: Ollama `ps` empty of unrelated residents. Retry: after a fresh
`ollama ps` shows no unrelated resident model, rerun the exact protocol establishment command,
write its declared tape, replay it offline, and stop on `NOT-ESTABLISHED`.

No model, task ownership, or substrate state was changed.
