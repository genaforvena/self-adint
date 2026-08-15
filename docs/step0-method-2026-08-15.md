# Step 0 — method record (adint window, 2026-08-15)

**Question:** which exchanges actually sell the operator's own phone's traffic?
**Artifact:** `data/bundle-exchange.tsv` — bundle → org, with observation counts. Not "app X shows ads".

This file records what was *measured from this node*, what was *rejected and why*, and what
therefore has to be asked of the operator. It is written before the instrument exists so that
the reasoning is auditable when the numbers arrive.

## What is reachable from `mesh-home`, measured 2026-08-15 05:5x UTC

| target | probe | result |
|---|---|---|
| phone (Redmi 10) LAN `<phone-lan-ip>` | ICMP | up |
| phone `:8022` (Termux sshd) | TCP connect | **open** — sshd alive |
| phone ssh `<termux-user>@…:8022` | ssh, our key | **Permission denied (publickey,password,keyboard-interactive)** |
| phone adb `<phone-lan-ip>:<adb-port>` | `adb connect` | connection refused |
| phone tailscale `<phone-tailscale-ip>` | tailscale status | offline, last seen 49d |
| phone LAN alt `<phone-lan-ip-stale>` | ICMP | no reply (this is what `mesh-phone-ip` still reports) |
| router `<router-lan-ip>` | TCP 22/53/80/443 | all open |
| router ssh `root@<router-lan-ip>` | ssh, our key | Permission denied (publickey,password) |

