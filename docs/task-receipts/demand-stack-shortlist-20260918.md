# Demand-stack shortlist — 2026-09-18

## Verdict

The canonical `ru-mobile` Stage B corpus contains **14 admitted domains**. The
reproducible shortlist is a wrapper-level demand-stack signal, not an exchange
or bidder attribution:

| observed label | admitted domains | interpretation |
|---|---:|---|
| `YaHeaderBiddingSettings(...)` | 14 | Yandex header-bidding settings were observed |
| `pbjs:pbjs` | 1 | Prebid.js global was observed (`ria.ru`) |
| organization/bidder name | UNKNOWN | schema 3 publishes no bidder list or organization label |

The bidder count embedded in `YaHeaderBiddingSettings(bidders=N)` is retained
as an observation, but it is not a count of named organizations. The corpus
therefore does **not** justify choosing an exchange or claiming sale volume.

## Admitted domains

`aif.ru`, `aviasales.ru`, `gismeteo.ru`, `hh.ru`, `interfax.ru`, `iz.ru`,
`kp.ru`, `magnit.ru`, `matchtv.ru`, `mk.ru`, `pikabu.ru`, `rbc.ru`, `ria.ru`,
`sports.ru`.

## Reproduction

From the repository root:

```sh
sed '/^#/d' public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl \
  | jq -s '{rows:length, admitted: map(select(.stageb_verdict=="admit")), admitted_count: (map(select(.stageb_verdict=="admit"))|length)}'
```

The source file is comment-prefixed metadata plus JSONL. The observed run had
212 JSON rows and 14 rows with `stageb_verdict=admit`; `ref/ru-frame-stageB-
ru-mobile-2026-08-19.tsv` contains the same 14 admitted domains as one row per
domain.

Input SHA-256:

```text
ref/CANONICAL-FRAME
05a24ad31aec60db2adea10cfe59e6500d73ec20f097245cee89bfb75383dccf

ref/ru-frame-stageB-ru-mobile-2026-08-19.tsv
b1edf784d8db3918a9f0c0abe7e1a717d76dd3e9c1a8c981a5441f0d8067cee3

public-data/frame-stageb-ru-mobile-2026-08-19-schema3.jsonl
8909efb780fe9de5a8c2957475245eed2be160655afe1f0bbf45e1789bf13c92

ref/org-aliases.tsv
216f5425ac8b85dc0a843232b9cb8c634059daa7ce1d443c21af743562ebca41
```

## Next exchange-selection gate

Before any seat, agreement, payment, or external contact, obtain a fresh
operator-device `bundle → demand-stack entry point` artifact from the existing
device-side capture path (`tools/adint-aggregate`). Admit an exchange candidate
only when the artifact has a hostname resolved to a known organization with
`role=exchange`, `external_seat=yes`, and a positive observation count. Keep
unresolved hosts and non-exchange delivery/analytics rows as `UNKNOWN` or out
of the exchange shortlist. The current Stage B corpus alone fails this gate by
design: it is public-web wrapper evidence, not the operator-device wire path.

No capture, external probe, publication, exchange contact, or irreversible
action was performed for this receipt.
