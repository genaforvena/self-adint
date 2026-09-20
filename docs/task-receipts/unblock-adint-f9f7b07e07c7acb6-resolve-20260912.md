# Resolver receipt: `unblock/adint/f9f7b07e07c7acb6/resolve`

- Checked: `2026-09-12T02:14:49Z` UTC on `mesh-home`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **resolver settled with an operator-action packet; parent remains BLOCKED/operator-input**

## Fresh evidence

The owner-scoped dispatch row was checked before taking it. The owner-authored command
`MESH_TASK_ACTOR=adint mesh-task take unblock/adint/f9f7b07e07c7acb6 resolve` succeeded;
the live task ledger now shows this resolver `active`, owned by `adint`.

At 02:14 UTC, `mesh-task status unblock/haunt/fd5d75268b44e1cc` still reported the parent
`blocked` with blocker `operator-input`. The parent's owner resume gate,
`mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`, exited `2`.
The latest board search for `operator-dictionary-input` contains no supplied manifest or
corpus after the request was sent at 01:48 UTC. The request and its Russian voice/text receipt
are recorded at `/home/mesh-home/lte-workstation/docs/task-receipts/unblock-adint-eb4c5cff68622566-resolve-20260912.md`
(SHA-256 `7490711a81fce0a7ecdaa0f4c1fd10083750aea6df0a796b0aa341aff07e3042`). No duplicate
request was sent.

No safe local choice can fill the gap: selecting the dictionary source, language, normalization,
immutable version, or corpus would invent study inputs. The runtime smoke does not exercise or
discharge the dictionary arm. No source, corpus, package, runtime, or model was changed, and no
study retry was run.

## Exact operator input and next action

Freeze all five fields in a tracked manifest and provide the exact matching corpus bytes:

```yaml
source: <operator-approved dictionary/lexicon URI or repository-relative path>
language: <ISO language code or explicit language set>
normalization: <exact Unicode and token-normalization policy, including implementation/version>
version: <immutable release, commit, or source checksum>
corpus: <immutable corpus path/URI and SHA-256 of the exact corpus bytes>
```

On `event:operator-dictionary-input`, validate the approved source and all five fields,
independently recompute the corpus SHA-256, run the parent dependency preflight and dictionary
arm, and record their outputs. Only if those checks pass, run:

```bash
cd /home/mesh-home/tiny-fleet
timeout 240 .venv/bin/python scripts/bbywvy_test.py
```

Until that input is supplied and validated, the `haunt` owner resume gate must remain refused.