So: **we have no shell on his phone and no credentials on the router.** `mesh-phone-ip` reports
`phone UP, key REJECTED (cached auth-rejected)`, and `mesh-phone-watch` has been logging
`PRESENT on LAN but ADB transport AND sshd both down` — the second half of that line is wrong,
sshd answers; what is down is *our authorisation*, not the daemon. (Genome-side observation, not
this project's lane — reported to the board as `[fyi]`, not fixed here.)

Note also: this node's LAN address has drifted by one from what `CLAUDE.local.md` records.
Same class of drift; likewise reported, not fixed here.

## Options considered, and the ones rejected

1. **Passive Wi-Fi capture from this node (monitor mode + the AP's PSK).** *Rejected on rule 4.*
   Monitor mode captures the whole cell; every other device in the flat lands in the capture by
   construction. "Nobody but him is in scope, ever, including incidentally" — a design step that
   needs foreign frames on the wire is the wrong step, and a post-hoc filter does not undo it.
2. **ARP/DHCP-level MITM on the LAN.** *Rejected* for the same reason plus it is a substrate
   change on a contended commons.
3. **Router DNS / flow log.** *Instruction-only, and secondary.* The router is read-only on this
   mesh, and we hold no credentials, so this needs him either way. It also gives **no app
   attribution**: a DNS log proves the household resolved `pubads.g.doubleclick.net`, not which
   package asked, and the artifact Step 0 demands is keyed by bundle. Useful later as a
   cross-check on *volume*, not as the primary instrument.
4. **On-device per-app capture (VPNService-based, no root), run by him.** The only channel that
   yields the required key — `(package, hostname)` — at zero cost and touching nothing but his
   own device. This is the primary instrument. Instrument survey and the step-by-step:
   `research-device-capture-2026-08-15.md`.
5. **Static SDK inventory from a Termux shell** (`pm list packages -f`, read the APKs, look for
   embedded ad-SDK signatures). Cheap, *complementary*, and honest about its limit: it yields
   **candidates** — which SDKs an app embeds — never observation counts. It cannot answer "which
   exchange actually sold an impression", only "which could". Needs the ssh key re-authorised.

## What that means for the operator

Two asks, both on his own device, both reversible, neither irreversible in the project's sense
(no seat, no contract, no deposit, no GAID reset):

- **A.** Re-authorise our ssh key in Termux (one command, restores the mesh's own phone body too).
- **B.** Install and run the capture app for a normal day of his usage, then export.

And one **question that has to be answered before either is worth doing**: *is the Redmi 10 the
phone he actually carries and uses* — the one with his real installed apps and his real ad
profile — or is it a mesh sense body with almost no consumer apps on it? Step 0 measures the
device he lives on; measuring the wrong handset produces a table that is honest and useless.
Per rule 6 this is not reconstructed from the mesh's config: it is asked.

## The pipeline that is already built and tested

`tools/adint-aggregate` (self-test: `python3 tools/adint-aggregate --test`) turns normalised
observation records into the artifact:

    {"ts":…,"source":…,"device":…,"package":…,"hostname":…,"count":…}   (JSONL)
      → data/bundle-exchange.tsv     bundle → org/role/rtb_exchange/external_seat + counts
      → data/unresolved-hosts.tsv    every host the reference table could not name

Properties asserted by the self-test, **each seen red under a deliberate mutation** before being
accepted (2026-08-15):

| property | mutation that must break it | seen red |
|---|---|---|
| a foreign device's rows never reach any output | `if dev != device:` → `if False:` | yes — 7 checks red |
| longest-suffix domain match, not first-match | `len(pat) > len(best[0])` → first wins | yes |
| an unnamed host is loud, with its count | drop the unresolved record / zero its count | yes (both) |
| a missing reference table is an error, not an empty map | `raise` → `return {}` | yes |

The last one matters more than it looks: an empty mapping renders *every* host `UNKNOWN` and the
run still "succeeds" — a table of nothing that looks like a finished measurement.

The reference table itself (`ref/exchange-domains.tsv`, domain → org → does that org run an
OpenRTB exchange → can an outsider get a seat there) is public knowledge, committed, and is what
turns hostnames into the decision Step 0 exists to make: **which exchange to approach at all.**

## The static half, run (2026-08-15)

He confirmed the handset is his daily phone, raised `sshd` on it himself, and it came back on a
different LAN address than the one the mesh had cached. With a shell:
`tools/adint-inventory-run` streamed every third-party APK's DEX through **one** `grep` pass on
the phone — nothing copied off the device, nothing written on it — and
`tools/adint-candidates` aggregated the result.

**39 of his 66 third-party packages carry at least one ad-SDK signature.** Two rankings, both
published, because they disagree:

| axis | order |
|---|---|
| by app count | Google 26 · AppLovin 12 · ironSource 11 · DigitalTurbine 10 · Huawei 9 · Yandex 7 |
| by hit count | Yandex 14658 · InMobi 5577 · Google 5084 · Mintegral 3974 · ironSource 3369 · ByteDance 2627 |

The disagreement is the finding, not noise in it. AppLovin sits in 12 apps at ~10 hits each —
that is a *mediation adapter naming a demand source*, not AppLovin's SDK; Yandex sits in 7 at
~2100 hits each, which is an embedded SDK. Rank on one axis alone and you pick a winner the
other axis contradicts. The split threshold between "adapter reference" and "embedded SDK" is
**not fixed here**: it is a property of this corpus and gets derived from it when the corpus is
worth deriving from.

One signature had to be refined against the real data: `com/google/android/gms/ads/identifier`
is its own row, role `identifier`, demand-side **no**. Eight of the 34 apps matching the bare
`com/google/android/gms/ads` prefix only read the advertising id — counting them as demand-side
inflated Google by a third. That split relies on `grep -oE` returning the leftmost-**longest**
alternative, which was verified *on the device* (GNU grep 3.11: 9 + 6, not 15 + 6) before
anything was built on it.

### A sharper static probe, and the trap in it

The signature scan asks *which SDK families ship*. A second, sharper question can be asked of the
same DEX for free: **which hostnames are written into the app as literals.** It settled a real
open question in minutes — an app carrying `com/tapjoy` ×20 turned out to contain **no Tapjoy host
of either era** (neither `connect.tapjoy.com` from ≤14.1.1 nor `gateway-b.offerwall.unity3d.com`
from ≥14.2.0), only listener class names: that app references Tapjoy through an adapter, it does
not ship the SDK. The control ran first — 63 `https://` literals and a dozen real ad hosts in the
same file — because an empty result from a broken probe and an empty result from a clean app look
identical.

**The trap:** DEX strings are length-prefixed MUTF-8 with no separator, so a permissive discovery
regex swallows the preceding bytes and *mints hostnames that do not exist* — this scan produced
`5com.applovin.com`, `.com.applovin.com`, `9wv.inner-active.mobi`. Harmless when matching a fixed
list of known hosts (the pattern anchors both ends), fatal for discovery: those strings would have
entered a reference table as new domains, sourced to "observed on the device". Any literal-mining
lane must anchor on a known suffix and treat the left edge as garbage until proven otherwise.

### Two static lanes with opposite blind spots

The literal probe was then run across the whole phone against all 87 reference domains
(`tools/adint-inventory-run --key hostname`), and cross-checked against the signature scan
(`tools/adint-crosscheck`). The disagreement is the point:

| org | apps with the SDK signature | apps with a host literal |
|---|---|---|
| Google | 18 | 12 |
| AppLovin | 12 | 4 |
| Unity/ironSource | 11 | 1 |
| Digital Turbine | 10 | 1 |
| Huawei | 9 | 0 |
| Yandex | 7 | 6 |
| Meta | 6 | 0 |
| MyTarget/VK | **3** | **5** |

**Signatures without literals** (Huawei, Meta, Chartboost, Mintegral, Pangle, Smaato, Amazon,
Appodeal, Tapjoy) are the expected case: the endpoint is assembled at runtime, or the reference
is a mediation adapter rather than the SDK. Not evidence of absence either way.

**Literals without signatures is the finding.** Three apps — a bank, a mobile operator, and one
more — carry myTarget/VK ad hosts (`r.my.com`, `ad.mail.ru`) with no `com/my/target` class path
anywhere in their DEX. The reason is mechanical: **R8/ProGuard renames class paths and does not
touch string constants.** A signature scan goes blind on an obfuscated app in exactly the
direction that matters — it reports *no demand relationship* where one is written into the code
— and the literal scan sees straight through it. The reverse blind spot is just as real: an SDK
that builds its host at runtime has no literal at all.

Neither lane is an observation, and this is the whole reason the cross-check exists rather than
a single "static scan" that would have silently inherited one of the two blindnesses. It also
gives the capture a falsifiable prediction: those three apps should contact `my.com` hosts on the
wire, and if they do not, one of the two static lanes is lying about something.

## The shape of Step 0's own verdict: `IN` / `UNKNOWN`, never `NOT`

The oracle his brief designed for segments — `LOSS` is not `NOT IN`, a negative is
indistinguishable from an absent result — turns out to govern **Step 0's own instrument** for a
different reason, and this is the finding that outlives every number here.

A device-side capture sees the connections his phone makes. It does **not** see what happens on
the other side of them. An impression can be transacted entirely server-to-server: the SDK talks
to its own host, and the auction that sells him runs between that host and an exchange his phone
never contacts. Two demand paths are invisible even on the near side — Smaato configures over DNS
with no TLS connection at all, and Meta multiplexes ad requests through `graph.facebook.com`,
separable only by URL path and never from a router flow log.

Therefore:

    exchange observed on the wire       → IN       (it carries his traffic)
    exchange not observed               → UNKNOWN  (always)
    exchange does not carry his traffic → NOT REACHABLE from this instrument, ever

An exchange missing from the table is a **statement about the instrument**, not about the
exchange. Ranking by observation count is legitimate; reading a zero as "this exchange does not
sell him" is the fabricated negative the whole project is built to avoid, and it would be
fabricated at the exact step whose job is to stop an irreversible decision from resting on a
guess.

**What this half cannot say, and must never be read as saying:** a signature in a DEX proves the
code ships. It does not prove an impression was ever sold through that org, it does not say which
exchange won an auction, and mediation SDKs deliberately name demand sources they do not contain.
That is why the candidates live in `data/bundle-sdk-candidates.tsv` and not in
`data/bundle-exchange.tsv`, and why Step 0 is not finished: **the counts that decide which
exchange to approach have to come off the wire, per app.**
