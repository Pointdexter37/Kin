# Kin
Kin is a beginner-friendly Go CLI for saving, organizing, searching, editing,
previewing, and optionally executing terminal commands. It uses local SQLite
storage and is intentionally built as a small learning project.

## Features and download
Kin supports persistence, repeatable tags, listing, substring search,
interactive selection, editing, removal, safe previews, and opt-in execution.
It includes tests, `go vet`, GitHub Actions CI, and GoReleaser releases.
Download archives at <https://github.com/Pointdexter37/Kin/releases>.
Use `.zip` on Windows and `.tar.gz` on Linux/macOS; extract and run `kin`.
Linux/macOS users may need `chmod +x kin`.

## Quick start
```powershell
kin init
kin add "git status" --tag git --tag common
kin add "docker ps" --tag docker
kin list
kin search git
kin run 1
```
Kin creates `kin.db` in its current working directory. Back up that file to
preserve or move snippets.

## Commands
```text
kin init
kin add <command> [--tag <tag>]
kin list
kin search <text> [--interactive]
kin edit <id> <new command> [--tag <tag>]
kin remove <id>
kin run <id> [--execute]
```
Examples:
```powershell
kin add "git log --oneline" --tag git --tag history
kin search git --interactive
kin edit 1 "git status --short" --tag git
kin remove 2
```
Tags may be repeated. Search checks command text and tags.

## Safety
`kin run 1` only prints a preview. `kin run 1 --execute` asks for `y` or `yes`
before starting the command. Review commands carefully: execution uses the
operating system shell, so deletion, permissions, downloads, credentials,
and production commands remain your responsibility.

## Placeholders
The reusable `internal/placeholder` package supports `${NAME}` and
`${NAME:=default}`. Missing required values return an error. Expansion is
tested as a library feature but is not connected to `run` yet.

## Build from source
Install Go 1.25 or newer:
```powershell
git clone https://github.com/Pointdexter37/Kin.git
cd Kin
go test ./...
go vet ./...
go build -o kin.exe
```
On Linux/macOS, use `go build -o kin`.

## Project structure
```text
main.go                  Application entry point
cmd/                     Cobra CLI commands and output
internal/snippet/        SQLite storage and model
internal/placeholder/    Placeholder expansion
.github/workflows/       CI and release workflows
.goreleaser.yml          Release build configuration
```
The CLI layer handles arguments and presentation; internal packages own
storage and domain behavior, keeping future changes isolated.

## Development and releases
Run locally with `go run . list`. Before a pull request, run
`gofmt -w .`, `go test ./...`, `go vet ./...`, and `git diff --check`.
CI repeats tests and vet on pushes to `main` and pull requests.
To publish a release:
```powershell
git tag v0.1.0
git push origin v0.1.0
```
GoReleaser builds Windows, Linux, and macOS archives with checksums and
publishes them on GitHub Releases.

## Roadmap and license
Planned work includes wiring placeholders into `run`, stronger fuzzy search,
import/export, tag filtering, usage tracking, shell selection, configuration,
and more cross-platform execution tests. No license has been selected yet;
add one before redistributing Kin as open source.
