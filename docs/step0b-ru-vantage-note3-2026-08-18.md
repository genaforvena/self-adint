# The RU-resident mobile vantage: what it took, what it cost the instrument, and what it answered

**Task**: board `adint-ru-vantage-via-note3` (owner `mesh-note3/adint`, `priority:incident`),
2026-08-18. **Question**: the 08-16 study closed with "the central number of this sweep — zero
observed auctions — is a statement about the vantage, not about the publishers", and named two
compounding defects: wrong geo, and a `hosting: true` datacenter ASN that invalid-traffic filters
exist to drop. This run removes both and re-asks.

**Gate the operator set, and it is the one that matters**: the vantage must be measured *from the
phone's side*, not from a shell on mesh-home — because on 2026-08-16 this lane recorded privoxy's
Montréal exit as the node's own. See "The artifact" below; nothing else in this document counts
until that one does.

---

## The artifact

Read by the phone, with `busybox wget` on the Note3 itself, against a numeric IP so no resolver
is involved, at a moment when the phone held only `lo` and `rmnet0` — wifi off, no other path
existed:

```
{"country":"Russia","regionName":"Moscow","city":"Moscow","isp":"MTS PJSC",
 "org":"PJSC \"MTS\" MBN","as":"AS8359 MTS PJSC","mobile":true,"proxy":false,
 "hosting":false,"query":"91.79.82.x"}
```

A second, independent echo (`ifconfig.co`, numeric) returned the same address. Full artifact with
the interface list, route table and radio props: `data/ru-vantage-note3-2026-08-18.txt`.

Control, measured in the same minutes on the node's own default route: **`77.246.104.228`,
Amsterdam, AS216071, `hosting: true`**. Two arms, two exits, one run — the thing the preflight
in `adint-hb-capture` exists to refuse without.

The carrier address is not stable: it moved to `91.79.90.x` partway through (re-checked:
still Moscow, AS8359, `mobile:true`, `hosting:false`). Every load records the exit its own
browser saw, so this costs nothing — but a document that quoted one IP as "the vantage" would
have been wrong within the hour.

## How the vantage is built

Full procedure, including the two steps that are not obvious, is in the operator memory
`ru-vantage-via-note3-tether-netns`. The short form:

1. `su -c "svc wifi disable"` on the phone. **As root** — plain `svc wifi disable` dies with a
   Samsung `sh: resetreason: can't execute: Permission denied` and silently does nothing. Wifi
   off is what makes Android raise the PDP context.
2. `setprop sys.usb.config rndis,adb`, then configure `rndis0` and NAT **by hand** — Android's
   tethering service is not involved, so no UI tap is needed (the older note required one).
3. **Android routes its own tether subnet out the radio.** netd's rule 22000 sends every lookup
   to the `rmnet0` table before the uid-0 → main rule, so `ip route get 192.168.42.130` answers
   `via 10.205.173.x dev rmnet0` and the host's ARP is never answered. Fix: a `pref 100` rule
   for `192.168.42.0/24` into a table holding the connected route.
4. Host side: **move the tether device out of the root namespace entirely.** This is not
   fastidiousness — **Chromium cannot bind a socket to an interface**, so `mesh-uplink-usb`'s
   scoped `oif` rule can never catch it. In a netns every process inside is on the tether by
   construction and there is no leak left to police. `main`, the wifi dongle and tailscale are
   untouched, and the rollback (`ip netns del`) is not on the path this node is reached over.

## Two instrument defects the vantage found, and the gate seen red for each

Both were in `adint-hb-capture`, both fixed, both mutants run.

