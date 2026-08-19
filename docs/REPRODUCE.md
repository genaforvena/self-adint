# How to re-run this, and how to check that we did not fool ourselves

You need this repository and a machine. You do not need our notes, our node, or our word
for anything. Where a step needs hardware we have and you may not, it says so and tells you
what the result would look like without it.

Everything below was run on Ubuntu 24.04, Python 3.12, Playwright Chromium.

---

## 0. Setup

```bash
git clone https://github.com/genaforvena/self-adint && cd self-adint
python3 -m venv ~/.venv-browser
~/.venv-browser/bin/pip install playwright && ~/.venv-browser/bin/playwright install chromium
```

Two interpreters are used on purpose. Tools that only compute run under the system
`python3`; tools that drive a browser need the venv. Each tool says which in its `--test`
banner, and a browser tool run under the wrong interpreter exits **n/a**, not FAIL — a
missing dependency is not a failed assertion.

## 1. Run the gates first

Every tool self-tests offline against a local fixture. Run them before trusting any output;
they take about three minutes in total.

```bash
for t in tools/adint-*; do
  echo "== $t"; python3 "$t" --test 2>&1 | tail -1
done
# the browser-driving ones:
~/.venv-browser/bin/python3 tools/adint-hb-capture --test   | tail -1
~/.venv-browser/bin/python3 tools/adint-pbjs-probe  --test  | tail -1
```

Exit codes: `0` pass · `1` a real failure · `2` **n/a** — the node lacks what the leg needs
(no browser, no RU vantage). An n/a is an honest pass, and it names its reason.

**What the gates actually assert.** A test that only proves the happy path is decoration, so
the load-bearing ones are listed with the mistake each exists to catch:

| gate | it fails if… | the mistake it caught |
|---|---|---|
| `adint-frame-stageb --test` | a 4xx is classified as "no header bidding" | a rejection reason becoming a market fact |
| " | bare `window.Ya` admits a domain | Yandex.Metrika (on much of runet) reading as ad-tech |
| " | a wrapper found on a 403 error page admits | classifying a page we never saw |
| `adint-hb-capture --test` | `no-answer` folds into `no-bid` | our own latency measured as bidder selectivity |
| " | both arms of a pair get different keys | pairs reconstructed by timestamp arithmetic |
| " | a `direct` arm's exit is read from the shell | a browser leaving by a path the shell does not use |
| `adint-pbjs-probe --test` | a subframe auction reads as "no prebid" | a claim about one frame printed as a claim about a site |
| `adint-publish --test` | a field nobody anticipated crosses to `public-data/` | a denylist promising something about columns that do not exist yet |

Each was broken deliberately and watched go red before it was accepted. A gate you have not
seen fail is not a gate.

## 2. Build the sampling frame (stage A → stage B)

Stage A is committed: `ref/ru-universe-tranco-V3JQN.tsv`, every registrable `.ru` / `.su` /
`.рф` domain in the Tranco top 50 000, **n = 1944**. Re-fetch and diff it yourself:

```bash
curl -sL https://tranco-list.eu/download/V3JQN/1000000 -o tranco.csv
sha256sum tranco.csv   # expect c68e363d5cc412a544024d947826db636fe84e56e16d95a3ae32f82684dcdd2a
```

Stage B walks that universe in rank order and admits a domain only if its page answers 200
with HTML **and** carries an ad wrapper. It publishes every domain it touches, including the
rejections and their reasons:

```bash
~/.venv-browser/bin/python3 tools/adint-frame-stageb \
    --vantage nl-direct --admit-max 20 --limit 120 --settle 8 --recheck-settle 45
```

- `--settle 8` is the frame probe's window. It is **not** the study's `window_s = 45`: the
  frame asks whether a wrapper is PRESENT, the capture asks whether an auction COMPLETED.
- `--recheck-settle 45` re-probes every `no-wrapper` verdict at the full window before writing
  it down, so a rejection is never cheaper than the acceptance threshold of the study itself.
- The run is resumable: re-running skips domains already in the ledger.

Then publish the ledger through the allowlist (never copy it by hand):

```bash
python3 tools/adint-publish --kind frame data/frame-stageb-nl-direct-<date>.jsonl
```

### A rejection ages differently from an admission

`admit` stands — we saw the wrapper. `no-wrapper` only ever means *we looked with this
detector and found nothing*, so when the detector learns a new signal, every earlier
`no-wrapper` silently changes meaning without changing value. Go back and look:

```bash
~/.venv-browser/bin/python3 tools/adint-frame-stageb --vantage nl-direct \
    --audit-no-wrapper data/frame-stageb-nl-direct-<date>.jsonl
```

## 3. Capture one cell of the schedule, paired

```bash
~/.venv-browser/bin/python3 tools/adint-paired-run \
    --frame ref/ru-frame-stageB-<vantage>-<date>.tsv --replicates 1 --window 45
```

Both arms are launched on **one site at a time** and the next site waits for both. The fast
arm idles; that idle is the price of the pair being a pair. Two independent walks cannot do
this — the arms differ in channel, so they drift apart monotonically and by the tail of a
20-site list the "pair" is an hour apart, with both timestamps perfectly honest.

Site order is randomised from a **recorded seed** and both arms get the identical order,
because the pair key is positional: a differing order would not mismatch loudly, it would
quietly pair different sites.

### Without the RU vantage

The RU arm runs inside a network namespace fed by a rooted Android phone's USB tether on a
Russian SIM (`docs/step0b-ru-vantage-note3-2026-08-18.md` has the full procedure). Without
one, run the NL arm alone:

```bash
~/.venv-browser/bin/python3 tools/adint-hb-capture --file sites.txt \
    --settle 45 --via direct --arm nl-direct
```

You will reproduce the NL column and not the comparison. That is a smaller result, not a
broken one — and it is worth knowing that some of what we record as `blocked` may be
geo-blocking that a Russian vantage would not see at all.

## 4. Read the capture

```bash
python3 tools/adint-hb-report --log data/hb-paired-<run>.nl.jsonl
```

The report refuses several things on purpose: it never pools schemas, never folds `no-answer`
into `no-bid`, counts **sites** rather than rows, and prints a load that ran a wrapper but
produced no candidate request as a **coverage gap** rather than as absent demand. Those two
readings point in opposite directions.

---

## What you cannot reproduce from here, and why

- **The server-side auction.** Everything here is a passive read of what a wrapper hands the
  browser. Server-to-server bidding is invisible from this vantage, full stop. A domain we
  call `no-wrapper` may be selling its inventory perfectly well.
- **Our exact vantages.** `77.246.104.228` (Amsterdam) and a rotating MTS mobile address in
  Moscow are ours. Yours will differ, and on a market where access is geo-shaped that changes
  what you see. This is why every row carries the egress the **browser** read — not the one
  the shell measured.
- **The clock.** Zones are Moscow time. A cell run at a different hour is a different cell,
  and the zone label is stored per run rather than inferred afterwards.
