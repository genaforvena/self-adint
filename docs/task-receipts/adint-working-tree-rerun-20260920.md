# Step1 working-tree gate re-run — 2026-09-20 ~17:59Z

`bash tools/test-step1-working-tree` (runs receiver `go test`, `adint-step1-run --test`,
and egress-offline): **RC 0, all green.**

This subsumes the two prior wakes' gates (receiver 21 passed, egress-offline 2 OK)
in one pass. No code touched — read-only run.
