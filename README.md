# self-adint — reading your own advertising profile from the demand side

An operator points the ad-tech machinery at himself: stand on the **buy** side of the real-time
bidding market and read what an audience buyer is handed, for free, about **his own device**.

The thesis is not "how much will impressions cost me". It is: *the audience profile arrives inside
the bid request itself.* You do not have to buy the impression to read what came with it —
`user.data[]` segments, `user.ext.eids[]` identity graphs, the consent string, and the `geo.type`
precision class per app bundle.

The full brief in the operator's own words is `docs/brief-2026-08-15-operator.md` — the fact base
of this project, kept verbatim. `PLAN.md` is what actually gets executed: his order of work with
one step inserted in front of it and four corrections folded in.

## Where it stands

**Step 0 — which exchanges actually sell his phone's traffic?** Inserted before everything else
because the one irreversible step in this project (onboarding a bidder seat: legal entity,
agreement, whitelist, deposit) currently rests on a *guess* about which exchange carries his
device — and that guess can be turned into a measurement, read from the device side, for zero
cost. If his apps sell through exchange A and the seat is opened at exchange B, his device never
appears in the sample, and the failure is discovered after the money.

The artifact is a `bundle → exchange` table with observation counts. Not "app X shows ads".

Method record and what was rejected, with reasons: `docs/step0-method-2026-08-15.md`.

## Rules the code is built to, not merely to follow

- **Only his device.** A record whose device id is not the target never reaches disk — dropped in
  the parser before any buffer, counted as a bare integer, never as a row, an aggregate, or a
  "noise statistic". Nobody else is in scope, ever, including incidentally: this is why
  monitor-mode Wi-Fi capture is rejected outright rather than filtered afterwards.
- **`UNKNOWN` is a first-class verdict.** A host the reference table cannot name is reported *with
  its count*, and the coverage ratio is printed on every run. A negative result and an absent
  result are different things — the entire three-valued oracle downstream (`IN` / `UNKNOWN` /
  `NOT`) exists because `LOSS` is not `NOT IN`.
- **A silent fallback is a fabricated success.** A missing reference table raises, rather than
  rendering every host `UNKNOWN` and letting the run "succeed" — a table of nothing looks exactly
  like a finished measurement.
- **A gate you have not seen fail is not a gate.** Each property above was broken deliberately and
  watched go red before it was accepted; the mutations are listed in the method record.

## Layout

    tools/adint-aggregate   observations (JSONL) -> bundle-exchange table + unresolved hosts
                            self-test: python3 tools/adint-aggregate --test
    ref/                    public knowledge: domain -> org -> does it run an OpenRTB exchange,
                            can an outsider get a seat there
    docs/                   the fact base, the method record, the research notes
    data/                   his observations — GITIGNORED, and stays that way

This repository is deliberately **not** part of the operator's mesh genome: nothing here is a mesh
organ, no cron, no reflex. It consumes the mesh (the phone body, the board, the voice channel) and
never vendors it.
