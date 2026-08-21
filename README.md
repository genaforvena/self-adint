# self-adint — reading your own advertising profile from the demand side

**The question.** When an ad is auctioned for your phone, the bid request that goes out
describes *you* — segments, identity graphs, consent state, location precision. Buyers see
that description for free, before anyone pays for anything. So: **what does the demand side
already know about one specific person's device, and can that person read it himself?**

The thesis is not "what would my impressions cost". It is that *the audience profile arrives
inside the bid request*. You do not have to buy the impression to read what came with it.

One operator, one phone, his own data. Nobody else is observed — not incidentally, not as a
"noise statistic". That constraint shapes the code, not just the intentions: see *The rules
the code is built to*, below.

---

## Where it stands, with the numbers that qualify it

Two threads run in parallel. **`docs/INDEX.md` is the map**; start there if you want to
navigate rather than read.

**Thread A — measurement.** Which exchanges actually sell this device's traffic, and what is
observable from where. This is where the current work is.

**Thread B — access.** What a person can legally demand from an exchange, and what it takes
for a small buyer to get a seat.

### The current measurement, stated with its size before its content

The live work is a study of **Russian header bidding**, chosen because the western literature
on the subject covers US/EU inventory and does not contain the Russian SSPs at all — `Yandex`,
`adriver`, `buzzoola`, `betweendigital`, `sape`, `smi2`, `MyTarget/VK`, `astralab`, `gnezdo`,
`hybrid`, `bidvol`, `alfasense` are visible by name in our own capture and absent from the
papers.

<!-- NUMBERS: regenerate with `python3 tools/adint-status` — do not hand-edit -->

*Generated 2026-08-21 03:28 UTC by `tools/adint-status` from the files on disk.*

### The sampling frame — and its rejections, which are half the result

The study runs on the **`ru-mobile`** frame, declared in `ref/CANONICAL-FRAME`. Which vantage a frame is built from changes its membership, so this is a choice with a reason, recorded in that file.

**Vantage `ru-mobile`** — **canonical** · 208 domains touched in rank order · 3 duplicate row(s) collapsed · browser egress 91.78.80.171, 91.79.81.62 · `frame-stageb-ru-mobile-2026-08-19-schema3.jsonl`

| verdict | n | what it means |
|---|---:|---|
| `admit` | 14 | **in the frame** — 200, HTML, and an AUCTION config: a bidder roster or a prebid global |
| `ad-serving-only` | 18 | Adfox/Yandex ad code present, no auction config readable — this publisher sells ads, but not by asking several bidders in the page |
| `no-wrapper` | 24 | page loaded, no Adfox and no prebid (re-probed at the full 45 s window) |
| `ya-generic-only` | 63 | a Yandex namespace but no advertising global — Metrika, not ad-tech |
| `blocked` | 23 | 4xx/5xx — **we never saw the page**, so this says nothing about its wrapper |
| `rate-limited` | 1 | 428/429 — caused by our own load, not by the site |
| `unreachable` | 20 | timeout or refused connection — our blindness |
| `no-web-apex` | 45 | DNS/TLS failure — the domain serves no page at its apex at all |

Admitted: `magnit.ru`, `pikabu.ru`, `rbc.ru`, `gismeteo.ru`, `hh.ru`, `ria.ru`, `iz.ru`, `matchtv.ru`, `kp.ru`, `aif.ru`, `mk.ru`, `sports.ru`, `aviasales.ru`, `interfax.ru`.

**Vantage `nl-direct`** · 254 domains touched in rank order · browser egress 77.246.104.228 · `frame-stageb-nl-direct-2026-08-19-schema3.jsonl`

| verdict | n | what it means |
|---|---:|---|
| `admit` | 13 | **in the frame** — 200, HTML, and an AUCTION config: a bidder roster or a prebid global |
| `ad-serving-only` | 20 | Adfox/Yandex ad code present, no auction config readable — this publisher sells ads, but not by asking several bidders in the page |
| `no-wrapper` | 33 | page loaded, no Adfox and no prebid (re-probed at the full 45 s window) |
| `ya-generic-only` | 81 | a Yandex namespace but no advertising global — Metrika, not ad-tech |
| `blocked` | 33 | 4xx/5xx — **we never saw the page**, so this says nothing about its wrapper |
| `rate-limited` | 1 | 428/429 — caused by our own load, not by the site |
| `unreachable` | 19 | timeout or refused connection — our blindness |
| `no-web-apex` | 54 | DNS/TLS failure — the domain serves no page at its apex at all |

Admitted: `magnit.ru`, `pikabu.ru`, `ria.ru`, `lenta.ru`, `iz.ru`, `matchtv.ru`, `kp.ru`, `aif.ru`, `mk.ru`, `sports.ru`, `liveinternet.ru`, `sportbox.ru`, `playground.ru`.

### Captured cells

