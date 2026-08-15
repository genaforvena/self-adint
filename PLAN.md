# Self-ADINT — working plan

His draft is `docs/brief-2026-08-15-operator.md` and stands as written. This file is the plan the
window executes: his order of work with **one step inserted in front of it** and our corrections
folded into the steps they belong to — each one marked `OURS` and lettered inside its step, so the
set is read off the headings and never off a count kept beside them (the count said "four" while
seven were already in the file).

His parts that needed no correction and are adopted as-is: the three-valued oracle (`IN` /
`UNKNOWN` / `NOT`), the argument that segment membership is monotone and therefore `ddmin`
degenerates to binary splitting on the "which segments am I in" question while being the right
tool for "what minimal subset of behaviour puts me in X", the decoy bidder's necessity, the
latency-first choice of VM region, and the privacy invariant.

---

## Method — the planted-invariant oracle (his doc, 2026-08-15; corrections OURS)

`docs/method-planted-invariant-oracle-2026-08-15-operator.pdf` (+ `.txt`) generalises the project's
core move into a domain-independent method: **find a quantity the system cannot corrupt because it
needs it itself, put your own mark in it, read the output.** The shift it names is the load-bearing
one — not *observe the system* but *control its input*: without input control any observation yields
correlation, with it the provenance of the fact is known and the reading is causal. Adopted as the
frame for Steps 4–6.

Adopted as written: the four applicability conditions (a write channel · an observable output, one
bit is enough · an invariant exists · the mark is globally unique); the route to finding the
invariant through the pipeline's *loss function* — what is preserved exactly is what is monetised;
the ladder of what can be extracted (connectivity → topology → time → provenance → inference model);
the one-sided oracle with `k` computed from the observed occasion rate rather than taken as a
constant; and the phase order, passive read **before** injection, because injection contaminates the
baseline irreversibly. The strongest single thing in the document is that the rival explanation is
written *into* the method: a field may arrive untouched out of **code inertia** rather than value,
and the two are told apart by an external sign (is there a market for that field) — never by the
data itself.

Six notes below. **M1–M4 are corrections to the method**; M5 is a naming, M6 an ordering fact.

**M1 (OURS) — the false `NO` is closed, the false `YES` is not.** The document handles the one-sided
direction correctly: absence of return proves nothing. It does not handle the other failure — a mark
can come back by a path other than the hypothesised one. Two concrete shapes: the SDK caches the
value and re-emits it inside the same session, and two of our vantages sit on one upstream, so the
"second confirmation" is a copy of the first observation rather than an independent one. Hence a
**fifth applicability condition: the return vantage must be one our own injection could not have
fed**, and it is established before any `YES` is trusted, not after. This is the mesh's own
`a-cached-sample-reread-faster-than-its-producer` and `two-methods-agreeing-is-not-corroboration`,
applied to a system we do not own — where it is harder, because we cannot inspect the wiring that
would prove independence and must infer it from the outside.

**M2 (OURS) — "the moment the mark stops arriving" is not the retention period.** It is a
right-censored lower bound: `min(true TTL, our observation window, occasion frequency, overwrite by
a fresher mark)`. The occasion base rate the document already computes for `k` must enter the time
estimate as well, and the result is reported as a **lower bound**, never as a storage term — the
error is one-directional and always flatters the system (measured retention comes out shorter than
real). Formally this is survival analysis with right censoring: with the occasion rate in hand a
Kaplan–Meier-style estimate is available; a single "last seen" date is not an estimate at all.

**M3 (OURS) — condition 4 is necessary and not sufficient.** The mark must be unique **in the key
space** and **typical in the field's own distribution**. A statistically odd value is caught by
anti-fraud before it is caught by anything else, and then the quantity measured is the fraud filter,
not the pipeline — with the side effect of landing in a sample that gets discarded. Practical form:
draw the mark from the value distribution Step 4's passive read observes, never from a random
string. Uniqueness buys proof; typicality buys survival, and the method needs both.

**M4 (OURS) — step 5 of the method (the plausible false fact) is understated as "the point of no
return for the baseline".** A planted fact propagates to third parties and is resold; it cannot be
recalled; and at that step his own caveat — *the method yields facts about paths, not intentions* —
stops holding, because the profile is no longer being read but shaped. The privacy invariant
(`CLAUDE.md` rule 2), which today governs the **inbound** direction, is extended **outbound**: a
planted false fact is chosen to be harmless in every downstream use — nothing health-, credit-,
employment-, political- or location-adjacent — and the choice is his, on the record, not the
window's.

**M5 (naming, not a correction).** The nearest formal relative is not the canary token but the
**chosen-plaintext attack**: the move from ciphertext-only (passive observation) to chosen
plaintext, which is a strictly stronger model — that is exactly the upgrade this method buys. The
"a distinct mark per source" step (topology) is **traitor tracing**, whose collusion theory applies
directly here, since data pooling between exchanges *is* collusion in that theory's sense.

**M6 (ordering).** The cheapest planted invariant needs no seat and no money: the **GAID reset** of
Step 6 — `t0` is known exactly and it answers the most interesting question available about the box
(does the profile survive an ID reset). Correction to how this was first put to him by voice, where
it was called simply free: it is free in money and independent of Step 3, but **not free in
baseline** — the reset is destructive to the very profile state being measured. The Step 6 ordering
constraint therefore stands unchanged: his explicit go, and only after the passive baseline exists.
Cheap is not the same as reversible.

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

This is the step the planted-invariant method governs: **M1** (a `WIN` counts only from a vantage
our own bid could not have fed), **M2** (any time figure read out here is a lower bound), **M3** (the
mark is drawn from the distribution Step 4 observed, not invented) all bind here, and **M4** binds
the moment a planted value stops being a meaningless token and starts being a plausible claim about
him.

## Step 6 — GAID reset → repeat the battery → profile recovery curve (his step 5)

Unchanged, and agreed: this is the strongest demonstration available here. It is also destructive
to his own profile state, so it runs on his explicit go and after the passive baseline exists —
resetting before there is a baseline destroys the comparison the demonstration is made of.

Read together with **M6**: in the method's own vocabulary this step is the cheapest planted
invariant in the project — no seat, no money, `t0` exact — which makes it tempting to pull forward.
It is not gated on Step 3; it *is* gated on Step 4. The recovery curve is measured against the
baseline or it is not measured at all, and **M2** applies to the curve's tail exactly as it does to
any other decay reading here: the point where the old profile stops showing is a lower bound on
re-linking time, censored by his own impression frequency.

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
