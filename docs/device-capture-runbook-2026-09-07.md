# Device-side capture handoff

The operator still controls PCAPdroid's VPN and export. `tools/adint-device-capture` does not
start or stop it, click Android UI, or contain a phone address or credential.

After exporting a complete CSV from the operator's own Redmi, run:

```text
tools/adint-device-capture --host <termux-user>@<phone-host> \
  --remote-csv <remote-export.csv> \
  --out data/device-observations.jsonl \
  --device redmi10-daily --pull
```

Preview the exact read-only transfer first with `--dry-run`. The command writes the JSONL and its
`.manifest.json` only after `scp` and the parser both succeed. The manifest binds the output to the
CSV SHA-256 and the declared device label. A failed pull or malformed export leaves no partial
measurement artifact.

The importer keeps system-resolver rows separate, reports malformed rows and ring overflow, and
stamps every retained record with the declared device. It does not claim that a hostname was sold
by an exchange: use `tools/adint-aggregate` for the demand-stack entry-point table.
