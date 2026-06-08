<div align="center">

```
  ███╗   ███╗██╗  ██╗ ██████╗██████╗
  ████╗ ████║██║ ██╔╝██╔════╝██╔══██╗
  ██╔████╔██║█████╔╝ ██║     ██║  ██║
  ██║╚██╔╝██║██╔═██╗ ██║     ██║  ██║
  ██║ ╚═╝ ██║██║  ██╗╚██████╗██████╔╝
  ╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝╚═════╝
```

**Make a directory and jump into it — instantly.**

[![Release](https://img.shields.io/github/v/release/suprabhaat/mkcd?style=flat-square&color=00ff88&labelColor=0a0a0a)](https://github.com/suprabhaat/mkcd/releases)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&labelColor=0a0a0a&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-white?style=flat-square&labelColor=0a0a0a)](LICENSE)
![Platforms](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-555?style=flat-square&labelColor=0a0a0a)

</div>

---

```bash
$ mkcd my-new-project
# Directory created ✔
# → now you're inside  ~/code/my-new-project
```

> One command. No more `mkdir foo && cd foo`.

---

## ✦ How it works

`mkcd` is a tiny Go binary + a shell function working together.

```
you type:  mkcd hello
              │
              ▼
    ┌─────────────────────┐
    │   mkcd-bin (Go)     │
    │  • validates name   │
    │  • mkdir -p         │
    │  • prints abs path  │◄─── stdout only
    └─────────────────────┘
              │
              ▼
    shell function captures path
              │
              ▼
         cd  /abs/path       ← happens in YOUR shell
```

The binary can't `cd` for you (subprocesses can't change the parent shell's directory). The shell function bridges the gap.

---

## ⬇ Installation

### 1 — Download the binary

Go to [**Releases**](https://github.com/suprabhaat/mkcd/releases) and grab the binary for your platform:

| Platform | Architecture | File |
|---|---|---|
| Linux | x86 32-bit | `mkcd-bin-linux-x86` |
| Linux | amd64 64-bit | `mkcd-bin-linux-amd64` |
| Linux | arm32 | `mkcd-bin-linux-arm32` |
| Linux | arm64 | `mkcd-bin-linux-arm64` |
| macOS | amd64 (Intel) | `mkcd-bin-macos-amd64` |
| macOS | arm64 (Apple Silicon) | `mkcd-bin-macos-arm64` |
| Windows | x86 32-bit | `mkcd-bin-windows-x86.exe` |
| Windows | amd64 64-bit | `mkcd-bin-windows-amd64.exe` |
| Windows | arm32 | `mkcd-bin-windows-arm32.exe` |
| Windows | arm64 | `mkcd-bin-windows-arm64.exe` |

---

### 2 — Linux / macOS

```bash
# Replace with your downloaded filename
mv mkcd-bin-linux-amd64 mkcd-bin
chmod +x mkcd-bin
sudo mv mkcd-bin /usr/local/bin/mkcd-bin
```

Then add the shell function to your `~/.bashrc` or `~/.zshrc`:

```bash
mkcd() {
  local path
  path=$(mkcd-bin "$@") || return 1
  cd "$path" || return 1
}
```

Reload your shell:

```bash
source ~/.bashrc   # bash
source ~/.zshrc    # zsh
```

---

### 3 — Windows (PowerShell)

Move the `.exe` to a folder on your `PATH` (e.g. `C:\Users\you\bin\`):

```powershell
Move-Item mkcd-bin-windows-amd64.exe C:\Users\you\bin\mkcd-bin.exe
```

Add the wrapper function to your PowerShell profile:

```powershell
# Open your profile
notepad $PROFILE
```

Paste this and save:

```powershell
function mkcd {
  $path = mkcd-bin @args
  if ($LASTEXITCODE -ne 0) { return }
  Set-Location $path
}
```

Restart PowerShell.

---

### 4 — Termux / Android (ARM64)

```bash
mv mkcd-bin-linux-arm64 mkcd-bin
chmod +x mkcd-bin
mv mkcd-bin $PREFIX/bin/mkcd-bin
```

Add to `~/.bashrc`:

```bash
mkcd() {
  local path
  path=$(mkcd-bin "$@") || return 1
  cd "$path" || return 1
}
```

```bash
source ~/.bashrc
```

---

## ⚡ Usage

```bash
mkcd my-project           # create and enter a directory
mkcd path/to/nested/dir   # creates all intermediate dirs
mkcd "dir with spaces"    # spaces handled automatically
mkcd existing-dir         # already exists? no problem, just cd into it
```

**Error cases handled:**

```bash
mkcd                      # → Usage: mkcd <directory>
mkcd $'\x01bad'           # → control characters are not allowed
mkcd existing-file        # → 'existing-file' exists but is not a directory
```

---

## 🔨 Build from source

```bash
git clone https://github.com/suprabhaat/mkcd
cd mkcd
go build -o mkcd-bin ./cmd/mkcd/src
sudo mv mkcd-bin /usr/local/bin/mkcd-bin
```

Cross-compile for any target:

```bash
GOOS=linux GOARCH=arm64 go build -o mkcd-bin-linux-arm64 ./cmd/mkcd/src
```

---

## 📁 Project structure

```
mkcd/
├── cmd/
│   └── mkcd/
│       └── src/
│           └── main.go        ← the binary
├── .github/
│   └── workflows/
│       └── release.yml        ← auto-release CI
├── go.mod
└── README.md
```

---

## 🚀 Releases

Releases are built automatically by GitHub Actions on every `v*` tag push.

```bash
git tag v1.0.0
git push origin v1.0.0
# → CI builds 11 binaries and publishes a GitHub Release
```

---

<div align="center">

MIT License · built with Go · works everywhere

</div>