Three counts, because they are three different facts. `this arm` is what this arm attempted. `both arms` adds the other arm's load — they differ exactly when one arm refused a load the other recorded. **`readable`** additionally requires that BOTH arms saw the page, and a pair with an unseen half carries no RU-versus-NL contrast at all. Only the last number is a comparison. The two ways of not seeing are kept apart: `refused` is a 4xx/5xx — the publisher's server answered and declined THIS vantage — while `no answer` is a navigation that never returned a response at all, which is our own link and says nothing about the site.

| cell | arm | loads | refused | no answer | sites | this arm | both arms | **readable** | window | hours UTC | vantage verified |
|---|---|---:|---:|---:|---:|---:|---:|---:|---|---|---:|
| `2026-08-19-day-2026081914` | `nl-direct` | 14 | 2 | 0 | 14 | 14 | 10 | **3** | 45 s | 14,15 | 14/14 |
| `2026-08-19-day-2026081914` | `ru-mobile` | 10 | 0 | 6 | 10 | 10 | 10 | **3** | 45 s | 14 | 7/10 |
| `2026-08-19-evening-2026081915` | `nl-direct` | 56 | 8 | 4 | 14 | 56 | 56 | **27** | 45 s | 15,16,17,18 | 56/56 |
| `2026-08-19-evening-2026081915` | `ru-mobile` | 56 | 0 | 26 | 14 | 56 | 56 | **27** | 45 s | 15,16,17,18 | 50/56 |
| `2026-08-19-evening-2026081919` | `nl-direct` | 14 | 2 | 2 | 14 | 14 | 14 | **9** | 45 s | 19,20 | 14/14 |
| `2026-08-19-evening-2026081919` | `ru-mobile` | 14 | 0 | 5 | 14 | 14 | 14 | **9** | 45 s | 19,20 | 12/14 |
| `2026-08-19-morning-2026081906` | `nl-direct` | 6 | 0 | 0 | 2 | 6 | 6 | **5** | 45 s | 6 | 6/6 |
| `2026-08-19-morning-2026081906` | `ru-mobile` | 6 | 0 | 1 | 2 | 6 | 6 | **5** | 45 s | 6 | 6/6 |
| `2026-08-19-night-2026081921` | `nl-direct` | 46 | 9 | 7 | 14 | 46 | 44 | **13** | 45 s | 21,22,23 | 46/46 |
| `2026-08-19-night-2026081921` | `ru-mobile` | 48 | 0 | 23 | 14 | 48 | 44 | **13** | 45 s | 21,22,23 | 41/48 |
| `2026-08-20-evening-2026082019` | `us-exit` | 14 | 1 | 0 | 14 | 14 | 14 | **13** | 45 s | 19,20 | 14/14 |
| `2026-08-20-evening-2026082019` | `ru-mobile` | 14 | 0 | 0 | 14 | 14 | 14 | **13** | 45 s | 19,20 | 14/14 |
| `2026-08-20-evening-2026082020` | `us-exit` | 18 | 2 | 1 | 13 | 18 | 10 | **9** | 45 s | 20,21 | 18/18 |
| `2026-08-20-evening-2026082020` | `ru-mobile` | 11 | 0 | 0 | 11 | 11 | 10 | **9** | 45 s | 20 | 11/11 |
| `2026-08-20-night-2026082021` | `us-exit` | 25 | 3 | 1 | 14 | 25 | 15 | **9** | 45 s | 21,22 | 25/25 |
| `2026-08-20-night-2026082021` | `ru-mobile` | 15 | 0 | 5 | 14 | 15 | 15 | **9** | 45 s | 21,22 | 15/15 |
| `2026-08-20-night-2026082022` | `us-exit` | 28 | 4 | 1 | 14 | 28 | 28 | **15** | 45 s | 22,23 | 27/28 |
| `2026-08-20-night-2026082022` | `ru-mobile` | 28 | 0 | 10 | 14 | 28 | 28 | **15** | 45 s | 22,23 | 28/28 |
| `2026-08-20-night-2026082023` | `us-exit` | 28 | 4 | 1 | 14 | 28 | 28 | **15** | 45 s | 0,23 | 28/28 |
| `2026-08-20-night-2026082023` | `ru-mobile` | 28 | 0 | 10 | 14 | 28 | 28 | **15** | 45 s | 0,23 | 28/28 |
| `2026-08-21-night-2026082100` | `us-exit` | 28 | 5 | 1 | 14 | 28 | 28 | **14** | 45 s | 0,1 | 28/28 |
| `2026-08-21-night-2026082100` | `ru-mobile` | 28 | 0 | 10 | 14 | 28 | 28 | **14** | 45 s | 0,1 | 28/28 |
| `2026-08-21-night-2026082102` | `nl-direct` | 28 | 5 | 0 | 14 | 28 | 28 | **15** | 45 s | 2,3 | 28/28 |
| `2026-08-21-night-2026082102` | `ru-mobile` | 28 | 0 | 11 | 14 | 28 | 28 | **15** | 45 s | 2,3 | 27/28 |

