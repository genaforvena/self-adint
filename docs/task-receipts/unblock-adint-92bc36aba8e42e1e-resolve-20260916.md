# Recovery receipt: unblock/adint/92bc36aba8e42e1e/resolve

- observed: 2026-09-16T04:42:31Z
- owner: adint
- task: Resolve the haunt blocker only when Ollama is provider-exclusive.
- prerequisite: `ollama ps` must show no unrelated resident model.

## Observation

Commands run from the adint turn:

```text
ollama ps
NAME                    ID              SIZE      PROCESSOR    CONTEXT    UNTIL
qwen3-vl:4b-instruct    ee4b975b58c1    4.2 GB    100% GPU     8192       4 minutes from now
```

The Ollama service and its corresponding `llama-server` are still running. No model was stopped or unloaded; stopping a resident model is outside this mind's safe authority.

## Exact retry edge

When `ollama ps` is empty (or otherwise contains no unrelated resident model), rerun the exact fresh repetition and replay:

```bash
cd /home/mesh-home/src/hyperhauntology_for_kids
python3 -m cryptohaunt run --model tiny-fleet-v1:latest --provider ollama \
  --rule zy --turns 8 --reps 1 --temperature 0.7 --seed 7 \
  --out runs/tiny-fleet-v1_zy_h2-establishment-2.jsonl -v
python3 -m cryptohaunt replay runs/tiny-fleet-v1_zy_h2-establishment-2.jsonl
```

This receipt records a concrete external-event block, not a completion: the fresh run and replay were not started because exclusivity was false.
