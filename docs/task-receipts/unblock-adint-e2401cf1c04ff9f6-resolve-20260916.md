# Unblock receipt — `unblock/adint/e2401cf1c04ff9f6/resolve`

Recorded: 2026-09-16T09:17:30Z

## Result

The dependent wake blocker remains a genuine external capability block. A fresh
credential check returned `Error: Not logged in` with rc=1. No token was read,
copied, printed, or persisted. The safe consumer dry-run was still run and
returned rc=0; it refused the internal board-log-trained fold and listed four
public upload targets, but uploaded nothing.

```text
/home/mesh-home/.venv-ai/bin/hf auth whoami
Error: Not logged in
rc=1

/home/mesh-home/.venv-ai/bin/python /home/mesh-home/finnegans-fake/wake/push_release.py --owner genaforvena --dry-run
REFUSED finnegans-fake-folds-lora share=local — trained on the internal board log
upload genaforvena/finnegans-fake-char257 PUBLIC
upload genaforvena/english-char257 PUBLIC
upload genaforvena/finnegans-fake-bpe4096 PUBLIC
upload genaforvena/finnegans-fake-lora-qwen3.5-0.8b PUBLIC
--dry-run: nothing was uploaded
rc=0
```

## Disposition and exact retry edge

Typed capability block: an operator-authorized Hugging Face credential is still
absent; device authorization was previously cancelled. After an operator-auth
event, run:

```text
/home/mesh-home/.venv-ai/bin/hf auth login
/home/mesh-home/.venv-ai/bin/hf auth whoami
/home/mesh-home/.venv-ai/bin/python /home/mesh-home/finnegans-fake/wake/push_release.py --owner genaforvena --dry-run
```

Only after `whoami` succeeds may the real publication step be considered. It
must continue to refuse the internal board-log-trained fold and must not expose
the credential.
