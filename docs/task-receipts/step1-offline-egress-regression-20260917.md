# Step 1 offline egress regression — 2026-09-17

## Result

PASS. The existing offline regression completed with exit code 0:

```text
..
----------------------------------------------------------------------
Ran 2 tests in 0.001s

OK
```

The test patches `subprocess.run`, invokes `egress_record(probe_external=False)`, and asserts
that no external probe runs. It also asserts the explicit offline result is `probe_status=unknown`
with reason `disabled by explicit offline mode`, no egress IP fields, and `legs_agree=None`.
The live default remains `probe_external=True`.

## Reproduction

```text
cd /home/mesh-home/self-adint
python3 tools/test-step1-egress-offline.py
```

Exit code: `0`

## Hashes

```text
fe008aab760eb03ccac7578e2b84665b1b50b0d7f855b17a0c7a8d039435b658  /tmp/adint-step1-offline-20260917.out
c4405ff6f0b0fcaf7b0f49b4144583408a10fe776f34a810400df30e2d6fa58c  tools/test-step1-egress-offline.py
6f36217846837152df25d0b6f50515fc2b91e95225c7b6dafa6401fc927004e5  tools/adint-step1-run
```

No external service, corpus, router, phone, namespace, staging area, or publication path was
accessed or changed by this regression.
