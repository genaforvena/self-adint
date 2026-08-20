# step0e — what the bid request actually says about the person

**Date** 2026-08-20 · **Ledger** `data/hb-bodies-2026-08-20-paired.jsonl` (schema 7, private)
· **Frame** the 23-domain stage-B RU frame · **Vantage** one exit, verified per load from
inside the browser · **Design** paired cold/warm, arms adjacent on each site, 23 loads per arm

---

## The question this step exists to answer

The README's thesis is that *the audience profile arrives inside the bid request* — that a
person could read what the demand side is told about them without buying anything. Steps 0–0d
measured **prices and outcomes**. They never read the payload: schemas 1–5 parsed
`post_data` through three narrow parsers and discarded it, so `req_openrtb` on disk was a
boolean and 579 candidate request rows from an earlier night retained nothing about what the
request said. The thesis was not merely unproven. It was **unmeasured**.

Schema 6 kept the body. Schema 7 kept the query string, which turned out to matter more.

## The correction that made the reading possible

Schema 6 recorded the URL as `url[:500]`. Every exchange that asks by **GET** therefore had
its question cut mid-query-string with nothing on the row saying so — the same trap the body
cap is documented against, applied to the half nobody re-read. On a fresh schema-7 sweep,
**174 of 380 request rows exceeded 500 characters**; `yandex.ru` reached 9529.

`ssp01.rambler.ru` is the case that names the cost. Under schema 6 it read `no-body`, sent
nothing, and was invisible. It is a real SSP, and its query string carries `adtech_uid` (a
uuid), `publisher_uid`, `rq_sess`, a `top100_session_id` map keyed by publisher site id, and
six Adfox `puidN` targeting parameters.

**A body-only reading of this market reports the opposite of what is there.**

## What the browser sends each exchange

`tools/adint-body-census` reports key-path SHAPES and never values, so the finding is
publishable while the payload is not. Restricted to paths the *capture tool* classified as
bid candidates (its verdict, read off the row — this tool does not re-decide what an auction
is by hostname), with at least 4 bodies read:

| | count |
|---|---|
| POST bid-request paths | **23** |
| carrying no audience field the taxonomy recognises | **18** |
| with at least one FILLED audience slot | **4** — `pb.adriver.ru`, `ad.mail.ru`, `exchange.buzzoola.com`, `ssp-rtb.sape.ru` |
| with a present-but-ALWAYS-EMPTY audience slot | **2** — `ssp-rtb.sape.ru`, `kimberlite.io` |

The 18 receive the Adfox `places` protocol and nothing else — seven to thirteen key paths,
every one of them about the placement, the format or the page:

```
places[].id  places[].placementId  places[].sizes[][]  places[].targetRef
settings.currency  settings.windowSize.width  settings.windowSize.height
```

**So on this path, for most of the demand side, the audience profile does NOT arrive inside
the bid request the browser sends.** The enrichment happens server-side, past the client's
vantage. What is client-visible is a *separate* channel: the cookie-sync hops
(`sync.bumlam.com`, `ad.adriver.ru`, `bid.sspnet.tech`, `a.ussp.io`, `adx.com.ru`,
`ev.adriver.ru`), which are `url-only` and carry identity, and the wrapper operators' own
requests, which are not bid requests at all.

That is a **refinement of the thesis, and a negative result for the naive form of it.**

## The most informative state is `present and empty`

`ssp-rtb.sape.ru` asks for `places[].sapeFpUids[]` — an eids-shaped identity graph, an array
of `{source, id}` pairs — on every bid request. On this browser:

```
cold   22 empty, 0 populated
warm   20 empty, 2 populated     both on aif.ru; two sources, a 19-digit id and a 32-hex id
```

Counted as "sape carries an identity field", that overstates. Counted as "sape asks nothing
about the person", it understates and loses the result. **The slot exists and this browser
has nothing to put in it** — a fact about the browser, not about the protocol.
`kimberlite.io` reads 0 filled / 1 always empty: it asks for a stable id on all 18 requests
and never gets one.

## Where the audience profile actually is: not the bid request, the analytics call

