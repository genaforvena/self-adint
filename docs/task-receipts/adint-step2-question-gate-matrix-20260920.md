# Step 2 question-to-gate matrix — 2026-09-20

Task: `adint-goal-20260920-step2-question-gate-matrix/step2-question-gate-matrix`.

Read-only local audit. No exchange was contacted, no account was opened, and no external or
irreversible action was taken.

## Reproducibility

Sources read at `2026-09-20T23:40Z` UTC:

| source | sha256 |
|---|---|
| `PLAN.md` | `878283bcbfa46b90b725b9240ada632427250efb44774b2dc3c4721ae1e5648e` |
| `docs/step2-questions-exchange-2026-08-15-ru.md` | `8326e8f5b174a8d53663f1a11a17afffe85c666276477223ec7c46a211a466b0` |

The plan defines six Step 2 gates. The question draft also contains q7 (creative) and q8
(admission/settlement) as supporting onboarding questions; they are recorded below as ancillary,
not silently counted as additional Step 2 gates.

## Matrix

| plan question | decision controlled | minimum answer/evidence | blocks | current boundary |
|---|---|---|---|---|
| 1. Is `user.data[]` exposed to a seat with no spend history? | Whether passive Step 4 can produce the profile fields the study seeks from day one. | Written confirmation of first-day availability, or the spend/eligibility threshold and exact fields unlocked. | Blocks Step 4 collection if access requires spend; therefore blocks the evidence path before onboarding is justified. | **UNKNOWN**: the local draft describes the gate but contains no exchange answer. |
| 2. Is `device.ifa` present, open, and under what redaction rules? | Whether the device can be identified in the live stream by passive observation, or whether the project must use the separately gated active-oracle branch. | Per-inventory and per-region/consent behavior, including hash/salt semantics and a sample or contractual field definition. | Blocks passive self-identification; it does not authorize active targeting. | **UNKNOWN**: no live seat response or bid-request sample from a real exchange. |
| 3. Is a forecast/reach API available, with what floor, rounding, limits, and login gate? | Whether forecast is available as cheap reconnaissance to narrow candidate segments. It cannot replace live bid-request evidence. | API/docs response naming access prerequisites, minimum non-zero audience, rounding, rate limits, and granularity. | Blocks only the forecast reconnaissance branch; it must not block or substitute for the live-stream study. | **UNKNOWN**: the draft explicitly rejects inference from forecast as proof about this device. |
| 4. What bid rate avoids throttling, and where is the seat PoP/timing boundary? | Whether the planned bidder can obtain a stable, interpretable sample without rate throttling or avoidable timeout loss. | Written minimum response rate, throttle behavior, PoP/region, timeout, and any QPS constraints. | Blocks reliable Step 4 sampling and informs Step 3 QPS negotiation. | **UNKNOWN**: no exchange-specific operational answer exists. |
| 5. May bid-request data be retained/used outside bidding purposes? | The legal/contractual go/no-go for passive collection and every downstream step. | Written clause or written clarification covering retention duration, purpose, permitted analysis, and deletion/exception terms. | **Hard gate on Step 3 onward**: direct prohibition stops the project; a time limit rewrites Step 4; silence is not permission and requires clarification. | **UNKNOWN / project-killer gate**: the draft identifies this as decisive, but no answer exists. |
| 6. Does targeting support OR over a segment list, with what maximum and billing semantics? | Whether grouped testing and its 250–300-test cost model are viable, or whether the oracle must be redesigned toward linear probing. | Written semantics (OR vs AND), maximum list length, per-segment charging, and behavior when the list is truncated. | Blocks finalizing Step 5 test design and its cost estimate; it does not authorize any test. | **UNKNOWN**: no platform-specific semantics or cap is established. |

## Ancillary questions in the draft

Q7 (creative hosting, `data:` URI, size, moderation, render deadline) controls creative preparation
and paid-zero risk after a seat path exists. Q8 (individual vs legal entity, pre/post-pay, deposit,
minimum monthly spend) controls admission and budget feasibility. Both remain unanswered and
operator-facing; neither is a reason for this mind to contact an exchange.

## Operator-only handoff

The document is ready to send as a question set, but sending is an explicit operator action under
the project charter. Until written answers exist, q5 remains the hard gate and Step 3 onboarding,
payment, and Step 4/5 execution do not start. No answer may be inferred from the frozen RU-mobile
pilot: its missing namespace/egress is a separate UNKNOWN boundary.

Acceptance evidence: both source hashes above were produced with `sha256sum`; the matrix covers all
six questions named by `PLAN.md` Step 2, preserves UNKNOWN where no answer exists, and records the
operator-only handoff.
