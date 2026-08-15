# Research note — hostname → exchange mapping (Step 0)

Companion to `data/exchange-domains.tsv`. 78 rows, written 2026-08-15.

This note exists because the table it accompanies is easy to over-read. The table maps an
**observed hostname** to an **organisation** and to that organisation's **seat policy**. It does
not — and cannot — map an observed hostname to *who sold that impression*. The gap between
those two statements is the whole content of this note, and it is the gap that decides whether
Step 0 can actually pick an exchange for Step 3.

**Not for publication.** Per `CLAUDE.md` rule 6 and the repo charter, nothing here goes public
without his explicit go. This file is a working artifact.

---

## 1. Method, and the evidence tiers used in the `source` column

Web search turned out to be the *weakest* available method and was used only for seat policy.
The strong method is reading the vendors' own Android SDKs.

Every row's `source` field names which tier it rests on:

| tag in `source` | what it means | strength |
|---|---|---|
| `own-extract <maven coord>` | I downloaded the official SDK artifact and read the hostname out of its constant pool myself. The field name (`SOMA_API_URL`, `adserver_endpoint`, `DtbConstants.AAX_HOSTNAME`) usually comes with it, which is what makes the *purpose* claim defensible and not just the name. | strongest |
| `own-dns 2026-08-15` | I resolved the name over DNS-over-HTTPS against Google and/or Cloudflare, bypassing this node's ad-blocking resolver (which returns `::` for everything and would have manufactured false negatives). Proves **existence and ownership** (via the CNAME chain), never purpose. | existence only |
| a vendor URL | official documentation. | strong for policy, weak for hostnames — see §5 |
| `delegated ... not re-verified` | a subagent reported it and I could not reproduce it in the time available. Flagged explicitly and capped at `medium`. | weakest; treat as a lead |

Fifteen SDKs were fetched and string-dumped: AppLovin 13.3.1, Unity Ads 4.16.2, ironSource
mediationsdk 8.10.0, Vungle 7.5.0, InMobi 10.1.4, Chartboost 9.2.1, Yandex mobileads 7.14.1,
myTarget 5.27.1, Amazon aps-sdk 11.1.1, Criteo 7.0.0, Fyber marketplace-sdk 8.3.7, Smaato
module-core 22.7.2 **and** 23.2.1, Meta audience-network-sdk 6.20.0, Google play-services-ads
24.5.0, AppsFlyer 6.17.0, Adjust 5.4.1, Branch 5.19.0, Amplitude 1.21.3.

Three of those extractions produced **negative** results that matter more than the positives, and
they are recorded as such in the table: Mintegral, Pangle and Amplitude ship **no hostname
literals at all**.

---

## 2. Mediation is not an exchange, and this is the central hazard

An in-app ad impression typically passes through three distinct layers, and only one of them
sells anything:

1. **The mediation SDK** decides *which* demand source gets the impression. AppLovin MAX,
   ironSource/LevelPlay, Google AdMob mediation, Chartboost Mediation, DT FairBid.
2. **The demand source** — a network, or an exchange — runs an auction or serves from its own
   direct demand.
3. **The buyer** — an external DSP holding a seat at that exchange, or the network's own
   advertisers.

A hostname observation lands on layer 1 or 2. The seat you would buy sits at layer 3.

Three concrete consequences for reading the table:

