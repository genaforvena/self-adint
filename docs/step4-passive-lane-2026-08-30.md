# Step 4 — the passive collection lane, wired, and the guard that was blocking it

*2026-08-30. `tools/adint-collect-passive`, wired hourly in the user crontab on mesh-home.*

## The measured reason this repo stalled

The board's own reading of the 2026-08-21 stop was not a missing idea. It was that **`adint`
matched zero cron lines**: every cell this study ever collected was launched by a hand at
the end of an agent's turn, and Step 4 asks for a *week*. A step everything downstream
depends on has to have its own unconditional cadence, because an agent's loop dies with its
engine and nobody remembers for seven days.

As of today `crontab -l | grep -c adint` is no longer 0.

## Why this is a new tool and not `adint-collect-cell`

`adint-collect-cell` drives `adint-paired-run` — Step 0's **vantage study**: two named arms,
and a preflight that correctly refuses the cell when they collapse onto one exit. On this
node they now do, for two independent reasons measured the same day:

- the Note 3 tether is physically gone — no `ruvantage` netns, no `enx*` device — so the
  `ru-mobile` arm has no device to run on;
- the host's own route and privoxy's far end **both** exit `38.49.216.141`, because this
  node consumes phaedra as its tailscale exit node while the socks tunnel lands on phaedra
  too. A proxy/direct arm pair is an A/A null here.

Wiring that tool on a cadence would have produced a lane that refuses every tick. **A reflex
that runs green and collects nothing is worse than no reflex**, because its cron line reads
as coverage.

Step 4 does not need a pair. It asks what arrives about him over a week — what is in
`user.data[]`, who is in `eids[]`, `geo.type` per bundle, how often his own impressions come
round. That is one vantage observed for a long time, and it makes no cross-vantage claim.

## The guard that was blocking it, and the third value it was missing

The first single-arm run refused outright:

    preflight vantage: via=direct exit=38.49.216.141 echo_attempts=1/1
    refusing: --via direct and the proxy's exit both leave from 38.49.216.141 —
    the labels differ and the treatment does not

That check exists for a real failure — a `--via direct` browser that leaked back through the
inherited privoxy is not the run its label claims. But the comparison is **only decidable
while the node's own route and the proxy's far end are different addresses.** Here they are
the same address, so a leaked browser and a perfectly direct one produce the *identical*
reading. Refusing there is not caution; it is treating an undecidable case as a positive.
Same shape as a probe that blames the far end for a failure mode both ends share.

`via_direct_leak_verdict()` now has three outcomes and the third is a word, not a refusal:

| world | verdict |
|---|---|
| the parent's two legs differ, the browser matched the **direct** one | `proven-distinct` (rc 0) |
| the legs differ, the browser matched the **proxied** one | refuse (rc 1) — a real leak |
| the legs **coincide**, or the parent's direct leg is unreadable | `unprovable-…` (rc 0) |
| no proxy in the environment at all | `n/a: nothing to leak into` |

The word travels onto **every load row** (`via_direct_leak_check`), because a run whose check
was `unprovable-*` is not a run that passed it, and a reader joining these rows a month from
now has no other way to tell a proven vantage from an undecidable one. Today's rows all read
`unprovable-single-exit: … both 38.49.216.141 …` — which is the honest description of this
node, and is exactly the sentence that would have been silently absent had the guard simply
been relaxed.

Seven gates, driven over every world the function can meet, including the two that must
never read as a pass. The refusal arm is still there and still refuses.

## What the lane refuses to get wrong

- **One writer, bound to the kernel.** The lock is checked by `os.kill(pid, 0)` *and* by
  reading `/proc/<pid>/cmdline` for this tool's own name — pid reuse on this node is hourly,
  so a bare pid check would eventually wedge the lane behind a recycled pid, a stand-down
  that looks exactly like a busy one. The guard never reads `$USER`: cron sets `LOGNAME` and
  no `USER`, which is precisely how this project's previous single-writer guard was green in
  every test and dead in production.
