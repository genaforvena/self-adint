# Device entry-point preflight — 2026-09-18

Status: **MEASURED from an existing operator-device export; exact-owner task registered and claimed after bounded replay-cache retries.**

## Evidence inspected

- Delegated review artifact: `/tmp/adint-goal-review.txt` (personally inspected). It selected the device-side `bundle-exchange.tsv` as the next bounded demand artifact, but incorrectly treated the existing output as absent.
- Existing outputs are present but dated 2026-08-15 14:40:00 UTC: `data/bundle-exchange.tsv` (7,195 bytes) and `data/unresolved-hosts.tsv` (2,777 bytes). They are gitignored device-derived files and were not overwritten.
- Existing normalized operator-device input: `data/out/obs-1739.jsonl`, source `pcapdroid-csv`, target device `redmi10-daily`.

## Checks

```text
python3 tools/adint-aggregate --test                 PASS (0 checks failed)
python3 tools/adint-study --test                     PASS
python3 tools/adint-study                            PASS; step4-device=BLOCKED_OPERATOR_EXPORT
python3 tools/adint-status --include-private         PASS
```

The parser was re-run into a fresh temporary directory, without changing `data/`:

```text
python3 tools/adint-aggregate --device redmi10-daily \
  --in data/out/obs-1739.jsonl --out-dir /tmp/adint-entry-preflight.yWNWcx
coverage: 643/1943 observations named (33.1%) · kept=317 foreign-dropped=0 malformed=0
wrote .../bundle-exchange.tsv (103 rows) · .../unresolved-hosts.tsv (149 hosts)
```

Input/reference/output hashes from that run:

```text
9e9ae4d247dd173ff9d275cf4f7c8952d0893ac00433db4f86232b87a64ff58d  data/out/obs-1739.jsonl
331712fffdac2cdbd41e6493208518b8429e0532af42bf106f061ec365e95e39  ref/exchange-domains.tsv
e5d153c08b96b5e1a59ab15d5a675fc2d24ffe850c4bdd4e7842f0cd670d2e3f  bundle-exchange.tsv (temporary output)
5af3daf621e122da31f66057c7358db009fe2acf2017bddcba5656ea6dd62063  unresolved-hosts.tsv (temporary output)
```

## Disposition and retry edge

This is a safe local measurement, not proof of an auction, winning bidder, or sale volume. The current export is historical; a fresh operator-device export is still required before treating the exchange shortlist as current. No capture, contact, publication, router write, seat onboarding, payment, or identifier reset occurred.

The first two bounded 20-second registration attempts timed out (`rc=124`); live `lslocks` showed 20 concurrent `WRITE*` waiters on `~/.mesh/.task-replay-cache.json.lock`, and a later 15-second attempt showed 22. A subsequent retry completed: task `adint-goal-20260918-device-entry-point-preflight/device-entry-point-preflight` was created and claimed by `adint`. The next retry edge is a fresh operator export; preserve `UNKNOWN` until it is supplied.
