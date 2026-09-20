# Independent adint pilot review — 2026-09-18

Read-only review of the on-disk live report and status/receipt inputs. No capture,
contact, publication, router write, or irreversible action was performed.

## Current measured counts

- Canonical `ru-mobile` frame: **209 domains**, **14 admits**; **3 duplicate rows
  collapsed**.
- Current paired readability: **239 / 4,480 (5.3%)**.
- Attempted-load coverage: **901 / 8,960 (10.1%)**.
- The checked report page describes **208 domains walked** (the report’s walk view),
  with 14 admitted and 194 rejected/blind outcomes; this is not a contradiction of
  the 209-domain canonical frame because the receipts explicitly distinguish the
  canonical frame from the walked report view.
- HB first-cell report: **478 simultaneous pairs**; candidate requests **224
  nl-direct** and **279 ru-mobile**.

## Pilot freshness verdict

The **239/4,480 pilot is current as the latest measured on-disk result**, but it is
also **stale for completion/current-state claims**: the latest reconciliation is
2026-09-18, while the underlying report is generated 2026-09-10, and the result
remains only 5.3% of the target. The receipts say the approved
`/run/netns/ru-mobile` prerequisite is absent, so the full 4,480-pair acceptance
has not been rerun or completed.

## One justified next measurement

When the approved `ru-mobile` namespace is available, run the canonical paired
measurement with both-arm vantage verification and reconcile the resulting count
against **4,480 pairs**. This is justified because it directly tests the stated
blocker and updates the incomplete 239/4,480 coverage without inferring a market
result from the pilot.

## Verifiable inputs

- `docs/task-receipts/step0d-status-reconciliation-20260918.md`
- `docs/task-receipts/hb-first-cell-live-reconciliation-20260917.md`
- `docs/task-receipts/live-study-status-reconciliation-20260918.md`
- `report.html` (SHA-256: `b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f`)
