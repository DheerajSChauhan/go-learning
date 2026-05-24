# Go Learning Repository 🚀

This repository is a personal collection of Go (Golang) examples and small projects used to learn and demonstrate Go language concepts, standard library usage, and small real-world patterns (HTTP, JSON, CRUD, auth).

## Goals
- Provide concise, runnable examples for Go language fundamentals.
- Demonstrate module-based projects, HTTP clients/servers, JSON handling, and simple CRUD and authentication patterns.
- Serve as a reference and teaching aid for common Go idioms.

## Prerequisites
- Go 1.20+ installed and on your PATH (see https://go.dev/dl).
- Git (for cloning and pushing changes).

## Quickstart — run an example
To run a single example folder that contains a `main.go`:

Windows (PowerShell):
```powershell
cd "01_Setup_first-program"
go run main.go
```

Generic (from repo root) — run a single file:
```bash
go run ./01_Setup_first-program/main.go
```

To run all packages (useful inside a module folder):
```bash
go run ./...
```

Notes:
- Folders numbered `01_...` through `27_...` are small, single-file examples demonstrating language features.
- Folders that contain their own `go.mod` are standalone modules; change into that folder before running or use `go run ./...` from inside the module.

## Repository Structure (high level)
- `01_Setup_first-program/` — Hello world and the first Go program.
- `02_variables_and_types/` to `27_methods_pointer_receiver/` — step-by-step language feature examples (variables, control flow, arrays, slices, maps, functions, errors, defer, pointers, structs, methods).
- `28_Go_lang_Module/` — example of a module layout using `cmd/` and `internal/` packages. Run with:

```bash
cd 28_Go_lang_Module
go run ./cmd/app
```

- `29_Http_module,_Working_with_JSON_unmarshal_and_http_client/` — HTTP server/client and JSON encoding/decoding examples. Subfolders include `01_basic_http_module`, `02_http_multiple_routes`, `03_json_encoder`, `04_json_decoder_decode_detail`, `05_http_get`, `06_reading_response_body`, `07_json_unmarshal_into_struct`, `08_external_api`, and a small `pokiApi` example.

Example (run one of the http samples):
```bash
cd "29_Http_module,_Working_with_JSON_unmarshal_and_http_client"/01_basic_http_module
go run main.go
```

- `30_GoLang_Notes_CRUD/` — small CRUD-style API demonstrating project layout with `cmd/api`, `internal/{config,db,notes,server}`. To run the API:

```bash
cd 30_GoLang_Notes_CRUD/cmd/api
go run main.go
```

Check `internal/config` for configuration details (port, DB settings). The example uses a simple MongoDB integration in `internal/db`.

- `31_Go_Lang_AUTH/` — demonstrates authentication patterns (JWT), middleware, user handlers, and a small API layout. To run the auth API:

```bash
cd 31_Go_Lang_AUTH/cmd/api
go run main.go
```

See `internal/auth`, `internal/middleware`, and `internal/user` for implementation details.

## Development workflow
- Format code: `gofmt -w .` or `go fmt ./...`.
- Vet code: `go vet ./...`.
- Run static checks (if you have `golangci-lint`): `golangci-lint run`.
- Run tests: `go test ./...` (this repo currently contains example code; tests may be limited).

## Committing and pushing changes
After updating files, standard Git workflow:
```bash
git add README.md
git commit -m "docs: update README with detailed usage and structure"
git push
```

If `git push` requires credentials, configure your credential manager or SSH keys as appropriate for your remote provider.

## Contributing
- Feel free to add examples, improve explanations, or add tests. Open a PR or issue with a short description of the change.

## Notes
- This repository is intended for learning and demonstration; it is not production-ready code.
- Check module `go.mod` files inside module folders for specific dependency versions.

---

If you'd like, I can also:
- run the commit and push the updated README now
- add a Table of Contents for quick navigation
- create a CONTRIBUTING.md or LICENSE file

Tell me which of the above you'd like me to do next.