# Self-ADINT — working plan

His draft is `docs/brief-2026-08-15-operator.md` and stands as written. This file is the plan the
window executes: his order of work with **one step inserted in front of it** and four corrections
folded into the steps they belong to. Everything below that is not his is marked as ours.

His parts that needed no correction and are adopted as-is: the three-valued oracle (`IN` /
`UNKNOWN` / `NOT`), the argument that segment membership is monotone and therefore `ddmin`
degenerates to binary splitting on the "which segments am I in" question while being the right
tool for "what minimal subset of behaviour puts me in X", the decoy bidder's necessity, the
latency-first choice of VM region, and the privacy invariant.

---

## Step 0 — which exchanges actually sell his phone? (OURS; cost ≈ 0; before anything irreversible)

**Why it exists.** His step 2 — seat onboarding — is the only irreversible step in the project
(legal entity, agreement, whitelist, deposit). Nothing in his draft measures, before that step,
**which exchange his own device's traffic actually flows through.** If his apps sell through
exchange A and the seat is opened at exchange B, his device never appears in the sample and the
failure is discovered after the money and the contract. The choice of exchange is currently a
guess, and it is a guess that can be turned into a measurement for free.

**Method** — read from the device side, not the buy side:

- MITM proxy on his phone (own device, own CA) → the ad SDK endpoints each app talks to.
- The router's DNS/flow log, read-only. **The router is instruction-only on this mesh — never
  change its settings; if a setting is needed, write the instruction and hand it to him.**
- Cross-check: `app.bundle` values he cares about → which SSP/exchange domains they resolve to.

**Artifact:** a `bundle → exchange` table with observed request counts, on disk, plus the list of
exchanges ranked by how much of *his* traffic they carry. Not "app X uses ads" — a table.

**Correction 0a (OURS, 2026-08-15, after the domain research).** Read literally, that artifact
cannot be built from the wire, and calling it `bundle → exchange` overstates what any hostname
observation can prove. **Mediation is not an exchange.** A request seen going to an AppLovin MAX
or an AdMob host may be sold by that org's own exchange, by AdX, by a third-party Open Bidding
exchange, or by a mediated network with no auction at all — *the wire looks identical in every
case.* What the capture yields is honestly a **`bundle → demand-stack entry point`** table: which
orgs' SDKs actually carry his impressions, ranked by observation count. That is enough to decide
which exchange to approach — the decision Step 0 exists for — and it is not sale volume. It must
never be reported as sale volume.

**Correction 0b (OURS).** Two structural facts the research settled, both of which shape every
step after this one:

- The orgs most likely to carry his traffic (Google, Meta) are exactly the ones that will not
  sell him a seat: Google's Authorized Buyers requires buying *"on behalf of multiple
  advertisers"*, which disqualifies a self-directed seat as written; Meta, Pangle and Criteo are
  closed too. The orgs that *are* open — Yandex, VK, Digital Turbine, Smaato/Verve, AppLovin,
  Unity/ironSource, Vungle/Liftoff, InMobi, Chartboost, PubMatic — are plausibly significant on a
  Russian device, but "plausibly" is a hypothesis about **his** phone until the capture ranks real
  hostnames. Yandex publishes a DSP→SSP intake with a named address and a contract template, and
  is the strongest candidate on jurisdiction and latency both.
- Some demand paths are **invisible to a capture by construction** and have to be closed from a
  second vantage or declared blind: Smaato fetches configuration over DNS with no TLS connection
  (so a connection log never sees it — it needs a DNS log beside it), and Meta multiplexes ad
  requests over `graph.facebook.com`, separable only by URL path and therefore never from a router
  flow log. A method that cannot see a demand path must say so; the org is `UNKNOWN`, not absent.

**What it decides:** which exchange to approach at all, and whether question 4 below (point of
presence) even has an acceptable answer. If no reachable exchange carries his traffic, the project
stops here having cost nothing.

## Step 1 — receiver + decoy bidder, locally, on synthetic requests (his step 1)

Go, ~200 lines, JSONL to disk, DuckDB offline. Three corrections fold in here:

**1a. "win rate нулевой, расход нулевой" is an assumption, not a property of the design.**
A bid below the floor wins whenever the floor is absent or not applied, and in a first-price
auction with no floor it wins outright. Rare, but it happens. Consequences to build for:

- a **valid fallback creative** must exist and be servable. Winning and failing to serve raises
  `error rate` — one of the six ranking axes his own draft lists, and a more expensive one to
  damage than the bid rate it was protecting.
- a **hard spend cap** enforced on our side, plus an alert on the first non-zero win.
- the win counter is an artifact, not a hope: if it is never asserted, "zero spend" is a silent
  fallback in the doctrine's exact sense.

**1b. The privacy invariant is a construction, not a discipline** (see `CLAUDE.md` rule 2).
The parser-before-buffer filter is right and insufficient: the leak path is a panic handler that
dumps the request body, a framework access log, or a debug metric labelled by IFA. Test for it
explicitly — drive a foreign-IFA request through the panic path and assert nothing lands on disk.

**1c. Gates must be seen red.** Synthetic requests are the whole test surface at this stage, so
they carry the weight: break the ifa filter and watch the foreign row appear, then restore it.

## Step 2 — the questions to the exchange (his list, + one; drafted here, sent by him)

His five stand. Adding a sixth, which decides whether group testing is viable at all:

6. **Does targeting support OR over a segment list, and what is the maximum list length?**
   His own arithmetic — ~250–300 tests instead of 2000 — depends entirely on one bid testing a
   whole pool. Many exchange and DSP interfaces cap the segment list, and some apply AND semantics
   or charge per segment. If the cap is small, binary splitting's advantage collapses toward linear
   and the active oracle's cost model has to be rewritten before it is built, not after.

Question 5 (contractual retention of bid request data outside bidding purposes) is the most likely
project-killer. Treat its answer as a gate on everything downstream, exactly as he wrote it.

## Step 3 — seat onboarding (his step 2)

**His call alone.** Legal entity, agreement, whitelist, QPS negotiation. The window drafts
documents and tracks the thread; it does not open accounts, sign, or pay.

## Step 4 — a week of passive collection (his step 3)

Report: what arrives in `user.data[]`, who is in `eids[]`, the `geo.type` distribution per
`bundle` (his stated primary result), and the observed frequency of his own impressions.

**Correction 4a — `k` is circular on a short window.** `k` is derived from the observed impression
frequency, and it is applied to the same window that estimates it. Early on, that estimate has
almost no power, and a `NOT` issued from it is fabricated. So the verdict lattice gains an explicit
state:

    IN         — one WIN
    UNKNOWN    — default
    NO-BASIS   — not enough observed impressions to compute k at all   (OURS)
    NOT        — k confirmed consecutive LOSS, and only once k is computable

`NO-BASIS` must be renderable in every report. Collapsing it into `UNKNOWN` is survivable;
collapsing it into `NOT` fabricates a negative, which is the one direction his oracle design
exists to prevent.

**Correction 4b — the throughput figures are unmeasured.** "30–100 оказий в сутки", "1–3 тысячи
бит в месяц", "единицы-десятки попаданий в неделю" are estimates, and this step is what measures
them. They are not quoted as facts anywhere until this step reports. (One clause of his brief
carrying a volume figure arrived garbled — see the flag in the fact base; ask him rather than
reconstruct.)

## Step 5 — active oracle, only if the payload is empty or hashed (his step 4)

Unchanged. Note his own condition: `WIN → состою` holds only if the exchange applies targeting
honestly, which question 6 partly probes.

## Step 6 — GAID reset → repeat the battery → profile recovery curve (his step 5)

Unchanged, and agreed: this is the strongest demonstration available here. It is also destructive
to his own profile state, so it runs on his explicit go and after the passive baseline exists —
resetting before there is a baseline destroys the comparison the demonstration is made of.

---

## Infrastructure

One VM, region chosen by the exchange's point of presence (his reasoning, adopted: the p99
100–120 ms requirement is punished harder than a low bid rate). Valid TLS, static IP for the
whitelist.

**Mesh constraint:** that VM is outside the mesh — no Tailscale, no mesh reachability, no mesh
credentials, and no node routes through it. It cannot be hosted on `mesh-home` in any case; public
ingress here is closed.

## What this project does not give (his list, kept)

Movement tracking at minute granularity; observation of anyone but himself; real time.
