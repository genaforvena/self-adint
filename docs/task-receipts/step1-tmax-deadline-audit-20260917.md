# Step 1 `tmax` deadline audit — 2026-09-17

## Verdict

`GAP`: the shipped Step 1 receiver does not read each request's `tmax` and does not enforce
`tmax - RTT`. The harness carries `tmax` and records elapsed service time, but that is
measurement of the local run, not enforcement of the exchange deadline.

## Evidence

Commands run from `/home/mesh-home/self-adint`:

```text
rg -n 'tmax|RTT|deadline|latency|elapsed|timeout|time.?budget' receiver tools/adint-step1-run docs/step1-run-2026-08-30.md PLAN.md
```

Relevant source evidence:

- `tools/adint-step1-run:97-101`: `tmax` and `at` are carried on every request, while the
  receiver reads neither; the harness reports observed latency against declared `tmax`.
- `tools/adint-step1-run:120`: generated requests receive `tmax` values from 80, 100, 120,
  or 150 ms.
- `tools/adint-step1-run:413-423`: the harness posts the request and stores measured `ms`
  plus the request's `tmax`; no budget is passed to the receiver.
- `tools/adint-step1-run:572-583`: the harness explicitly records
  `tmax_read_by_receiver = false` and `at_read_by_receiver = false`.
- `receiver` source contains no `tmax`, `RTT`, or deadline-read implementation.
- `docs/step1-run-2026-08-30.md:140-145`: the documented run says RTT is zero by construction
  and records `tmax_read_by_receiver: false`.

Runtime verification:

```text
timeout 120s python3 tools/adint-step1-run --test
...
selftest: PASS  (workdir /tmp/adint-step1-selftest-arban5ay)
TEST_STATUS:0
```

The green self-test proves the existing receiver/privacy/spend-cap gates, not `tmax`
deadline enforcement. No external probe, device data, publication, or irreversible action
was used.

## Next action

Implement a separate Step 1 deadline gate only after deciding the receiver contract for
request timing: parse `tmax`, measure request-arrival/response RTT, and assert a bounded
decision budget. This receipt intentionally does not implement that change.
