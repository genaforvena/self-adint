# AdFox lays its cards on the table — which cards, and which it keeps

**2026-08-21.** Operator's hypothesis, in his words: *«AdFox раскрывает все карты своих
торгов»*. This tests it directly against the wire, not against an earlier draft.

Tools: `tools/adint-adfox-cards`, `tools/adint-yandex-vocab` (both `--test`, each gate
seen RED under a mutant before being trusted).
Artifacts: `data/adfox-cards-2026-08-21.json`, `data/yandex-vocab-2026-08-21.json`.
Corpus: 20 paired ledgers, 2026-08-19..20, 23,202 request rows, 14 RU publishers,
two simultaneous arms (`ru-mobile` = MTS PDP in Moscow, `us-exit` = Tailscale exit node).

## What is being read, and why "reveals" is the right verb

On a page running AdFox the browser calls

```
https://yandex.ru/ads/adfox/<ownerId>/getBulk/v2?…&bids=<base64>
```

`bids` is base64 of a JSON **array**: the header-bidding results the page already
collected, reported back to Yandex. One element per competing bidder.

This is a **request the page itself makes**. No privileged vantage, no interception —
anyone can read it in their own browser's network tab on their own screen. That is what
makes it a disclosure rather than a leak.

## Coverage

| | |
|---|---|
| getBulk/v2 requests | 1,110, from 16 distinct AdFox owner ids |
| carrying `bids=` | 1,093 |
| decoded | 1,093 (0 failures) |
| decoded to `[]` | 786 |
| decoded to a non-empty array | 307 → **1,259 bidder entries** |

So the auction is legible on **28.1 %** of the requests that report one.

## (a) The roster — the hypothesis HOLDS, completely

Every field, over all 1,259 entries:

| field | present |
|---|---|
| `bidderName` | 100.0 % |
| `response_time` | 100.0 % |
| `error` | 99.1 % |
| `campaign_id` | 98.3 % |
| `placement_id` | 82.8 % |
| `bid` / `cpmAdjustment` / `currency` / `unit` | **0.9 %** |

**29 distinct bidders** are named — `mytarget`, `adriver`, `buzzoola`, `betweendigital`,
`mediasniper`, `roxot`, `gnezdo`, `astralab`, `sape`, `videonow`, `smi2`, `adwile`, `umg`,
`bidvol`, `segmento`, `otm`, `hyper`, `getintent`, `alfasense`, `hhru`, `hybrid`, four
`adfox_*` and four `pb_*` adapters. Per-bidder latency: n=1,259, min 68 ms, p50 1,002 ms,
p90 1,521 ms, max 2,015 ms. Seven distinct failure codes, `3` and `1` covering 89 %.

Who was invited, how slow each one was, and why each dropped out is fully legible from
the page's own outbound request. On this half the operator is right without qualification.

## (b) The price — the hypothesis BREAKS

**11 of 1,259 entries (0.87 %) carry a number.** The other 1,248 carry an error instead.
The cards shown are the **players and their exits, not the stakes**. Reading this as
"AdFox publishes its clearing prices" is wrong on 1,248 of 1,259 entries.

The 11 prices: `roxot` 1–30 RUB (`unit=1`) and `adwile` 1,617–2,097 RUB (`unit=2`), on
three publishers, `cpmAdjustment=1` throughout.

## The part that was worth the paired design: **a price is a function of the vantage**

All 11 priced entries are on the **RU arm**. Zero on the US arm.

That is only interesting if it is not a site-composition artifact, so it was checked. On
the three sites where a price appeared, the US arm was **more** legible, not less:

| site | non-empty auctions ru / us | entries ru / us |
|---|---|---|
| matchtv.ru | 5 / 18 | 40 / 144 |
| aif.ru | 2 / 6 | 24 / 72 |
| kp.ru?section=zenyandex | 24 / 32 | 114 / 152 |

And both bidders were **present on the US arm more often than on the RU arm**:

| bidder | ru entries → priced | us entries → priced |
|---|---|---|
| `roxot` | 31 → 6 (19 %) | 41 → **0** |
| `adwile` | 6 → 5 (83 %) | 8 → **0** |
| pooled | 37 → 11 | 49 → **0** |

One-sided Fisher **p = 3.5 × 10⁻⁵**. The same bidders, on the same sites, in
simultaneous paired loads, asked *more* often from the US vantage, bid **zero** times
there. This is a bid/no-bid decision keyed on where the browser appears to be.

### One roster difference is real; the other is confounded — say which

Four Prebid adapters — `pb_smartadserver`, `pb_adf`, `pb_onetag`, `pb_cpmstar` — appear in
**12 of 12** US auctions on gismeteo.ru and **0 of 14** RU auctions on the same site. Both
arms were legible there, so this one stands.

Six bidders looked RU-only (`umg`, `bidvol`, `hhru`, `otm`, `segmento`, `hyper`) — **this
does not survive checking.** They live on iz.ru, hh.ru and rbc.ru, and on those three sites
the US arm produced **zero** non-empty auctions at all. Absence of the site's auction is
not absence of the bidder. Not a finding; recorded so nobody re-derives it as one.

## The naming, since the odd names landed

220 distinct query-parameter names on yandex.ru ad endpoints in this corpus.

**`puidN` are not names, they are POSITIONS.** 54 distinct `puid` keys, up to `puid63`, and
the same key means different things per publisher. `puid2` is a **weather word** on
gismeteo (`cloud` ×38, `rain` ×21, `sun` ×4) and an opaque 24-, 69- and 2-character token on
hh.ru, kp.ru and sports.ru. Anything downstream reading `puid2` as a field is reading four
different fields under one name. `puid44/45/46` on matchtv carry Russian content categories
in plain text. Most high-numbered slots carry the literal four-character string `null` —
neither absent nor empty, a third state.

**`puid29` carries the vantage.** `RU_` on 96/96 RU-arm requests, `CA_QC` on 72/72 US-arm
requests — Quebec, which is where this node's exit node egresses. The publisher's own
targeting parameter is a ground-truth oracle for the arm, embedded in the ad request.

**`ylv`/`ybv` look like a floor and are a build number.** Both equal `pcode-version / 10⁷`
on 1,104 of 1,104 requests where both appear (`0.1306302` vs `1306302`). *Caveat kept in the
tool's own output:* only **one** build version exists in this corpus, so the relation is
fitted on a single point and is equally consistent with `ylv` being a constant. It cannot be
told apart until a second version is observed.

Also in the vocabulary: `utf8=✓` — the Rails hidden-field idiom, on a Yandex ad request.

## Scope

14 publishers, two nights, one browser profile shape, desktop viewport, night hours MSK.
307 legible auctions is not a market. Every number here is re-derivable from the two tools
against the ledgers named above; the ledgers themselves are not published.
