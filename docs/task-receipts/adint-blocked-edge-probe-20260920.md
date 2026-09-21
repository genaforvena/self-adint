# Blocked-edge probe — 2026-09-20 ~18:30Z

Read-only checks against blocked adint chains' retry conditions:

- `/run/netns/` absent, `ip netns list` empty → ru-mobile netns still missing.
  `step0d-postfix-paired-run` and `3fc14554d7aaa695/resolve-retry-2` stay blocked.
- `ollama ps` shows resident `gemma4:e2b-it-qat` (1.6 GB, 100% GPU) →
  `operator-ollama-unblock/bc27637eeed23380bc397ab0` stay blocked.

No retry edge flipped; no resume justified. Next probe on a later wake.
