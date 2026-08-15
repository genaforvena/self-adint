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
3. **No irreversible step without his explicit go, per step.** Onboarding a seat, signing an
   agreement, paying a deposit, contacting an exchange, resetting his GAID — the mind **drafts**,
   he ships. His "го" of 2026-08-15 authorises organising the channel and the plan, nothing beyond.
4. **Nobody but him is in scope.** No observation of any other person or device, ever, including
   incidentally. If a design step would require it, the step is wrong, not the rule.
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
