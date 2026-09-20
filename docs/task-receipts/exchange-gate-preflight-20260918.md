# Exchange-selection gate preflight — 2026-09-18

## Verdict

**PASS for the repository preflight; UNKNOWN for an actual device demand-stack
candidate.** The reference resolver and aggregate privacy/self-test paths pass,
and the reference table contains 24 rows for `role=exchange` with
`external_seat=yes` across 10 organizations. This proves the static gate is
available; it does not prove that any of those organizations appeared in a
fresh operator-device bundle.

## Commands and exit codes

Run from `/home/mesh-home/self-adint`:

```text
python3 tools/adint-aggregate --test   # exit 0; selftest: 0 checks failed
python3 tools/adint-ref-check --test   # exit 0; selftest: 0 checks failed
awk -F '\t' 'NR>1 && $3=="exchange" && $5=="yes" {n++; org[$2]++} END {print n, length(org)}' ref/exchange-domains.tsv
# 24 rows, 10 organizations
```

The two executables are present and executable. The aggregate self-test
explicitly covers foreign-row exclusion, unknown-host accounting, and source
recording. The reference self-test covers the exchange/seat consistency rules,
resolver reuse, and missing-reference failure.

## Input hashes

```text
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf  ref/CANONICAL-FRAME
331712fffdac2cdbd41e6493208518b8429e0532af42bf106f061ec365e95e39  ref/exchange-domains.tsv
216f5425ac8b85dc0a843232b9cb8c634059daa7ce1d443c21af743562ebca41  ref/org-aliases.tsv
95fdad902d7f3ea9d114f395d23d6449e1436aefa49f21d179e9026ead23a088  tools/adint-aggregate
bc800c87e66dbd33785334eccf82773f6fb73790cee8b28a1753c1b3f1940166  tools/adint-ref-check
```

## Boundary and retry edge

No capture, external probe, exchange contact, publication, or irreversible
action was performed. The next gate remains a fresh operator-device
`bundle → demand-stack entry point` artifact from `tools/adint-aggregate`.
Only a hostname resolved to a known organization with `role=exchange`,
`external_seat=yes`, and a positive observation count may enter the shortlist;
otherwise the result stays UNKNOWN. The required device artifact is not
manufactured by this static preflight.
