# Self-ADINT Full Study Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Complete the safe, reproducible self-ADINT study loop through measured artifacts, while leaving seat onboarding, payments, outbound letters, and GAID reset explicitly gated for the operator.

**Architecture:** Keep the existing passive browser lane and device-side PCAPdroid importer as separate instruments. Make each instrument write append-only, provenance-bearing artifacts; add an independent echo control to prevent a single third-party service from suppressing a collection tick; add a study coordinator that reports which planned phase is measured, blocked, or gated without turning absence into a negative.

**Tech Stack:** Python 3 standard library tools, Go receiver, JSONL/TSV artifacts, existing browser virtualenv, existing Android/Termux and PCAPdroid workflows.

## Global Constraints

- Only the operator's device may reach persistent capture storage; foreign `device.ifa` is filtered before buffering.
- Request bodies are never logged by the passive lane; metrics are counters only and never IFA-labelled.
- A foreign device may be parsed transiently only where required to answer an auction, and is never retained or aggregated.
- The VM stays outside the mesh and no mesh credential is added to the project.
- No seat onboarding, agreement, payment, outbound data-subject letter, or GAID reset is performed by this implementation.
- `data/` remains gitignored; published files contain method and synthetic fixtures, not device identifiers.
- Every claimed capability has a test artifact and every unavailable measurement remains an explicit `UNKNOWN`, `BLOCKED`, or `NO-BASIS` state.

---

### Task 1: Remove the passive lane's single echo failure point

**Files:**
- Modify: `tools/adint-collect-passive`
- Modify: `docs/step4-passive-lane-2026-08-30.md`
- Test: `tools/adint-collect-passive --test`

**Interfaces:**
- `egress_legs(runner=None, hosts=None)` returns both process legs, the providers attempted per leg, and `legs_agree` as `True`, `False`, or `None`.
- The default providers are the already-established Yandex IP endpoint and `ifconfig.co`; a provider failure is recorded, never silently replaced with a fabricated IP.

- [x] Write a failing self-test that injects one dead provider and two agreeing providers, and asserts the result records the dead provider plus the successful provider rather than returning an error-only leg.
- [x] Run `tools/adint-collect-passive --test`; the assertion first failed on the missing injectable provider interface.
- [x] Implement provider rotation and explicit per-provider evidence using only the standard library and the existing curl subprocess boundary.
- [x] Run the focused self-test and then the full local tool suite.
- [x] Update the step-4 document with the new artifact fields and the interpretation of an all-provider failure.
- [ ] Commit the focused change.

### Task 2: Make device-side capture handoff executable and auditable

**Files:**
- Create: `tools/adint-device-capture`
- Modify: `tools/adint-import-pcapdroid`
- Create: `docs/device-capture-runbook-2026-09-07.md`
- Test: `tools/adint-device-capture --test` and `tools/adint-import-pcapdroid --test`

**Interfaces:**
- `adint-device-capture --host user@phone --remote-csv PATH --out PATH --device LABEL --pull` performs only an explicit, read-only CSV pull and writes a manifest beside the imported artifact.
- `--dry-run` prints the exact SSH/SFTP command without contacting the device.
- A missing host, missing remote CSV, failed copy, malformed CSV, or device-label mismatch is a named failure; no empty artifact is emitted.

- [x] Write failing tests for command construction and manifest fields.
- [x] Run the focused test and observe RED on the missing manifest implementation.
- [x] Implement the explicit pull/import/manifest flow without embedding a phone address, username, key, or credential.
- [x] Document the operator action needed to start/stop/export PCAPdroid; the tool must not silently drive Android UI or VPN state.
- [x] Run focused and full tests.
- [ ] Add an end-to-end synthetic CSV fixture test and commit.

### Task 3: Add an executable phase/status coordinator

**Files:**
- Create: `tools/adint-study`
- Modify: `docs/INDEX.md`
- Create: `docs/study-runbook-2026-09-07.md`
- Test: `tools/adint-study --test`

**Interfaces:**
- `adint-study --out data/study-status.json` emits one status row for Steps 0, 0D, 1, 2, 3, 4-device, 4-web, 5, and 6.
- Safe phases derive status from existing artifacts; irreversible phases are always `GATED_OPERATOR_GO` until an explicit marker supplied for that exact step exists.
- Step 4 reports device and web instruments separately; web data cannot masquerade as phone data.
- Step 5 remains `NOT_ELIGIBLE` until Step 4 establishes payload-empty/hashed evidence; Step 6 remains gated until a baseline artifact exists and the operator authorises the reset.

- [x] Write failing tests for missing artifacts, explicit operator gates, and the distinction between `UNKNOWN`, `NO-BASIS`, and `NOT_ELIGIBLE`.
- [x] Run the focused test and observe RED on the missing phase resolver.
- [x] Implement the coordinator with deterministic JSON output and no network calls.
- [ ] Add the runbook mapping each status to the exact next local command or operator question.
- [x] Run the full suite and generate the first status artifact without inventing private measurements.
- [ ] Commit.

### Task 4: Close the study with measured artifacts and review

**Files:**
- Modify: `docs/step4-passive-lane-2026-08-30.md`
- Modify: `README.md` through `tools/adint-status --write` only
- Create: `docs/full-study-review-2026-09-07.md`

- [ ] Run all relevant self-tests, `go test ./...`, and the study coordinator.
- [ ] Inspect the private ledger for actual collection outcomes; never infer a week from a cron line.
- [ ] If the phone is reachable and an operator-exported CSV exists, run the device tool and produce `bundle → demand-stack entry point` artifacts.
- [ ] If the phone is unreachable or no export exists, record that as a specific blocked artifact and ask the operator for the one missing action.
- [ ] Verify no private `data/` file is staged and no mesh-genome file was modified.
- [ ] Post one `[done]` or `[yield]` board line with exact artifacts and unresolved obligations; write the handoff.

## Self-review

- The plan covers the known open gaps from the restored handoff: device-side Step 4 and echo redundancy.
- It does not claim that public-web browser captures are phone captures.
- It does not execute the irreversible Steps 2, 3, 5, or 6.
- It preserves the existing parser-before-buffer and no-body-logging guarantees.