`mc.yandex.com` — Yandex Metrica, the same operator that runs the wrapper — receives a
`site-info` map in its query string with **295 key paths**. Most are the publisher's own UI
taxonomy (page sections, menu items). Among them is a block of **publisher-declared visitor
labels, written in plain Russian, as key names**:

| key path | reading | shape on this browser |
|---|---|---|
| `Пользователь.Подозрительный` | "User: Suspicious" | `str:2` |
| `Пользователь.Лояльность.Disloyal` | "User: Loyalty: Disloyal" | `str:2` |
| `Пользователь.Комментатор` | "User: Commenter" | `str:6` |
| `Пользователь.Достижения` | "User: Achievements" | `str:7` |
| `Пользователь.Авторизация` | "User: Authorization" | `str:16` |
| `__ymu.Пользователь.Лояльность (просмотры за 30 дней).Новый` | "loyalty, views in 30 days: New" | `str:1` |
| `__ymu.authStat` | Metrica userParams | `str:13` |
| `__ymu.pgUserID` | | `str:1` |
| `__ymu.pgUserName` | **the visitor's username** | **`str:empty`** |

`__ymu` is Metrica's `userParams` — publisher-declared attributes attached to the visitor.
The host reads **14 filled audience slots and one present-but-always-empty**: `pgUserName`.
The username slot exists, and a browser that never logged in has nothing to put in it.

**What this is NOT.** `mc.yandex.com` is the analytics endpoint, not a bid request. Nothing
measured here shows these labels reaching an auction. Yandex operates both, and Metrica
segments are usable for targeting as a matter of product documentation — but that is a link
this capture does not evidence, and it is not claimed.

**What it is.** The plainest answer to "what does the demand side get told about a person" in
this market is not in the bid request at all. It is a publisher telling the wrapper operator,
in its own language, that this visitor is suspicious and disloyal and has not logged in.

## The identity that does travel

Paired design, one hour, one vantage, the two arms adjacent on each site. The treatment is
verified on the row, not inferred from which run produced it: 23 cold loads at
`profile_age_s = 0` with no profile directory, 23 warm loads in a named profile aged
811–2210 s.

Shapes cannot answer identifier persistence — `uuid` and `uuid` are the same shape whether a
field carries one identifier forever or mints a new one each load. Counting **distinct
values** answers it and publishes nothing (values are folded to one-way tags; only the count
is reported).

`yandex.ru`:

| field | cold (fresh context each load) | warm (one persistent profile) |
|---|---|---|
| `?pcode-uid` | 19 distinct over 19 sites | **3 distinct over 20 sites** |
| `?pcode-icookie` | 20 distinct over 19 sites | **3 distinct over 20 sites** |
| largest span of a single value | **1 site** | **19 sites** |

One identifier travelling across 19 of 21 publishers, readable from the demand side without
buying anything, existing only because the profile has history. **The cold arm is what makes
this mean something:** the same field *can* be per-load, so the 19-site span is the profile
and not the field's nature. Two loads (aif.ru, mk.ru) carry their own value — noted, not
explained.

`ssp01.rambler.ru`'s `top100_session_id` map moves the same way: 0–2 entries per request and
**2** distinct publisher site ids when cold, 2–4 entries and **7** when warm. The bid request
carries session identifiers for publishers the browser visited *earlier*.

### The control that keeps it honest

From the same measurement on the same requests, two fields stay at one distinct value **per
load in both arms**: `?duid` and `?site-info.__ym.adSessionID`. They are request nonces, not
identities. `?duid` independently base64-decodes to 19 decimal digits whose leading ten are a
unix timestamp — an id minted from the clock. Two readings agree.

Without this control, "19 distinct values over 19 cold loads" would prove nothing: a fresh
context mints fresh state by construction. The nonces show what a *non-*carrier looks like in
the same ledger.

## What Yandex sees that the other bidders do not

`?bids` on `yandex.ru` is base64 JSON. Decoded (`--decode`, off by default, and a decoded
path is marked `~b64` because it is a weaker claim than a field sent in the clear):

```
?bids~b64[].bidderName  .campaign_id  .placement_id  .error.code  .response_time
```

