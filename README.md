# Kin
Kin is a beginner-friendly Go CLI for saving, organizing, searching, editing,
previewing, and optionally executing terminal commands. It uses local SQLite
storage and is intentionally built as a small learning project.

## Features and installation
Kin supports persistence, repeatable tags, listing, substring search,
interactive selection, editing, removal, safe previews, and opt-in execution.
It includes tests, `go vet`, GitHub Actions CI, and GoReleaser releases.

There are four ways to install Kin. The installer scripts and release archives
download published binaries, so Go is not required unless you choose the
`go install` method.

### Option 1: Windows PowerShell installer

Requirements:

- Windows 10 or newer
- PowerShell 5.1 or PowerShell 7
- Internet access

Run this command in PowerShell:

```powershell
irm https://raw.githubusercontent.com/Pointdexter37/Kin/main/scripts/install.ps1 | iex
```

The installer finds the latest GitHub release, downloads the Windows 64-bit
archive, verifies its SHA-256 checksum, installs `kin.exe` to `$HOME\bin`,
and adds that directory to the user PATH. Open a new PowerShell window after
installation:

```powershell
kin --help
kin init
```

To choose another installation directory, download the script and run it with
`-InstallDir`:

```powershell
irm https://raw.githubusercontent.com/Pointdexter37/Kin/main/scripts/install.ps1 `
  -OutFile "$env:TEMP\kin-install.ps1"
& "$env:TEMP\kin-install.ps1" -InstallDir "$HOME\AppData\Local\Kin"
```

The custom directory is also added to the user PATH. The Windows installer
currently targets `amd64` systems.

### Option 2: Linux and macOS installer

Requirements:

- Linux or macOS
- `curl` and `tar`
- `sha256sum` on Linux or `shasum` on macOS

Run:

```sh
curl -fsSL https://raw.githubusercontent.com/Pointdexter37/Kin/main/scripts/install.sh | sh
```

The installer detects the operating system and CPU architecture, downloads
the matching archive, verifies its SHA-256 checksum, and installs `kin` to
`$HOME/.local/bin`. It currently supports Linux and macOS on `amd64` and
`arm64`.

If `$HOME/.local/bin` is not already on your PATH, add it for the current
shell:

```sh
export PATH="$HOME/.local/bin:$PATH"
```

To make that change permanent for common shells:

```sh
printf '\nexport PATH="$HOME/.local/bin:$PATH"\n' >> "$HOME/.profile"
```

You can choose another installation directory with `KIN_INSTALL_DIR`:

```sh
KIN_INSTALL_DIR="$HOME/bin" \
  curl -fsSL https://raw.githubusercontent.com/Pointdexter37/Kin/main/scripts/install.sh | sh
```

Verify the installation:

```sh
kin --help
kin init
```

### Option 3: Install with Go

This option is useful for Go developers and requires Go 1.25 or newer:

```sh
go install github.com/Pointdexter37/kin@latest
```

Go places the executable in its Go binary directory, normally
`$HOME/go/bin` on Linux/macOS or `%USERPROFILE%\go\bin` on Windows. Ensure
that directory is on your PATH, then run:

```sh
kin --help
kin init
```

To see the directories configured for Go:

```sh
go env GOBIN GOPATH
```

### Option 4: Manual GitHub Release download

Manual downloads are available at
<https://github.com/Pointdexter37/Kin/releases>. Choose the archive matching
your operating system and CPU architecture:

| Platform | Archive |
| --- | --- |
| Windows 64-bit | `Kin_<version>_windows_amd64.zip` |
| Linux 64-bit | `Kin_<version>_linux_amd64.tar.gz` |
| Linux ARM64 | `Kin_<version>_linux_arm64.tar.gz` |
| macOS Intel | `Kin_<version>_darwin_amd64.tar.gz` |
| macOS Apple Silicon | `Kin_<version>_darwin_arm64.tar.gz` |

Each release also contains `checksums.txt`. Verify the archive before using
it. On Linux:

```sh
sha256sum Kin_<version>_<os>_<arch>.tar.gz
```

On macOS:

```sh
shasum -a 256 Kin_<version>_<os>_<arch>.tar.gz
```

On Windows PowerShell:

```powershell
Get-FileHash .\Kin_<version>_windows_amd64.zip -Algorithm SHA256
```

Compare the resulting hash with the matching line in `checksums.txt`. Extract
the archive and place the executable in a directory on your PATH.

Linux/macOS example:

```sh
mkdir -p "$HOME/.local/bin"
tar -xzf Kin_<version>_<os>_<arch>.tar.gz -C "$HOME/.local/bin"
chmod +x "$HOME/.local/bin/kin"
```

Windows PowerShell example:

```powershell
Expand-Archive .\Kin_<version>_windows_amd64.zip -DestinationPath "$HOME\bin"
```

Open a new terminal after updating PATH, then run `kin --help`.

### Updating and uninstalling

Run the same installer command again to update a binary installed by an
installer script. For a Go installation, run
`go install github.com/Pointdexter37/kin@latest` again.

The installers do not remove Kin's database. To remove the executable:

- Windows: delete `$HOME\bin\kin.exe` or the custom installation path.
- Linux/macOS: delete `$HOME/.local/bin/kin` or the custom installation path.
- Go installation: delete `kin` from the directory shown by
  `go env GOBIN GOPATH`.

Kin stores snippets in `kin.db` in the current working directory. Keep that
file if you want to preserve your data; deleting the executable does not
delete it.

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
and more cross-platform execution tests. Kin is available under the MIT
License; see [LICENSE](LICENSE).
