# Preregistration — does a bidder's decision to return a PRICE depend on the vantage?

**Written 2026-08-21, before the collection window it governs opens.** That ordering is
the whole point of the document and it is checkable: `tools/adint-prereg --verify` hashes
this file and compares the hash against the one recorded at registration, and every cell
captured from the window's start carries a timestamp later than this file's commit.

Registered by: `adint@mesh-home`. Task: `adint-two-week-collection-2026-08-21`.
Design this refines, not replaces: `docs/step0c-hb-ru-market-study-2026-08-18.md`.

---

## 0. Why preregister at all, in one paragraph

The pilot found something. On the corpus of 2026-08-19..20, every one of 11 priced AdFox
entries appeared on the RU arm and none on the foreign arm, while the same two bidders
were asked *more* often from the foreign vantage (`roxot` 31→6 RU vs 41→0 foreign;
`adwile` 6→5 vs 8→0; pooled 37→11 vs 49→0; one-sided Fisher p = 3.5 × 10⁻⁵). That is a
striking result resting on **11 priced records over two nights**, and the honest reading
of a two-night finding is that it is indistinguishable from a two-night artifact. The
operator asked for significance; significance is not a bigger n, it is a test that could
have come out the other way and was chosen before the data that decides it. Choosing the
test after seeing the corpus is p-hacking, and it would destroy exactly the credibility
the extra fortnight is being spent to buy.

**This corpus is published continuously and cannot be blinded.** Anyone, including the
analyst, can read the incoming cells. The defence against that is not pretending
otherwise — it is fixing the confirmatory analysis here, computing it **once**, and
labelling every figure produced before the window closes as exploratory. `adint-prereg`
enforces the "once, at the end" part mechanically rather than by intention.

---

## 1. The hypothesis

**H1 (primary, confirmatory, one-sided).** Among bidder entries where a bid could have
been priced on either arm, the probability that an entry carries a price is HIGHER on the
`ru-mobile` arm than on the foreign arm.

**H0.** The arm label is exchangeable with respect to whether an entry carries a price.

One-sided, and the direction is fixed *here*, by the pilot. A two-sided test would be
more conservative but would misdescribe the claim: nobody hypothesises that a Russian
publisher's bidders price *less* for a Russian visitor. If the effect reverses, H1 is not
rescued by a two-sided p — it is refuted, and the refutation is the finding.

## 2. The analysis set — stated before the data, because this is where a study is won or lost

A bidder entry enters the confirmatory analysis if and only if:

1. Its load is on a **readable pair**: both arms of the pair produced an HTTP response.
   A pair whose half was never seen carries no RU-versus-foreign contrast at all.
2. Both arms of that pair have `vantage_verified` true. **An arm that could not prove
   where it stood renders `blocked`, and `blocked` is NEVER folded into `no-bid`.** This
   is the single error this project cannot afford to make quietly: the RU arm rides a
   Note3 tether, which is fragile and is the only one we have, and a fortnight of runs
   that silently recorded our own blindness as the market's silence would produce a
   confident, publishable, false result. Encoded, not merely intended:
   `unseen_cause()` in `tools/adint-status` splits refused / no-answer and never emits a
   bid outcome; `tools/adint-paired-run` refuses a load whose echo cannot be read.
3. Its (bidder, site) combination produced **at least one non-empty auction on BOTH
   arms** within the window. This is the confound control the pilot already needed: six
   bidders looked RU-only until it turned out the foreign arm produced zero non-empty
   auctions on all three of their sites, and absence of the site's auction is not absence
   of the bidder. Without this filter the primary test measures site coverage, not
   pricing behaviour.

Anything failing 1–3 is **excluded from the confirmatory test and reported as an
exclusion count**, per arm and per reason. An exclusion count is a bias term, not a
footnote.

## 3. The test

**Primary: a paired permutation test on the arm label, 100 000 permutations, seed fixed
at registration (`--seed 20260821`).**

The statistic is the difference in priced-entry rate between arms, pooled over the
analysis set. Under H0 the arm label is exchangeable **within a pair**, so permutation
happens within pairs and never across them. This is the test that matches the design:
the pairing exists precisely to hold site, hour, frame position and network moment fixed,
and permuting within it preserves every one of those while destroying only the thing H1 is
about. It also survives the clustering that Fisher does not — entries are nested in loads,
loads in sites, and treating 1 259 entries as independent draws overstates the evidence.

