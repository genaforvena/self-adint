# Receiver red-before-green — 2026-09-07

The receiver privacy gate was tested against a real mutation in an isolated temporary copy. The
mutation changed the parser decision from matching the configured target IFA to matching every
request.

```text
baseline repository:  go test ./...  PASS
mutated copy:        go test ./...  RED
  TestDecoyRule: foreign requests reported matched=true
  TestForeignIFANeverWritten: expected 1 persisted row, got 3
restored repository: go test ./...  PASS
```

The temporary mutant was outside the repository and was not used for the fresh run. The restored
repository produced `receiver/testdata/step1-2026-09-07-final/`.
