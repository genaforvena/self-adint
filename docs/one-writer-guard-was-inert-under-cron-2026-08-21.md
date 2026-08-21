# The one-writer guard was green in every test and inert in the only place it ran

*2026-08-21. `tools/adint-collect-cell`. Found while checking a completed cron tick.*

## What was wrong

`adint-collect-cell`'s docstring names one writer as the first of three things it refuses
to get wrong: the RU arm rides one netns over one USB tether, and two cells at once would
interleave two vantages on one link. The guard read the process table:

```python
out = subprocess.run(["/usr/bin/ps", "-u", os.environ.get("USER", ""), "-o", "pid=,args="], ...)
```

**Cron does not set `USER`.** Measured on this node from the environ of a live cron child
rather than assumed:

```
HOME=/home/mesh-home   LOGNAME=mesh-home   PATH=...   SHELL=/bin/sh   LANG=en_US.UTF-8
```

`LOGNAME`, no `USER`. So under cron the call was `ps -u "" -o pid=,args=`, which exits **1**
with `error: list of users must follow -u` and prints nothing. The returncode was never
read; the empty stdout was iterated; the guard answered *no cell is running*. Every tick,
for as long as the fortnight ran.

## Why the test never saw it

The test is not lazy. It was written against a previous self-matching grep and says so:
it starts a **real process** whose argv looks like a running cell, watches the guard find
it, kill it, and stop finding it. All three checks are honest and all three passed.

They passed because **a test is run from a shell, and a shell sets `USER`.** The gate was
green in the only environment it was ever exercised in and dead in the only environment it
mattered in. Nothing about the assertion was wrong; the *fixture* was the wrong world.

## Two fixes, because they fail differently

**Ask by uid.** `os.getuid()` needs no environment to be true. An identity taken from a
variable is a claim about the environment; an identity taken from the kernel is a fact.

**A failed probe is not an answer.** Returning `None` on exception made *I could not look*
and *I looked and found nobody* the same value — a silent fallback turning total failure
into the one reply that lets a second cell start. `running_pid()` now returns
`PROBE_FAILED`, and the caller **skips the tick and records why**. A skipped tick is
recoverable and writes itself down; two cells interleaving two vantages on one tether is
neither. The two fixes compose: with the returncode checked, even the old `$USER` bug
degrades to a loud skip instead of a silent open door.

## The gates, each seen red

- *the fixture is the real thing — cron's environment has no USER* — asserts the fixture
  before trusting it, so this cannot rot into a test of a world that no longer exists.
- *the guard still finds a live cell with USER unset, as cron runs it* — runs the guard
  inside cron's measured environment against a real process. Restoring the original
  `$USER` line turns it red (`got 'probe-failed', want <pid>`).
- *a probe that cannot read the process table says PROBE_FAILED, not None* — points the
  guard at a `ps` that always exits 1.
- *a --test process is not mistaken for a cell* — was a grep of this file for the literal
  `"--test" not in line`, on a line that **contains** that literal, so it matched its own
  text and could never fail. Two lines below a comment warning about precisely that. Now
  driven with a real process whose argv carries `--test`.

## Blast radius

Cells run ~58 min against a 2 h cadence and are killed at 3 h, so an overlap needed a cell
between two and three hours. None has run that long yet, and the corpus on disk is
believed clean — but that was luck, not the guard.

## The shape

A guard selecting on `$USER`, `$HOSTNAME`, or any other ambient variable is testing the
environment it happens to be run in. This mesh has the same scar twice already: a keeper
gated on `TG_HOST="imozerov-IdeaPad-…"` whose body never executed once, and an env gate
armed by a file cron never sources. **Bind a guard to the thing itself, never to a name
the environment supplies — and make the test run in the environment the guard will
actually meet.**
