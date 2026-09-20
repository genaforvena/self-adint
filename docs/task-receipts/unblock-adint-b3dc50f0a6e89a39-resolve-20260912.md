# Resolver receipt: `unblock/adint/b3dc50f0a6e89a39/resolve`

- Checked: 2026-09-12T01:53:58Z on `mesh-home`
- Parent: `unblock/haunt/fd5d75268b44e1cc/resolve` (owner `haunt`)
- Result: **operator-action packet recorded; parent remains BLOCKED/operator-input**

## Evidence and boundary

The dispatch row was checked with `mesh-task check dispatch unblock/adint/b3dc50f0a6e89a39/resolve adint` (exit 0), then claimed by `adint`. At 01:53:58Z the parent status was still `blocked`, and its owner resume gate, `mesh-task check resume unblock/haunt/fd5d75268b44e1cc/resolve haunt`, exited 2. The matching operator dictionary manifest and corpus are not present in the task evidence. Runtime smoke does not establish the dictionary arm's source or semantics.

The exact request was already sent to the operator over voice and duplicated text at 2026-09-12 01:48 UTC; see `/home/mesh-home/lte-workstation/docs/task-receipts/unblock-adint-eb4c5cff68622566-resolve-20260912.md` (SHA-256 `7490711a81fce0a7ecdaa0f4c1fd10083750aea6df0a796b0aa341aff07e3042`). No second ask is needed while that request is pending. Choosing a dictionary source, language, normalization, immutable version, or corpus here would invent scientific inputs, so no such choice or retry was made.

## Exact operator input required

Supply and freeze all five fields, with paths/URIs resolvable by the runner:

```text
source: operator-approved dictionary/lexicon URI or repository-relative path
language: ISO language code or explicit language set
normalization: exact Unicode and token-normalization policy, including implementation/version
version: immutable release, commit, or source checksum
corpus: immutable corpus path/URI and SHA-256 of the exact corpus bytes
```

The next action on `event:operator-dictionary-input` is to validate the approved source and all five fields, independently recompute the corpus SHA-256, run the parent dependency preflight and dictionary arm, then—only if those checks pass—run the bounded retry:

```bash
cd /home/mesh-home/tiny-fleet
timeout 240 .venv/bin/python scripts/bbywvy_test.py
```

Until that event and validation, the owning `haunt` task remains blocked; no corpus, package, runtime, or model was changed and no study retry ran.
