# Publication-safety audit — 2026-09-16

Status: **unsafe for publication in the current working-tree state**.

## Scope and reproducible checks

This was a read-only audit of the self-adint working tree. The live repository reported 7 modified
tracked paths and 80 untracked paths. `git diff --check` exited 0. No `data/` files were included
in the publication set; that directory remains gitignored.

Checks run:

```text
git status --short                         -> 7 modified tracked paths; 80 untracked paths
git diff --check                           -> exit 0
git diff --stat                            -> 52 insertions, 33 deletions across 7 tracked paths
added-line private-identifier scan         -> 0 IFA/GAID/IDFA matches
added-line IPv4 scan                       -> 6 matches
untracked-file IPv4 scan                   -> 11 files with matches
```

The IPv4 scan is a publication-boundary check, not a claim that every address is sensitive. The
project doctrine says committed files must not carry node identifiers, so a real egress address is
unsafe until replaced by a placeholder or kept outside the published tree.

## Disposition

Keep as method/code candidates, pending a second review: `tools/adint-step1-run`,
`tools/test-step1-egress-offline.py`, `docs/INDEX.md`, and task-plan/receipt files whose scan has no
address match.

Quarantine or redact before any commit/publication: the modified `README.md`,
`docs/step1-run-2026-08-30.md`, `public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl`,
`ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv`, and `report.html`; their changed lines contain
runtime egress addresses. Do not hand-edit the measured corpus as a cosmetic fix: regenerate it
with publication-safe placeholders or retain the measurement outside the public tree.

Review separately before staging: the 11 untracked receipts containing IPv4-shaped values are
`docs/task-receipts/live-state-20260914.md`, `docs/task-receipts/live-state-20260914T0350Z.md`,
`docs/task-receipts/live-state-20260914T0800Z.md`,
`docs/task-receipts/ru-vantage-wifi-off-20260916.md`,
`docs/task-receipts/retry-camera-ssh-block-20260916T212317Z.md`,
`docs/task-receipts/step0d-first-cell-post-fix-paired-run-20260916.md`,
`docs/task-receipts/step0d-namespace-retry-check-20260916T211452Z.md`,
`docs/task-receipts/unblock-adint-3fc14554d7aaa695-resolve-20260916.md`,
`docs/task-receipts/unblock-adint-a9534782d8636395-resolve-20260916T212836Z.md`,
`docs/task-receipts/unblock-adint-aec9a98f69e425a5-resolve-20260916.md`, and
`docs/task-receipts/unblock-adint-1d9907f10902c94f-resolve-20260912.md`.

The untracked 8-byte lock file is not a publication artifact and should remain untracked; remove
it only through the project’s lock-owner recovery, not as part of this audit.

## Boundary result

No request bodies, advertising identifiers, or observation rows were copied into this receipt.
No router, phone, namespace, capture, or external service was changed. The next safe action before
publication is to resolve the quarantined runtime-address paths, then rerun this exact audit.

Delegation: none. This was a tightly coupled, read-only single-owner audit; I personally inspected
the live status, diff summary, marker scans, and untracked-file scan. No subagent report was used as
evidence.
