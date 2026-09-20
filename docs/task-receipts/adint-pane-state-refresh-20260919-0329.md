# Step 0d pane-state refresh — 2026-09-19 03:32 UTC

Task: `adint-goal-20260919-pane-state-refresh-0329/pane-state-refresh-0329`

## Fresh observations

- `python3 tools/adint-status --include-private`: exit 0.
- Fresh report: canonical `ru-mobile` frame has 209 touched domains, 14 admits, and 239/4480 readable pairs (5.3%); attempted-load coverage is 901/8960 (10.1%).
- `mesh-dash --once adint`: exit 0; `pane:adint` rendered `step0d-hb-first-cell — ru-mobile frame 14 admits/208 domains; 239/4480`, with the task shown active.
- The 208-versus-209 domain wording is presentation drift already visible in the generated report versus the pane marker; the admitted count and readable-pair count agree.
- `ip netns list`: exit 0, with no `ru-mobile` namespace; `test -e /run/netns/ru-mobile`: exit 1. The next collection gate is therefore `UNKNOWN`, not a negative result.

## Input hashes

```text
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3  ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92  public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
b3ff1f29da4e822fc0c448d2f4e86e18b42c16a15489bf4953010f127ed7493f  report.html
```

`python3 tools/adint-status --test` ran the status self-tests but exited 1 because the existing
`docs/status-next-action-runbook-20260918.md` is not listed in `docs/INDEX.md`; the substantive
status assertions shown in the same run passed. No file was changed by the read-only checks.

## Delegation record

Delegated `adint-step0d-independent-read` for an independent read-only inspection. The worker
produced no durable artifact before stopping; no worker report was treated as evidence. I personally
inspected the live `tools/adint-status` output, the three hashes above, the namespace checks, and
the `pane:adint` render.

## Retry edge

When the approved `/run/netns/ru-mobile` exists, verify its RU-mobile egress and rerun the
canonical paired measurement with both-arm vantage verification. No capture, router write,
external contact, onboarding, payment, or GAID reset was performed in this turn.