**Coverage against the design:** 147 of 4480 readable pairs (3.3 %) of the 14-day, two-arm schedule in §2.3 of `docs/step0c-hb-ru-market-study-2026-08-18.md`. That is the study's own unit: a pair is readable only when BOTH arms saw the page, and only such a pair carries a contrast. Measured in loads attempted the same work reads 591 of 8960 (6.6 %) — the larger figure is what we have SPENT, the smaller is what we can USE, and the gap is the systematic per-site blindness named in the table above, not noise. Everything measured so far is a pilot of that design, not a sample of the Russian market.

> **Some cells carry loads outside the daypart they are named for.** `2026-08-19-day-2026081914` is named `day` but 1 evening load(s) of 24 fall outside it; `2026-08-20-evening-2026082020` is named `evening` but 5 night load(s) of 29 fall outside it; `2026-08-21-night-2026082102` is named `night` but 18 morning load(s) of 56 fall outside it. A cell takes its name from the daypart it STARTED in and keeps it for the whole run, so a long cell crosses the boundary. The design in §2.3 stratifies by daypart, so these loads are filed under a daypart they were not sampled in. They are named here rather than re-bucketed — which stratum they belong to is a decision about the study, not about this table.

<!-- /NUMBERS -->

**The word *market* is not available to us yet.** What we have is a design that says exactly
what would earn it (`docs/step0c-…`), a frame built by a published rule, and the first cell of
that design actually run (`docs/step0d-…`). Every figure below and in the documents carries
its `n` and its coverage in the same line as the number, and no chart in this repository is
captioned as though it covered more.

---

## What we can see, and what we are blind to

This matters more than any result here, so it is on the front page rather than in a footnote.

**Observable:** what an ad wrapper hands the browser — which bidders a publisher's own config
asks, which of them answered, which stayed silent, what price came back when one did, and the
currency and floor when the request carries them.

**Not observable, at all, from this vantage:**

- **The server-side auction.** Server-to-server bidding never touches the browser. A site we
  record as carrying no wrapper may be selling its inventory perfectly well.
- **Any page we did not see.** An HTTP 401/403/5xx is *our own access failing*, and it is
  recorded as `blocked`, never as "this publisher runs no header bidding". Converting our
  blindness into a market fact is the one mistake this project cannot afford.
- **Anything about anyone else.** By construction.

---

## The rules the code is built to, not merely to follow

- **Only his device.** A record whose device id is not the target never reaches disk — dropped
  in the parser before any buffer, counted as a bare integer, never as a row or an aggregate.
  This is why monitor-mode Wi-Fi capture is rejected outright rather than filtered afterwards.
- **`UNKNOWN` is a first-class verdict.** A host the reference table cannot name is reported
  *with its count*, and the coverage ratio is printed on every run. `LOSS` is not `NOT IN`.
- **Absence is split by cause, always.** `no-bid` / `no-answer` / `blocked` / `unreachable` /
  `no-web-apex` / `no-wrapper` are six different facts. Folding any two of them is how a
  measurement of our own latency becomes a finding about bidder selectivity.
- **A silent fallback is a fabricated success.** A missing reference table raises rather than
  rendering every host `UNKNOWN` and letting the run "succeed" — a table of nothing looks
  exactly like a finished measurement.
- **A gate you have not seen fail is not a gate.** Every property above was broken
  deliberately and watched go red before it was accepted. `docs/REPRODUCE.md` lists which
  mistake each gate exists to catch.
- **A vantage is part of every reading.** Each row records the exit address the **browser**
  read, not the one the shell measured — on our node those differ, and a row that cannot prove
  its own path is discarded rather than labelled.

## Layout

    tools/         collectors and reporters, one concern each, every one with a --test that
                   runs offline against a local fixture:
                     adint-frame-stageb    builds the sampling frame from the pinned universe
                                           and PUBLISHES the rejections with their reasons
                     adint-paired-run      runs one cell of the schedule from both vantages,
                                           barriered site by site so a pair is really a pair
                     adint-hb-capture      drives a browser, records the wrapper's auction,
                                           and proves its exit from INSIDE the browser
                     adint-hb-report       counts it; refuses to pool schemas or fold no-answer
                     adint-publish         copies observations to public-data/ through an
                                           explicit field ALLOWLIST
                     adint-aggregate       observations -> bundle-exchange table
                   self-test any of them:  python3 tools/<name> --test
                   (browser-driving tools need ~/.venv-browser/bin/python3; run under the
                   wrong interpreter they exit 2 = n/a, never a false FAIL)

    ref/           the pinned inputs: the Tranco universe, domain -> org -> exchange tables,
                   and each frame produced from them
    public-data/   the observations, world-readable, CC0. Written only by adint-publish
    data/          HIS capture. Gitignored, and stays that way — this is the privacy boundary
    docs/          design, method, fact base. See docs/INDEX.md
    tests/         fixtures

This repository is deliberately **not** part of the operator's mesh: nothing here is a mesh
organ, no cron, no reflex. It consumes the mesh (the phone body, the board, the voice channel)
and never vendors it.

## Licence

**CC0 1.0** — public domain. Take it, fork it, use it to argue we are wrong. See `LICENSE`.

A public-domain dedication is a statement about rights, not about truth: every number here is
published with its `n` and its vantage so that you can judge it rather than inherit it.
