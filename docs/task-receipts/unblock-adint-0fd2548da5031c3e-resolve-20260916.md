# adint resolver receipt — 2026-09-16

Task: `unblock/adint/0fd2548da5031c3e/resolve`
Owner: `adint`

## Gate recovery

With explicit operator authorization, stopped the two unrelated resident models reported by the
previous check:

```text
ollama stop qwen3-vl:4b-instruct
ollama stop all-minilm:latest
```

The subsequent `ollama ps` was empty and `pgrep -af 'llama-server'` found no live model server.
The exclusivity gate therefore passed. No mesh credentials or network routing were changed.

## Exact fresh run and replay

From `/home/mesh-home/src/hyperhauntology_for_kids`, ran the exact retry edge:

```bash
python3 -m cryptohaunt run --model tiny-fleet-v1:latest --provider ollama \
  --rule zy --turns 8 --reps 1 --temperature 0.7 --seed 7 \
  --out runs/tiny-fleet-v1_zy_h2-establishment-2.jsonl -v
python3 -m cryptohaunt replay runs/tiny-fleet-v1_zy_h2-establishment-2.jsonl
```

Measured run output: `not-established`, broke at turn 1, `obeyed 0/8 spoken`, `calls 8/8`.
Replay independently read the tape and returned `not-established`; all four families were
`NOT-ESTABLISHED` with `0/0` treatment/control/noise observations because the rule was never
applied. Both commands exited 0.

Measured artifact:

`/home/mesh-home/src/hyperhauntology_for_kids/runs/tiny-fleet-v1_zy_h2-establishment-2.jsonl`

The resolver's external gate is resolved. The experiment's substantive verdict is preserved as
`NOT-ESTABLISHED`, not interpreted as a behavioral result.