The wrapper reports **every other bidder's bid back to Yandex** — 22 distinct bidders in one
ledger (mytarget, buzzoola, adriver, mediasniper, betweendigital, gnezdo, sape, videonow,
otm, roxot, ohmybid, smi2, astralab, adwile, hybrid, bidvol, getintent and five `adfox_*`
seats), each with its campaign id and how long it took (n=93, min 416 ms, median 1397 ms, max
1703 ms).

`yandex.ru` also carries **35 distinct `puidN` parameters** in this ledger. In Adfox these are publisher-declared
*user targeting* parameters: the name says identity, the role is a segment.

## Scope, stated before the conclusions are used

- **One browser, one operator, one night.** 23 sites, 23 loads per arm.
- **Warming is ~15–37 minutes and one pass over the frame, with no logins.** This says
  nothing about a months-old profile, and the direction of the effect is what is measured,
  not its ceiling.
- **The vantage is a datacentre exit, not a Russian consumer connection.** Which bidders
  participate may differ from a resident's view; what the *browser sends* is built
  client-side and is less exposed to that, but it is not immune.
- **`sapeFpUids` populated is n=2, on one site.** The cold arm's 0/22 is the control, but the
  positive is small.
- **The taxonomy is name matching, and it under-reported four times in one day** —
  `adtech_uid`, `puidN`, `sapeFpUids`, and then every Russian-language visitor label on the
  market's dominant analytics path, because the taxonomy was written in English. *A taxonomy
  is a copy of the world in one language.* Read the headline as *"no audience field this
  taxonomy recognises"*, never *"no audience data"*. `--paths` prints the unbucketed census
  so a reader can disagree without re-running the capture.
- **`mc.yandex.com` is analytics, not an auction.** The visitor labels above are not shown
  reaching a bid request, and no such link is claimed.

## The published dataset, and the leak that nearly went with it

`public-data/census-2026-08-20-paired.jsonl` — 2655 rows, CC0, one per (host, key path),
carrying the shape, the buckets, the load counts, the distinct-value counts and the
`empty_slot` flag. The ledger it was read from is not published and will not be.

**The first version of that file leaked, and no gate could see it.** Every gate in
`adint-publish` asserts a field NAME; the leak was in the CONTENT of an allowed field.
Diffing the published file against every 12+ character token that had been on the wire
returned **349 matches** — because a KEY can be an identifier (the census publishes key paths,
and `top100_session_id.<19-digit id>` puts one in a path), and a HOSTNAME can embed one
(`2-6a871d048f403bcd46b18008.id.…`, a per-sync domain, published verbatim in the `host`
column).

The fix redacts identifier-shaped segments of paths and hostnames to their shape, keeping
short segments so a 7-digit publisher site id — not a person, and the thing that makes the
cross-site result readable — survives. **The diff that found it is now a gate**, narrowed to
the wire's VALUES plus ids hiding inside names (asserting "no wire token survives" fails on
`top100_session_id`, a field name this study exists to report). After the fix: 0 identifiers,
16 remaining overlaps, all publisher domains and field names.

Recorded here rather than quietly fixed, because it is the same failure the study measures —
an absence that was really a blindness — committed by the tool built to avoid it.

## Presented

A shareable write-up of this step: <https://claude.ai/code/artifact/67ffcca9-5901-4241-9f03-aff93616e433>
(private until shared from the page's own menu). It carries the same numbers as this document
and the same scope section; nothing is in it that is not here.

## Reproducing

```bash
python3 tools/adint-hb-capture --file <frame urls, each twice> --settle 12 --nav-timeout 30 \
  --via direct --profile-dir data/profiles/<name> --warm-arm warm --arms cold,warm \
  --arm-note "..." --out data/<ledger>.jsonl

python3 tools/adint-body-census <ledger> --candidates only --min-req 4
python3 tools/adint-body-census <ledger> --host ssp01.rambler.ru --paths --decode
python3 tools/adint-body-census <ledger> --by-arm cold warm --audience-only --min-loads 3
```

The ledger itself is **not published**: it is a description of one real person. Four gates in
`tools/adint-publish` name `req_body`, `req_body_bytes`, `req_url_full` and `req_url_bytes`
and go red under a mutant that adds any of them to the allowlist. Publishing profiles
requires deleting a gate.
