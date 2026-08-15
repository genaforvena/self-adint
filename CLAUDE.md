# self-adint — channel charter

Separate repository, deliberately **not** part of the mesh genome (`lte-workstation`). Nothing here
is a mesh organ: no `mesh-*` tool, no `# reflex-cadence:` header, no cron. The genome's
`mesh-doctor` / `mesh-autowire` / `mesh-land` must never see this code — a project with its own
logic living in the genome only generates orphan warnings and false reflex gates.

The mesh relationship runs one way: this project **consumes** mesh organs (the phone body, the LAN
sense, `mesh-tell`, `mesh-chat`, `mesh-voice-tx`) and never imports or vendors them.

## What this is

The operator reading **his own** advertising profile from the demand side: stand up an RTB seat,
receive bid requests, and read what an audience buyer is handed for free about his own device.
His full brief is the fact base — `docs/brief-2026-08-15-operator.md`, verbatim. The working plan,
including the corrections he accepted, is `PLAN.md`.

**The goal of the first pass is not data. It is finding out what arrives for free.** A cycle that ends
in a summary is a failure; it ends in a measured artifact or an answered question.

## Hard rules — none of these are negotiable by a mind in this window

1. **Only his device.** Rows whose `device.ifa` is not the target are never written to disk — not
   as raw lines, not as aggregates, not as "noise statistics". The filter lives in the parser
   **before the buffer**, never in a query against storage.
2. **The leak will not come from storage.** The decoy path sees 100% of foreign traffic by
   construction. So: no request-body logging anywhere, a `recover` that never prints the request,
   no access log carrying the payload, metrics as **counters only**. Any framework default that
   dumps a request on panic is a defect to be removed before the first real packet.
3. **No UNRECOVERABLE step without his explicit go, per step** (narrowed 2026-08-15 on his own
   objection — *«да ну а как же автопоэзис? давайте и почту вы сделаете и отправите тоже вы»*: a
   project that stops for a human hand at every step does not self-produce, and the rule as written
   made the hand mandatory everywhere). The gate now stands exactly where an error does not play
   back: **onboarding a seat, signing an agreement, any payment, resetting his GAID.** There a
   mistake costs money or destroys the baseline the comparison is made of.

   **DELEGATED to the mind, end to end:** creating the dedicated mailbox and sending the data-subject
   letters. Three commitments carry that delegation, and they are the reason it is safe: (a) **no
   invented fact** — every evidentiary line in a letter is drawn from his own capture (host, time,
   bytes, app), never from memory or "approximately"; (b) **every sent letter lands on disk and
   reaches him the moment it goes out** — for reading, not for approval; a correction is an ordinary
   follow-up in this genre; (c) **one ledger** of who was written to, when, what came back, who is
   silent. His "го" of 2026-08-15 authorises the channel, the plan, the local Prebid step, and this
   delegation — nothing beyond.
4. **Nobody but him is in scope — and the line is the DISK, not the socket** (revised 2026-08-15 on
   his explicit word: *«на жёсткий диск никогда не попадает, если это не я»*). The old text banned
   seeing a foreign request "even incidentally", which the project cannot honour the moment it
   BIDS: to answer an auction you must parse the request you are answering. Pretending otherwise
   would have meant a rule the code quietly breaks — worse than a rule that says what we do. So,
   binding: a foreign request lives **in memory, for the duration of the bid, and nowhere else**.
   Not one line, not one aggregate, not one counter keyed by anything about that person reaches
   storage; the filter stays in the parser **before the buffer** (rule 1), and rule 2's no-body-log
   discipline is what makes this true under panic, access log, and metrics alike. Nothing about
   another person is ever *retained*, *derived from*, or *reported* — his ad may be served to
   whoever's request arrives, and that person leaves no trace behind them. If a design step would
   require keeping anything about someone else, the step is wrong, not the rule.
5. **The VM is outside the mesh.** It is a public TLS endpoint with a static whitelisted IP. It
   carries no Tailscale, no mesh reachability, no mesh credentials, and no node of the mesh routes
   through it. It cannot run on `mesh-home` regardless — public ingress here is closed.
6. **His words are the fact base.** Any number, claim, or intent used anywhere in this project
   comes from `docs/brief-2026-08-15-operator.md` or from him directly. A paraphrase is a number
   with no source. Where his brief arrived garbled, the gap is marked — ask, do not reconstruct.

## Doctrine inherited from the mesh (it applies here in full)

- **Every claim is an artifact.** Not "the receiver parses bid requests" — a JSONL file on disk
  with parsed rows, and a test seen RED before it was seen green.
- **A silent fallback is a fabricated success.** `|| 204`, a default `k`, a swallowed parse error:
  if a default is indistinguishable from a real result, it will be mistaken for one.
- **Absence is not a negative.** This is the whole shape of the oracle here: `LOSS` is not
  `NOT IN`. `UNKNOWN` is a first-class verdict and must survive all the way to the report.
- **Measure before committing.** The cheapest measurement that can invalidate an expensive,
  irreversible step is run first. That is why `PLAN.md` opens with a step his own draft did not
  have.
- **Calibrate against the real corpus, never an assumed constant.** `k`, the geo/bundle filters,
  the impression frequency: all derived from what is observed, all re-derived as the window grows.

## Working conventions

- The mind of the `adint` window owns this repo. It works in its own pane, posts to the board
  (`mesh-chat`) in its own voice, and speaks to the operator over `mesh-voice-tx`.
- Subagents for heavy reading (spec sweeps of OpenRTB, exchange documentation, log analysis).
  A subagent's report is a claim, not an artifact — check the artifact before acting on it.
- **The repo is PUBLIC on his go of 2026-08-15** (`genaforvena/self-adint`) — the code and the
  reasoning are published. That go covers the tracked tree and nothing else: **`data/` is
  gitignored and stays gitignored**, because it is where his observations land — his installed
  apps, his hosts, his device id. Committed files carry no node identifiers either (the
  reachability table in `docs/step0-method-2026-08-15.md` is written with `<phone-lan-ip>`-style
  placeholders on purpose). Before any commit, check that what you are publishing is the method,
  not the measurement.
- Language with the operator: Russian, voice + duplicated text.