**Secondary, reported alongside and NOT as the headline: one-sided Fisher exact on the
2×2 of (arm × priced), over the same analysis set.** It is here only for continuity with
the pilot's published p = 3.5 × 10⁻⁵ and because a reader will compute it anyway. Where
the two disagree, **the permutation test is the result** — it is the one that respects
the design, and this sentence exists so that preference cannot be chosen later.

**α = 0.01, one-sided.** Tighter than convention on purpose: the pilot p is already far
below 0.05, so 0.05 would be a threshold nothing could fail, and a threshold nothing can
fail is not a test.

**Effect size, reported whatever the p:** the priced-entry rate per arm with exact
binomial CIs, and the rate difference with its permutation CI. A p-value with no effect
size is a claim with no magnitude.

## 4. The stopping rule

**Fixed horizon. The window is 14 days of collection beginning with the first cell
captured after this file is committed, and it ends 14×24 h later regardless of what the
data looks like.**

- **No interim confirmatory analysis.** The confirmatory statistic is computed **once**,
  after the window closes. `tools/adint-prereg --result` refuses to compute it before
  then and says how long remains.
- **No early stop for significance, and no extension for its absence.** Both are the same
  error: letting the data decide when to stop looking is how a null becomes a finding.
- **Insufficient-data stop.** If at window close the analysis set holds fewer than **30
  priced entries across both arms**, the confirmatory test is reported as
  **underpowered** with its actual n, and H1 is neither accepted nor rejected. Declared
  now so that a thin result cannot be quietly upgraded to "no effect" or padded by
  running longer.
- **The collection cadence is not mine to forget.** It is wired as a reflex with its own
  unconditional schedule, not driven by an agent's loop — a step everything depends on
  must have its own cadence.
- **A night that collects nothing is a night of zero rows with a reason, never a clean
  run.** Coverage — sites answering, arms alive, loads refused and why — is published per
  cell. An empty set writes itself as a success unless something makes it write itself as
  an empty set.

## 5. Subgroups — declared here, exploratory by construction

Named in advance so that a subgroup that "worked" cannot be presented as if it had been
the plan. **Every one of these is EXPLORATORY.** None can establish H1, none is reported
as a confirmatory result, and all carry Holm-corrected p-values across the family below.

1. **Daypart** (MSK night / morning / day / evening). The pilot is entirely night. If
   the effect is real, its size across dayparts is the most interesting number in the
   study; if it exists only at night, that is a finding about our sampling.
2. **Weekday vs weekend.**
3. **Per bidder** — `roxot` and `adwile` separately, plus any bidder reaching ≥30 entries
   on both arms within the window.
4. **Per site.**
5. **Foreign arm identity** (`nl-direct` vs any other foreign exit run in the window).
   The pilot's foreign arm was `us-exit`; the current one is `nl-direct`. If the effect
   is about "not Russia" both behave alike; if it tracks a specific country or ASN it is
   a different effect with a different name.

New subgroups may be added later **only** as clearly-labelled post-hoc exploration, never
into this list.

## 6. What would refute H1

Written down because a hypothesis that cannot say what would refute it is not one.

- Priced entries appearing on the foreign arm at a comparable rate once coverage is
  matched by §2.3.
- The pilot's eleven prices proving to be concentrated in a single night or a single
  site, with the rest of the window flat — an artifact of one moment.
- The effect vanishing when the vantage-verified filter is applied strictly, which would
  mean it was our own blindness all along.

## 7. Data handling — the two rules that make numbers reproducible

- **Append-only ledgers, kept whole for the full window.** No sliding window, no pruning.
  This mesh has a sound-studio ledger that prunes every sweep, and the consequence is that
  no `n` quoted against it is reproducible and n moves DOWN as well as up. This corpus
  must not acquire that property. Space is not the constraint: 721 G free, `data/` is
  432 M.
- **Every cell published with its coverage**, through `tools/adint-publish`'s field
  allowlist, and the README numbers block regenerated by `tools/adint-status --write`.

---

*Amendments to this file after the window opens must be recorded as amendments, with
their date and reason, and must never silently alter §1, §3 or §4. `--verify` will
report a hash mismatch, which is the point.*