- **A tick that collected nothing still writes a row.** Every path appends an outcome with a
  reason — `collected`, `partial`, `collected-nothing`, `stood-down`, `dry-run`, `crashed`.
  A tape of only the successful ticks is a numerator, and over a week the difference between
  "the market was quiet" and "our browser never started" is the difference between a finding
  and a fabrication.
- **The outcome word is decided by what reached the corpus, not by the child's exit code.** A
  capture can exit 0 having refused every load, and exit non-zero after writing most of them.
  Both facts are recorded; only one of them is the coverage.
- **The cursor moves before the work.** The frame is walked with a persistent cursor so a
  week covers it round and round instead of re-sweeping its head forever, and it is advanced
  *before* the sites are loaded. A tick killed halfway must not hand its successor the same
  sites; the cost is that a crashed tick's sites are skipped once and the ledger says so,
  where the opposite cost — repeating them silently — is unbounded, because nothing
  downstream can tell a re-load from a fresh one.
- **Egress, both legs, in the same run as the rows.** Neither "this node is proxied" nor
  "this node is direct" is a fact about the node — it is a fact about the process. Under
  cron there are no proxy variables; in a mind's pane there are. The row carries both legs
  and `legs_agree`.
- **The ledger is append-only and whole.** No window, no pruning: an `n` quoted against a
  sliding ledger is not reproducible and can move *down*.

## The first live tick

    outcome  collected            3/3 sites (rc=0)
    sites    magnit.ru, pikabu.ru, ria.ru      cursor_from 0
    loads    3        bid_candidates 21        bytes_added 1052193
    seconds  199.4    free_gb 772.94
    egress   as-invoked 38.49.216.141 US/AS197537 ·
             proxy-stripped 38.49.216.141 US/AS197537 · legs_agree true
    proxy_env HTTP_PROXY=HTTPS_PROXY=ALL_PROXY=http://127.0.0.1:8118

258 rows, 21 bid candidates, named bidders across the three sites — `hybrid`,
`betweendigital`, `sape`, `mytarget`, `buzzoola`, `astralab`, `getintent`, `videonow`.

## The cadence, and where it is written

    13 * * * * cd $HOME/self-adint && ./tools/adint-collect-passive \
        >> $HOME/self-adint/data/passive-collect.cron.log 2>&1

Hourly, 4 sites a tick at 30 s settle — about 5 minutes of browser per hour, ~96 loads a
day. It is wired in the **user crontab and nowhere else**: this project's charter forbids its
code being a mesh organ, so the tool carries no `# reflex-cadence:` header and is invisible
to `mesh-autowire` and `mesh-doctor` by construction. The existing 301 crontab lines were
verified byte-identical after the write.

## Echo redundancy added 2026-09-07

Each passive tick now records the ordered attempts against two independent IP-echo providers for
each process leg. A dead provider is retained as a failed attempt and the next provider may still
establish the leg; if all providers fail, the leg's IP remains `null` and the tick is UNKNOWN rather
than a fabricated egress. This closes the single-provider outage shape without turning provider
availability into a market result.

## What this lane is not

It is not the seat. Step 4 as the plan writes it reads bid requests arriving *for his
device* through an exchange seat, and that needs Step 3. What this collects is the bid
requests the public web's own header-bidding wrappers send from **this** vantage — real
OpenRTB traffic with real audience payloads, and the right corpus for the `user.data[]` /
`eids[]` / `geo.type` questions, but about the browser doing the browsing, not about his
phone. The device-side half is measurable too — PCAPdroid is installed on the Redmi and
reachable over ssh — but its capture cannot yet be started, stopped or exported headlessly:
Termux's `am` wrapper fails with a reflection error against the control activity on that
device, so driving it is an open problem and not a wired lane. Saying so is the point; a
lane that silently collected only half of what its name implies is the failure this document
opens with.