**1. The `--via direct` gate compared two identical measurements.** It refuses when the browser's
exit equals `egress_ip_proxied` — but that value came from a curl run with the *same* environment
as the direct one whenever no proxy variable is set. Same command, same env: the comparison is a
tautology, and it fired, refusing a run that had no proxy to leak through. (It fired here because
this pane's shell carries no proxy vars at all — `restore.env`'s three `export` lines are
commented out, so the "shells here are proxied" rule is a property of a pane's age, not of the
node.) Absent a proxy env there is nothing for the browser to inherit, so the honest value is
`None`. Mutant: proxy left in the browser's launch env → `refusing: --via direct still leaves
from 38.49.216.141`. Fixed tool, same environment → passes at `77.246.104.228`.

**2. The IP echo was a hardcoded host with a fibre-tuned timeout, and the vantage broke it.**
MTS's own resolvers answer `api.ipify.org` **and** `example.com` with `8.47.69.0` / `8.6.112.0` —
the same /24s Google answers with, last octet zeroed — and nothing listens there. Every `.ru`
host and both Cloudflare hosts resolve byte-identically, so this is narrow, not a block. The
effect is that the one host the gate had to reach was the one host the vantage could not, and it
presents as `ERR_NAME_NOT_RESOLVED` — indistinguishable from the HSDPA link wedging, which is
what it was first read as.

The fix is to move the *instrument*, never the subscriber's resolver: `--ip-echo` (default
unchanged) and `--echo-timeout` (patience scales with `--nav-timeout`; the gate does not move —
an unread echo still refuses). Mutant: an unroutable echo at a 30s timeout still refuses.
Caught in review before it shipped: `echo=IP_ECHO` as a **default argument** binds once at
def time and would have made `--ip-echo` silently inert.

## The confound this treatment carries, stated before any count

Changing the vantage changed the **link**: fibre → HSDPA. Exactly one outcome column moves, and
its mirror moves the opposite way (rendered loads only, per load):

| outcome | Amsterdam /load | RU MTS /load |
|---|---|---|
| no-answer | 13.4 | **50.7** |
| no-bid | 33.7 | **12.8** |
| total candidate requests | 94.6 | 107.8 |
| unread | 19.0 | 19.7 |
| binary | 18.7 | 16.6 |

`no-answer` is defined as "no response arrived inside OUR observation window" — a claim about the
capture, not about the bidder. The same auction is *asked* (total requests barely moves); its
answers migrate out of `no-bid` and into `no-answer`. **The confound points the same way as the
hypothesis**, which is the worst case, so:

- the comparison is made on `priced` — evidence *of* a bid — and never on `no-answer`, which is
  not evidence of a bid's absence;
- the window was widened from 14s to 45s and the effect was **measured, not argued**: at
  `--settle 14` all 21 of sparrow's requests were `no-answer`; at `--settle 45` sparrow answered
  on e1.ru with a real price (0.18265177380087824 RUB) and one zero bid. The window was the
  silence.

A correction that belongs here because it was published first and was wrong: `unread` was
initially named as a second link effect. It is 19.0 vs 19.7 — flat — and 364 of 372 such rows are
`mc.yandex.ru` Metrica beacons, not auction messages. The baseline carries 1588 of the same rows.
The error was quoting per-load counts from the largest HSDPA loads against an *impression* of the
baseline instead of computing it, over a file already on disk.

> **Redaction note.** This repo is public. The node's own exit (`77.246.104.228`, a rented
> datacenter host) is already published here and stays in full. The **mobile subscriber**
> addresses are truncated to their first three octets, the way this repo truncates advertising
> IDs — the ASN, the geo and the `mobile`/`hosting` flags carry the entire claim, and the last
> octet adds nothing to it. The full values live in `data/`, which is gitignored.

---

## The result

**One participant that had never returned a bid to this study returns one from the RU mobile
vantage, and it is not the one the baseline turned on.** The baseline sentence was *"32 bidders
answered at least one auction message; exactly one ever returned a bid object (sparrow)"*. From
the RU vantage the count of answering participants does **not** grow — 32 against 35, on half the
loads — and the new bid comes from `ssp.24smi.net`, which the Amsterdam corpus records as
declining every single time it was asked.

The claim is made on kp.ru, the site that asks that host, and it is made against a **same-day,
same-window control run on this node's own route** rather than against the two-day-old baseline
alone. That control exists because the first version of this comparison was confounded three ways
at once — geo, calendar day, and observation window — and any one of them could have produced the
whole effect.

| kp.ru, rows from `ssp.24smi.net` | NL 08-16 (settle 8–14) | **NL 08-18 control (settle 45)** | RU MTS 08-18 |
|---|---|---|---|
| rendered loads | 4 | **14** | 16 |
| loads carrying a price | 0 | **0** | **12** |
| `priced` rows | 0 | **0** | **12** |
| `no-bid` rows | 8 | **27** | 16 |
| `no-answer` / `unread` | 0 | 1 / 0 | 3 / 1 |
| exit, read inside the browser | 77.246.104.228 ×4 | 77.246.104.228 ×14 | 91.79.82.x ×1, 91.79.90.x ×15 |

Fisher exact, one-sided, against the matched control — same day, same hour band, same 45s window,
same browser config (`ru-RU`, `Europe/Moscow`), the only difference being the path:

- **by load: 12/16 vs 0/14 → p = 2.1 × 10⁻⁵**
- by answered row: 12/28 vs 0/27 → p = 6.9 × 10⁻⁵
- pooling both NL arms, by load: 12/16 vs 0/18 → p = 3.3 × 10⁻⁶

**The control is not vacuous and that was checked before it was trusted.** `ssp.24smi.net` is
asked in the control arm — 27 `no-bid` rows over 14 loads — so the zero is a record of declining,
not of never being asked. A control whose request never fires would have produced the same zero
and meant nothing.

**The prices.** Twelve observations, two distinct values: `2.206303398` and `2.22496865604` RUB,
both at placement `28986`, each echoed a second time inside the returned `displayCode` as
`yhbPrice`. Exact repeats within a session, a small move between sessions — the same shape the
08-16 sweep found in sparrow's numbers, and neither value is on the 0.5 grid the round atoms sit
on. As always in this protocol there is no `floorData`, so a fixed price and a floor-follower
remain inseparable.

**The host stays unattributed.** `adint-attribute-by-adapter` finds no Prebid adapter that cites
`ssp.24smi.net` in its bytes. kp.ru's roster does carry `smi2` and `24smi.net` is SMI2's domain,
but that is a brand inference of exactly the kind this repo's attribution rule exists to refuse
(`r.ussp.io` → `anzuSSP` is the standing example). The finding is filed against the HOST.

### What this does and does not establish

It establishes that **the vantage was a live cause of at least one bidder's silence**, not merely
a suspected one — the 08-16 conclusion ("a statement about the vantage, not about the publishers")
now has a positive instance behind it rather than only an argument.

It does **not** identify which property of the vantage did it, and this is the honest limit: the
path changed in two ways at once. `91.79.x` is Russian **and** a mobile carrier ASN with
`hosting: false`; `77.246.104.228` is Dutch **and** `hosting: true`. The 08-16 write-up named both
as suspects and they moved together here, so this run cannot apportion the effect between them.
Separating them needs a third vantage: a **RU datacenter** IP (geo held, `hosting` flipped) or a
**non-RU mobile** IP (geo flipped, `hosting` held). Neither exists on this mesh today.

Three further limits, each of which would have been easy to leave unsaid:

- **The effect is one host on one site.** 25 other named bidders were asked from the RU vantage
  and none of them started bidding. "The geography was the problem" is not what this shows; "the
  geography was *a* problem for *one* participant" is.
- **The arms are sequential, not interleaved.** Both ran on 2026-08-18 within the same hour band,
  but the RU block preceded the NL block. Interleaving per load would close the last of the
  temporal gap, and the tool's `--treatment`/`--arms` machinery cannot currently do it across a
  network namespace.
- **The count of answering participants did not rise** (32 vs 35 on half the loads). If the
  expectation going in was "a RU residential vantage unlocks the field", the answer to that is
  **no**, and it is an answer, not a failed run. What it unlocked was one participant's price.

### Reproducing it

```bash
# vantage up (see the memory note for the three non-obvious steps)
sudo ip netns exec ruvantage runuser -u mesh-home -- \
  env -u HTTP_PROXY -u HTTPS_PROXY -u ALL_PROXY HOME=/home/mesh-home \
  ~/.venv-browser/bin/python3 tools/adint-hb-capture --file <kp.ru x14> \
    --via direct --arm ru-mts --settle 45 --nav-timeout 150 \
    --echo-timeout 60 --ip-echo https://icanhazip.com --out data/hb-auction-ru.jsonl

# the control, same window, this node's own route
env -u HTTP_PROXY -u HTTPS_PROXY -u ALL_PROXY ~/.venv-browser/bin/python3 \
  tools/adint-hb-capture --file <kp.ru x14> --via direct --arm nl-control \
    --settle 45 --nav-timeout 150 --echo-timeout 60 --out data/hb-auction-nlctl.jsonl

tools/adint-hb-report --log data/hb-auction-ru.jsonl
tools/adint-hb-report --log data/hb-auction-nlctl.jsonl
```

The logs are gitignored (`data/`), as is every capture in this repo; the counts above are
recomputed from them by the reporter, never quoted from a console tail.