**An AppLovin host does not mean AppLovin sold the impression.** MAX is *both* a mediation layer
and an exchange. The MAX In-App Bidding Exchange "collects bids for individual impression
opportunities from SDK-based demand partners and DSPs via the OpenRTB 2.5 protocol"
([AppLovin](https://support.applovin.com/en/max/demand-partners/auction-dynamics)) — so a
request to `ms.applovin.com` may resolve into a win by AppLovin's own ALX demand, by a
third-party SDK network mediated *through* MAX (Unity, Mintegral, Vungle…), or by a DSP with an
ALX seat. The wire looks identical in all three cases. This is why `mediation.unity3d.com`
carries `role=mediation` and `rtb_exchange=unknown` rather than `yes`.

**A Google host does not mean Google's AdX sold the impression.** Google Ad Manager's Open
Bidding requests bids from "eligible Authorized Buyers, third-party exchanges, and/or mediation
networks" ([Google](https://support.google.com/admanager/answer/7128453)), and the publisher's
bill separates "revenue from third-party exchanges and AdX." The SDK talks to
`googleads.g.doubleclick.net` regardless of which of those wins. So the single highest-volume
hostname on the phone is the *least* informative about who is buying.

**Absence of a hostname is not absence of a seller.** Index Exchange appears to publish no in-app
bid endpoint of its own: Prebid's `ix` bidder ships disabled with a placeholder endpoint, and IX's
mobile guidance is for the publisher's Prebid Server host to add `ix` as a bidder. In-app IX demand
therefore arrives over the *publisher's* Prebid Server hostname. A capture with zero
`casalemedia.com` hits does not establish that IX is absent from the supply path.

**Practical rule for Step 0:** the bundle→exchange table PLAN.md asks for is honestly a
**bundle→demand-stack-entry-point** table. Ranking exchanges by "how much of his traffic they
carry" ranks *SDK integrations*, not *sellers*. That is still the right thing to measure — it
bounds which exchanges *could* be selling him — but the ranking must not be reported as a
measurement of sale volume.

---

## 3. What a hostname observation can and cannot prove

**Can prove:**
- That a given SDK is integrated in a given app, and fired.
- That the impression opportunity was *offered* into that org's stack.
- Which orgs are structurally reachable at all from his device — the genuine Step-0 decision.

**Cannot prove:**
- Who won the auction.
- Whether an auction happened at all (a waterfall fill is not an auction).
- Whether the impression was resold onward (`schain` in the bid request would tell a *buyer*;
  nothing on the device tells the *device owner*).
- What the bid request contained about him — which is the actual object of this project and is
  only observable from the seat, never from the phone.

That last point is worth stating plainly, because it is the honest limit of Step 0: **no amount of
device-side observation answers the project's real question.** Step 0's job is narrower and
sufficient — stop Step 3 from opening a seat at an exchange his phone never reaches.

### Six traps that will silently break a hostname classifier

These are all confirmed in the table and each one produces a *false negative* — the failure mode
that looks like a clean result.

1. **Backup and alternate domains under a different registrable name.** AppLovin's
   `*_backup_endpoint` constants all live on `applvn.com`, not `applovin.com`. Adjust ships
   `adjust.io` alongside `adjust.com`. AppsFlyer ships `appsflyersdk.com` alongside
   `appsflyer.com`. myTarget has `alt-ad.mail.ru` on the same IP as `ad.mail.ru`. A one-domain rule
   per vendor is wrong for at least four vendors here.
2. **Per-account and per-publisher hostname prefixes.** AppsFlyer composes
   `https://%slaunches.%s/api/v…` where the leading `%s` is a per-account prefix; OpenX uses
   `<publisher>-d.openx.net` (literally the Prebid `delDomain` param); Telaria uses
   `<supplyCode>.ads.tremorhub.com`. Exact-host matching cannot work — suffix matching is
   mandatory, which is why the table's `domain_pattern` is specified as a suffix.
3. **Server-supplied host overrides.** Chartboost 9.x accepts per-endpoint host overrides from its
   own `/api/config`, validated only against the `.chartboost.com` suffix. Amazon's APS SDK
   bootstraps from `mads.amazon-adsystem.com/msdk/getConfig` and then uses runtime-supplied
   regional hosts. Huawei's HMS Ads SDK hard-codes **no** ad hostnames at all and resolves
   everything through its GRS service. For these three, a static list is a snapshot, not a contract.
4. **An endpoint on a domain with no vendor string in it.** `telemetry.tk0x1.com` is a Chartboost
   endpoint whose name — through two CNAME hops to `ov1o.com` — contains nothing identifiable.
   No suffix rule catches it. It is in the table as a standing exception.
5. **Config fetched over DNS, with no TLS connection.** Smaato's
   `PUBLISHER_CONFIGURATION_DNS_URL = sdk-hb.smaato.net` has no A record (I confirmed:
   `NOERROR` with no answer). An SNI-only capture will never see it. **The MITM proxy in PLAN.md
   Step 0 must be paired with a DNS log**, or this class of endpoint is invisible by construction.
6. **Multiplexing over a shared host.** Meta Audience Network ad requests go to
   `graph.facebook.com/network_ads_common` — the ordinary Graph API host. SNI and DNS cannot
   separate a Meta ad request from an app event or any other Graph call. Meta is only
   distinguishable at the *path* level, i.e. only inside the MITM, never from the router's flow log.

### Three names in the original brief that do not exist

Recorded because a dead entry in a match list is a permanent silent miss, and because it shows the
brief's hostnames were assumptions rather than observations:

- **`mobileads.mobile.yandex.net` — NXDOMAIN.** Confirmed twice, against both Cloudflare and
  Google DoH, with an SOA returned from `ns3.yandex.ru`. The Yandex Mobile Ads SDK's compiled-in
  host is `adsdk.yandex.ru` (own-extract, `com.yandex.android:mobileads:7.14.1`), and
  `mobile.yandexadexchange.net` also exists and resolves to a Yandex IP.
- **`ia.inner-active.mobi` — NXDOMAIN.** The Digital Turbine ad-request host is
  `wv.inner-active.mobi` (own-extract, `com.fyber:marketplace-sdk:8.3.7`).
- **`adx.vungle.com` — NXDOMAIN.** The Vungle exchange host is `adx.ads.vungle.com`
  (own-extract, `com.vungle:vungle-ads:7.5.0`).

Also dead, from the same sweep: `aan.amazon-adsystem.com`, `ads.smaato.net`,
`rtb.mtgglobals.com`, `adx.rayjump.com`, `gads.pubmatic.com`, `hb.pubmatic.com`,
`exchange.rubiconproject.com`. And `an.facebook.com` *does* resolve to Meta infrastructure
(verified against a control query proving `facebook.com` is not a wildcard) but nothing
establishes its purpose — it is in the table with `role=unknown` rather than the guess its name
invites.

---

## 4. The decision this table feeds: who would actually take his seat

This is the question Step 0 exists to answer, and it splits into three groups. **None of this is a
quote.** No vendor in this table publishes a minimum spend, a QPS floor, or an onboarding fee.
Every "yes" below means *a documented path exists for an external DSP to receive bid requests* —
it does not mean a one-person buyer will be accepted, and that residual risk cannot be closed from
the outside. It is closed by sending the questions in PLAN.md Step 2.

### 4a. Structurally closed — do not spend effort here

- **Google (AdX / Authorized Buyers).** The eligibility statement is explicit and disqualifying:
  buyers "must buy on behalf of multiple advertisers"
  ([Google](https://support.google.com/authorizedbuyers/answer/6138000)). A seat whose entire
  purpose is to observe the profile of *one* device fails that test as written. This matters more
  than any other row, because Google is almost certainly the largest single share of his traffic.
  **The largest exchange on his phone is the one he cannot buy from.**
- **Meta Audience Network.** Demand comes from Meta advertisers via Meta's own buying surface;
  no external OpenRTB buyer path was found. `rtb_exchange=no`.
- **Pangle / ByteDance.** Self-describes as "the ad network of TikTok for Business" with demand
  from "direct TikTok For Business advertisers" ([Pangle](https://www.pangleglobal.com/)) and
  publishes no external-DSP integration. Marked `no` at `low` confidence — this is absence of
  evidence, not proof of a closed door.
- **Criteo.** A DSP, not a neutral exchange. Its SDK's two hosts carry Criteo's own direct-bidder
  demand. Its SSP (Commerce Grid) is a separate product not reachable from that SDK path.

### 4b. Documented external-seat path — the real candidates

Ordered by how *specific and reachable* the documented onboarding is, not by traffic share:

1. **Yandex** — the strongest candidate for this operator, and the one Step 0 most changes the
   answer on. Yandex publishes an explicit DSP-to-SSP connection procedure
   ([yandex.ru/dev/rtb](https://yandex.ru/dev/rtb/doc/api/connection.html)) with a **named intake
   address** (`newpartners@yandex-team.ru`) that supplies access, **a contract template**, and
   connection details. The doc states no minimum budget and no IP-whitelist requirement; it does
   require manual creative moderation and that responses return to the source IP. It is also the
   only candidate whose legal entity, currency, language and *latency* (PLAN.md's p99 100–120 ms
   constraint, the stated driver of VM region choice) are all natively favourable from Russia.
2. **VK / myTarget** — publishes a Buyer protocol *and* a Buyer in-app bidding protocol on
   OpenRTB 2.x ([targetmycom.gitbook.io](https://targetmycom.gitbook.io/openrtb_spec)), with
   partner intake at `ads.vk.com/partner`. Same jurisdictional and latency advantages. Onboarding
   terms are not public.
3. **Digital Turbine / DT Exchange (Fyber)** — the most openly advertised intake in the whole
   table: a public self-serve **"Questionnaire for Bidders"** form at
   `https://dsp-form.prod.fyber.com/`, plus public OpenRTB 2.5 specs. If the goal were purely
   "get *a* seat with minimum friction", this is the front door.
4. **Smaato / Verve** — public DSP onboarding with a *stated process*: credit check, contract, tech
   questionnaire, then a test account with a Bid Response Compliance stage and a Traffic
   Validation stage ([Smaato](https://wiki.smaato.com/display/DSP/DSP+Onboarding+Guide)). The
   most transparent about what onboarding actually involves — which is exactly the kind of detail
   Step 2's questions need.
5. **AppLovin (ALX)**, **Unity / ironSource Exchange**, **Vungle / Liftoff**, **InMobi**,
   **Chartboost** — all five publish DSP-facing OpenRTB specs (regional endpoints, auction type,
   timeouts, test apps). All five are commercially rep-mediated with no published terms. High
   in-app traffic share; unknown willingness to take a very small buyer.
6. **PubMatic** — publishes a real **Demand Partner Master Services Agreement**
   ([PubMatic](https://pubmatic.com/legal/demand-service-agreement/)), i.e. a contractual path
   exists in public. Its in-app path (OpenWrap SDK, `ow.pubmatic.com`) is real but was not
   binary-verified here.

### 4c. Exchange yes, seat unknown — do not plan around these

**Magnite**, **Index Exchange**, **OpenX**, **Amazon APS**, **Huawei/Petal**, **Mintegral**. All
six run real exchanges; none publishes an onboarding path a small independent buyer could assess
from outside. Magnite's bidder docs are login-gated and route to an account rep; IX has publicly
positioned itself as sell-side-only; Amazon names 20+ demand partners but no intake; Mintegral's
public OpenRTB spec names no endpoint, no QPS, no whitelist and no contact at all.

### 4d. The uncomfortable structural point

**The orgs most likely to take a small seat are not the orgs most likely to carry his traffic.**
Google and Meta between them plausibly dominate the volume on an ordinary Android phone, and both
are closed. Yandex and VK are open *and* likely significant for a Russian device — that
coincidence is the single most valuable finding in this pass, and it is exactly the finding that
Step 0 was inserted to produce before the irreversible step. But it is a *hypothesis about
his device* until the MITM capture ranks the actual hostnames, which is the artifact PLAN.md
Step 0 still owes.

---

## 5. What is still unverified, stated plainly

- **Vendor documentation is a poor source for hostnames.** Six of the vendors here publish no
  usable domain list at all (AppLovin, Mintegral, Pangle, Chartboost, Digital Turbine, and
  ironSource post-migration), and Chartboost's nearest equivalent — its Charles-proxy page — names
  `api.chartboost.com`, which the SDK reportedly does not contact. The SDK binaries are not merely
  the best source; for most of this table they are the only one.
- **Mintegral and Pangle have no hostname evidence of any tier above DNS existence.** Both build
  endpoints at runtime; the Pangle artifact is a shell over a downloaded payload. Every Mintegral
  and Pangle row is `low` confidence and should be treated as a lead, not a fact.
- **Huawei rows rest on a naming convention plus DNS**, with one Akamai CNAME chain through
  `huawei.com` doing the org attribution for the `dre` host. The `drru` (Russia) form was *not*
  independently probed.
- **Amazon's runtime host set is region-dependent** and was sampled from exactly one egress.
- **The rows tagged `delegated ... not re-verified`** — `telemetry.tk0x1.com`'s binary provenance,
  the Yandex `mobile.yandexadexchange.net/v4/ad` doc quote, `tremorhub.com`, AppMetrica,
  PubNative's SDK constants, `r.my.com` — are subagent findings I could not reproduce in this
  pass. Per the mesh rule that a subagent's report is a claim and not an artifact, they are capped
  at `medium` and flagged in-row. Re-extract before anything depends on them.
- **Nothing in this table has been checked against his actual device.** Every row is a claim about
  what an SDK *would* do. The measurement PLAN.md Step 0 asks for — MITM capture plus the router's
  DNS log, producing observed request counts per bundle — has not been taken. This table is the
  *decoder* for that capture, not a substitute for it.

---

## 6. Sources

Vendor documentation cited above, in the order used:

- Google Authorized Buyers eligibility — https://support.google.com/authorizedbuyers/answer/6138000
- Google Open Bidding — https://support.google.com/admanager/answer/7128453
- Google mobile session capture (host names) — https://support.google.com/admanager/answer/6206401
- Google app conversion tracking API — https://developers.google.com/app-conversion-tracking/api/request-response-specs
- Meta Audience Network ad-request validation — https://developers.facebook.com/docs/audience-network/setting-up/test/validate-ad-requests/
- AppLovin MAX auction dynamics — https://support.applovin.com/en/max/demand-partners/auction-dynamics
- AppLovin oRTB spec for DSPs — https://support.applovin.com/en/max/demand-partners/demand-side-platforms/applovin-ortb-specification/introduction
- Unity demand-side partners — https://docs.unity.com/demand-side-partners/en-us/manual/GettingStarted
- ironSource Exchange DSP onboarding — https://docs.unity.com/en-us/grow/is-ads/user-acquisition/ironsource-exchange/getting-started-dsp-onboarding
- Vungle Exchange OpenRTB 2.5 — https://support.vungle.com/hc/en-us/articles/360045953431-Vungle-Exchange-OpenRTB-2-5-Integration-Guide
- InMobi OpenRTB overview — https://support.inmobi.com/advertise/ad-tracker/overview-ortb/
- Chartboost partner OpenRTB 2.6 — https://docs.chartboost.com/en/partners/exchange/openrtb/2.6-bid-request/
- Digital Turbine bidder questionnaire — https://dsp-form.prod.fyber.com/
- Fyber Marketplace OpenRTB 2.5 — https://marketplace-demand.fyber.com/docs/openrtb-25-spec
- Smaato DSP onboarding — https://wiki.smaato.com/display/DSP/DSP+Onboarding+Guide
- Verve integration process — https://developers.verve.com/docs/overall-integration-process-2
- Mintegral ADX OpenRTB spec — https://github.com/Mintegral-official/mtgrtb_doc/blob/master/mtgrtb_en.md
- Pangle — https://www.pangleglobal.com/
- Yandex DSP→SSP connection — https://yandex.ru/dev/rtb/doc/api/connection.html
- Yandex RTB overview — https://yandex.ru/support2/partner/ru/technologies/rtb
- VK Ad Network OpenRTB protocols — https://targetmycom.gitbook.io/openrtb_spec
- Amazon Publisher Services, programmatic bidders — https://aps.amazon.com/aps/solutions-for-programmatic-bidders/
- Petal Ads / Huawei exchange — https://www.adexchanger.com/mobile/petal-ads-hopes-to-blossom-with-global-advertisers/
- PubMatic Demand Partner MSA — https://pubmatic.com/legal/demand-service-agreement/
- Magnite OpenRTB spec (login-gated) — https://help.magnite.com/help/openrtb-specification
- Index Exchange DSP OpenRTB fields — https://kb.indexexchange.com/dsps/open-rtb/list_of_supported_openrtb_bid_request_fields_dsp.htm
- OpenX docs / Prebid delDomain — https://docs.openx.com/ , https://docs.prebid.org/dev-docs/bidders/openx.html
- Criteo Commerce Grid — https://www.criteo.com/platform/commerce-grid/

SDK artifacts read directly (Maven Central, Google Maven, and vendor Maven repositories) are
named inline in each row's `source` field with their exact group:artifact:version.
