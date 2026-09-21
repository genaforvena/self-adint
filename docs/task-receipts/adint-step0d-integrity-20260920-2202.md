# Step 0D integrity check — 2026-09-20T22:04Z

Task: `adint-goal-20260920-step0d-integrity-2202/step0d-integrity-2202`.

## Checks

| command | result |
|---|---|
| `python3 tools/adint-status --test` | exit 0; `selftest: ok` |
| `go test ./...` from `receiver/` | exit 0; `selfadint/receiver (cached)` |
| `python3 tools/adint-status` | exit 0; generated 2026-09-20 22:04 UTC |

The live disk-backed `ru-mobile` frame reports 209 domains touched, 14 admits, and
239/4480 readable pairs (5.3%); attempted coverage is 901/8960 (10.1%). The 14 admitted
domains match the prior first-cell snapshot. The prior snapshot's 209-domain count and
239/4480 coverage therefore remain consistent; the pane's 208-domain display is a
presentation discrepancy from three collapsed duplicate rows, not a disk-data change.

The receiver privacy suite passed its foreign-IFA persistence gate: `TestForeignIFANeverWritten`
and the surrounding receiver tests are green. No capture, network mutation, or publication
was performed. No new demand evidence was produced. `ru-mobile` remains the frozen frame;
the missing fast RU vantage remains the next external edge.

Observation: `pane:adint` rendered the live `mesh-dash --once adint` state at 22:01:44Z;
this receipt is the disk-backed recheck at 22:04Z.
