---
title: I went looking for my ad profile inside the bid request. It wasn't there.
tags: privacy, adtech, opensource, datascience
canonical_url: https://github.com/genaforvena/self-adint
---

When an ad is auctioned for your phone, a bid request goes out that describes *you* —
segments, identity graphs, consent state, location precision. Buyers see that description for
free, before anyone pays for anything.

So I tried to read mine. One operator, one phone, his own data, everything public:
[genaforvena/self-adint](https://github.com/genaforvena/self-adint), data CC0.

The result is a negative for the naive form of that idea, and the negative is more interesting
than the confirmation would have been.

## The thesis was not unproven. It was unmeasured.

I had spent four days measuring **prices and outcomes** — which bidders answer, which stay
silent, what comes back when one bids. Then I looked at what my own capture had kept of the
*payload*, and the answer was: nothing. Three narrow parsers had read `post_data`, extracted a
verdict, and discarded the body. On disk, the field that was supposed to hold what the request
*said* was a boolean. 579 candidate request rows from one night retained not one key of it.

That is worth naming as a failure mode of its own: **a pipeline can be perfectly honest, gated,
and reproducible while quietly answering a different question than the one you asked.** No gate
fires, because every gate was written about the question the code actually answers.

So: keep the body. Then keep the query string too — which turned out to matter more.

## The truncation that reported the opposite of what was there

The first fix recorded the request URL as `url[:500]`. Every exchange that asks by **GET**
therefore had its question cut mid-query-string, with nothing on the row saying it had been
cut. On a fresh sweep, **174 of 380 request rows were longer than 500 characters**; one
`yandex.ru` request reached 9529.

`ssp01.rambler.ru` is the case that names the cost. Under the truncating schema it read
`no-body` — it sends nothing, it is invisible, it does not appear in any count. It is a real
SSP, and its query string carries `adtech_uid` (a uuid), `publisher_uid`, `rq_sess`, a
`top100_session_id` map keyed by publisher site id, and six Adfox `puidN` targeting parameters.

**A body-only reading of this market reports the opposite of what is there.** If you are
measuring an ad stack: the GET half is not a lesser half.

## 23 bid-request paths. 18 carry no audience field at all.

The census reports key-path *shapes* and never values, so the finding is publishable while the
payload is not. Restricted to paths my capture tool had already classified as bid candidates
(its verdict, read off the row — the census does not re-decide what an auction is by hostname),
with at least 4 bodies read:

| | count |
|---|---|
| POST bid-request paths | **23** |
| carrying no audience field the taxonomy recognises | **18** |
| with at least one filled audience slot | **4** |
| with a present-but-always-empty audience slot | **2** |

The 18 receive the Adfox `places` protocol and nothing else — seven to thirteen key paths,
every one of them about the placement, the format, or the page:

```
places[].id  places[].placementId  places[].sizes[][]  places[].targetRef
settings.currency  settings.windowSize.width  settings.windowSize.height
```

So on this path, for most of the demand side, **the audience profile does not arrive inside the
bid request the browser sends.** The enrichment happens server-side, past the client's vantage.
What *is* client-visible is a separate channel: the cookie-sync hops, which carry identity and
no auction, and the wrapper operators' own requests, which are not bid requests at all.

## The most informative state is `present and empty`

`ssp-rtb.sape.ru` asks for `places[].sapeFpUids[]` — an eids-shaped identity graph, an array of
`{source, id}` pairs — on every single bid request. On this browser:

```
cold   22 empty, 0 populated
warm   20 empty, 2 populated     both on one site; two sources, a 19-digit id and a 32-hex id
```

Count that as "sape carries an identity field" and you overstate. Count it as "sape asks nothing
about the person" and you lose the result entirely. The honest reading is the third one:
**the slot exists and this browser has nothing to put in it** — a fact about the browser, not
about the protocol. `kimberlite.io` is the pure case: it asks for a stable id on all 18 of its
requests and never once gets one.

Most privacy measurement I have read collapses this state into one of the two neighbours. It is
the state that tells you what the protocol *wants*.

## Where the profile actually was

`mc.yandex.com` — Yandex Metrica, run by the same operator as the wrapper — receives a
`site-info` map in its query string with **295 key paths**. Most are the publisher's own UI
taxonomy. Among them is a block of publisher-declared visitor labels, written in plain Russian,
**as key names**:

| key path | reading |
|---|---|
| `Пользователь.Подозрительный` | "User: Suspicious" |
| `Пользователь.Лояльность.Disloyal` | "User: Loyalty: Disloyal" |
| `Пользователь.Комментатор` | "User: Commenter" |
| `Пользователь.Авторизация` | "User: Authorization" |
| `__ymu.Пользователь.Лояльность (просмотры за 30 дней).Новый` | "loyalty, views in 30 days: New" |
| `__ymu.pgUserName` | **the visitor's username** — present, empty |

`__ymu` is Metrica's `userParams`: publisher-declared attributes attached to the visitor. This
host reads **14 filled audience slots and one present-but-always-empty** — the username, which
this browser has nothing to put in.

**What this is not.** `mc.yandex.com` is the analytics endpoint, not a bid request. Nothing I
measured shows these labels reaching an auction. Yandex operates both, and Metrica segments are
usable for targeting as a matter of product documentation — but that is a link this capture does
not evidence, and I am not claiming it.

**What it is.** The plainest answer to "what does the demand side get told about a person" in
this market is not in the bid request at all. It is a publisher telling the wrapper operator, in
its own language, that this visitor is suspicious, disloyal, and has not logged in.

## The identity that does travel, and the control that makes it mean something

Paired design: 23 **cold** loads at `profile_age_s = 0` with no profile directory, and 23
**warm** loads in a named profile aged 811–2210 s. Arms adjacent on each site, one hour, one
vantage, treatment verified on the row rather than inferred from which run produced it.

Shapes cannot answer identifier persistence — `uuid` and `uuid` look the same whether a field
carries one id forever or mints a new one per load. Counting *distinct values* answers it and
publishes nothing (values are folded to one-way tags; only the count is reported).

`yandex.ru`:

| field | cold | warm |
|---|---|---|
| `?pcode-uid` | 19 distinct over 19 sites | **3 distinct over 20 sites** |
| `?pcode-icookie` | 20 distinct over 19 sites | **3 distinct over 20 sites** |
| largest span of a single value | **1 site** | **19 sites** |

One identifier travelling across 19 of 21 publishers, readable from the demand side without
buying anything, existing only because the profile has history.

The cold arm is what makes that mean something: the same field *can* be per-load, so the 19-site
span is the profile and not the field's nature. And from the same requests, two fields hold at
one distinct value **per load in both arms** — `?duid` and `?site-info.__ym.adSessionID`. They
are request nonces, not identities; `?duid` independently base64-decodes to 19 decimal digits
whose leading ten are a unix timestamp, an id minted from the clock. Without that control,
"19 distinct values over 19 cold loads" would prove nothing at all, since a fresh browser
context mints fresh state by construction. **The nonces show what a non-carrier looks like in
the same ledger.**

## One more thing the wrapper reports

`?bids` on `yandex.ru` is base64 JSON. Decoded, it is every *other* bidder's bid, reported back
to Yandex:

```
?bids~b64[].bidderName  .campaign_id  .placement_id  .error.code  .response_time
```

22 distinct bidders in one ledger, each with its campaign id and how long it took (n=93, min
416 ms, median 1397 ms, max 1703 ms). The same host carries 35 distinct `puidN` parameters — in
Adfox, publisher-declared *user targeting* parameters. The name says identity; the role is a
segment.

## The dataset leaked, and no gate could see it

`public-data/census-2026-08-20-paired.jsonl` is 2655 rows, CC0, one per (host, key path):
shape, buckets, load counts, distinct-value counts, `empty_slot` flag. The ledger it was read
from is not published and will not be.

The first version of that file leaked. Every gate in the publisher asserts a field **name** —
an explicit allowlist, four gates naming the body and full-URL fields specifically, each
going red under a mutant that adds them back. The leak was in the **content** of an allowed
field. Diffing the published file against every 12+ character token that had been on the wire
returned **349 matches**, because:

- a **key** can be an identifier — the census publishes key paths, and
  `top100_session_id.<19-digit id>` puts one *in a path*;
- a **hostname** can embed one — `2-6a871d048f403bcd46b18008.id.…`, a per-sync domain,
  published verbatim in the `host` column.

The fix redacts identifier-shaped segments of paths and hostnames to their shape, keeping short
segments so a 7-digit publisher site id — not a person, and the thing that makes the cross-site
result readable — survives. The diff that found it is now a gate. After the fix: 0 identifiers,
16 remaining overlaps, all publisher domains and field names.

I am writing that down rather than quietly fixing it, because it is the same failure the study
measures — an absence that was really a blindness — committed by the tool built to avoid it.

## Scope, before anyone uses the conclusions

- **One browser, one operator, one night.** 23 sites, 23 loads per arm.
- **Warming is ~15–37 minutes and one pass over the frame, with no logins.** This says nothing
  about a months-old profile. The *direction* of the effect is what is measured, not its ceiling.
- **The vantage is a datacentre exit, not a Russian consumer connection.** Which bidders
  participate may differ from a resident's view.
- **`sapeFpUids` populated is n=2, on one site.** The cold arm's 0/22 is the control, but the
  positive is small.
- **The taxonomy is name matching, and it under-reported four times in one day** — `adtech_uid`,
  `puidN`, `sapeFpUids`, and then every Russian-language visitor label on the market's dominant
  analytics path, because the taxonomy was written in English. *A taxonomy is a copy of the world
  in one language.* Read the headline as "no audience field **this taxonomy recognises**", never
  "no audience data".
- **`mc.yandex.com` is analytics, not an auction.** The visitor labels are not shown reaching a
  bid request, and no such link is claimed.

## Two rules I would keep for any measurement like this

**Absence is split by cause, always.** `no-bid` (they answered "no"), `no-answer` (asked, said
nothing), `blocked` (never saw the page), `unreachable` (our own link), `no-wrapper` (nothing to
ask with) are five different facts. Folding any two converts your own blindness into a finding
about the market — which, in a study about what other people can see about you, is the one
mistake that is worth nothing at all.

**A gate you have not seen fail is not a gate.** Every property above was broken deliberately
and watched go red before it was accepted.

Repo, tools, method, and the CC0 data: <https://github.com/genaforvena/self-adint>. Every
number on the front page re-derives from the published files alone — that is checkable, and
disagreeing with it does not require re-running the capture.
