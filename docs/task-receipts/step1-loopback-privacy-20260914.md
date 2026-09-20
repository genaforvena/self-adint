# Step 1 loopback privacy rerun — 2026-09-14

Task: `adint-step1-privacy-rerun-20260914/verify-loopback-privacy-gates`

## Result

The fresh synthetic receiver run and its fail-under-mutation self-test passed their receiver
gates. **The task acceptance condition did not pass:** the runner's `egress_record()` calls
`curl https://ifconfig.co/json` twice per run, including from each self-test control/mutant
cell. This contradicted the task's explicit no-external-network boundary. I found this only
after running it. `unshare -n true` failed with `Operation not permitted`, so this node could
not isolate another run from external networking. Do not mark the task done or rerun the
harness until an isolated runner or an explicitly revised boundary is available.

The standalone run made two successful egress queries. The self-test ran one control and nine
mutant cells, each with two successful egress queries: **22 requests to `ifconfig.co` total**.
The egress IP is omitted here; the standalone summary retains the raw measurement under `/tmp`.
The queries carried no corpus or device data, but they were outside this task's declared scope.

## Verification

- `python3 tools/adint-step1-run --seed 20260914 --out-dir /tmp/adint-step1-privacy-20260914`: exit 0. Receiver bound to `127.0.0.1`; all 15 displayed gates passed, including matched-row completeness, no foreign row/IFA or coordinates on disk, panic path exercised, foreign pass/deny behavior, missing IFA rejection, floor, spend cap, and win artifacts.
- `python3 tools/adint-step1-run --test`: exit 0. The control was green and all nine named mutants made their target gate red.
- `unshare -n true`: failed with `Operation not permitted`.
- Task audit: no other active, queued, or blocked row matched Step 1 receiver/privacy verification. Existing open project work is the separate Step 0D frame study; the exact-owner camera resolver remains blocked on its documented external capability condition.

## Artifact hashes

All generated test data is synthetic and remains under `/tmp`; no device capture was read or written.

```text
1056628cab8e49d5deefa69f5db9fb3501b1c4ba44526b4d8ca2664ea6b97d6d  summary.json
f2f8638a81ad7a0073ca8b17661e186465689141f81404e3f2cef6aeb952ca07  results.jsonl
0c0b184e1864af6e11d6c60a76e1b7564dcaa362160bc6ad41303c715e9ab630  data/receiver.jsonl
f08073acff78215bd6bbaa31b5eb58aa9f3e7ad7e2e2569dbb134cef0ab8f244  data/win-counter.txt
c60de29edf253503c93d8f1afbe4cc8e5c1ecc9f0a8cb7f08b61deb8208c2c74  receiver.stdout.log
f2dc28a0f5afd78a6855b02015d2c1d1509d1528adc1093eeb4c8211f168072c  corpus.jsonl
```

## Follow-up — offline probe mode, 2026-09-14

Implemented an explicit `--offline-egress` path in `tools/adint-step1-run`. It skips both
`ifconfig.co` subprocess calls and labels node-egress readings `unknown`; the live probe
remains the default. The harness's `--test` now uses offline mode for its synthetic control
and mutant cells. The scope is explicit: receiver requests are verified as loopback-only;
the node's public egress is UNKNOWN in offline mode, never asserted to be loopback.

Verification:

- `python3 tools/test-step1-egress-offline.py`: PASS (2 tests). The test failed first because
  `egress_record` had no offline parameter, then passed after the implementation.
- `python3 tools/adint-step1-run --test`: PASS; the control remained green and all nine
  mutants turned their target gates red. Ran with a `curl` guard first in PATH; its marker
  was absent after the test, proving the self-test made no curl invocation.
- `python3 tools/adint-step1-run --seed 20260914 --out-dir
  /tmp/adint-step1-privacy-offline-20260914 --offline-egress`: exit 0; 15/15 receiver
  gates passed. A `curl` guard first in PATH remained untouched. The summary records the
  receiver path as loopback-only and node egress as `unknown`; it contains no egress IP.
- `cd receiver && go test ./...`: PASS.

The Step 1 scope is clarified here: offline verification establishes that the harness made
no external probe and that receiver requests stayed on loopback; it does not measure the
node's WAN route. Thus the prerequisite is satisfied without changing the live-probe
default. The unblock resolver chain is complete; the owner resumed the original rerun row
after its resume check returned 0. This receipt is the artifact for settling that rerun.

Follow-up artifact hashes:

```text
6f36217846837152df25d0b6f50515fc2b91e95225c7b6dafa6401fc927004e5  tools/adint-step1-run
c4405ff6f0b0fcaf7b0f49b4144583408a10fe776f34a810400df30e2d6fa58c  tools/test-step1-egress-offline.py
b59583e3b339f127253f507510473ad8a87d8385e7eefb6e590d2121a0679d57  docs/step1-run-2026-08-30.md
27390337cbfc2f39c11613680427a1b9dfe150b3b96bebe453cb112611e1478a  /tmp/adint-step1-privacy-offline-20260914/summary.json
15e2191f8bd5ae480e6b397f8b9aded4841d44cb8367ba897470d46b99a55318  /tmp/adint-step1-privacy-offline-20260914/results.jsonl
4376b6454ecf5edf3763f34bd58a2e60fe58943fd591349bf8dd480b90bc0696  /tmp/adint-step1-privacy-offline-20260914/data/receiver.jsonl
f08073acff78215bd6bbaa31b5eb58aa9f3e7ad7e2e2569dbb134cef0ab8f244  /tmp/adint-step1-privacy-offline-20260914/data/win-counter.txt
7734b686c960d7c89858e49554b1fb12eef1aeeb8db489b6dd3a4c90ff71c15c  /tmp/adint-step1-privacy-offline-20260914/receiver.stdout.log
f2dc28a0f5afd78a6855b02015d2c1d1509d1528adc1093eeb4c8211f168072c  /tmp/adint-step1-privacy-offline-20260914/corpus.jsonl
```
