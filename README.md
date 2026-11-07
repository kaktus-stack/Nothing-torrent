<div align="center">

# NOTHING (TORRENT) 🔴

```text
 █▄ █ █▀█ ▀█▀ █ █ █ █▄ █ █▀▀   ( ▀█▀ █▀█ █▀█ █▀█ █▀▀ █▄ █ ▀█▀ )
 █ ▀█ █▄█  █  █▀█ █ █ ▀█ █▄█    █  █▄█ █▀▄ █▀▄ ██▄ █ ▀█  █  
```

**Minimalist BitTorrent TUI Client inspired by the Nothing OS aesthetic.**

<br>

<!-- Language Switcher Bar -->
<p>
  <a href="README.md"><img src="https://img.shields.io/badge/Language-English-D71920?style=for-the-badge&logo=readme&logoColor=white" alt="English"></a>
  <a href="README_RU.md"><img src="https://img.shields.io/badge/Переключить_на-Русский-white?style=for-the-badge&logo=readme&logoColor=black" alt="Русский"></a>
</p>

<!-- Big Download Button -->
<p>
  <a href="https://github.com/kaktus-stack/Nothing-torent/releases/latest">
    <img src="https://img.shields.io/badge/⬇️_DOWNLOAD_LATEST_RELEASE-v1.0.1-D71920?style=for-the-badge&logo=github&logoColor=white" height="38" alt="Download Release">
  </a>
</p>

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-D71920?style=flat-square)](LICENSE)
[![Platform: Linux & Windows](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows-black?style=flat-square)](https://github.com/kaktus-stack/Nothing-torent/releases)
[![UI: Bubble Tea](https://img.shields.io/badge/UI-Bubble%20Tea-F0F0F0?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Style: Nothing OS](https://img.shields.io/badge/Style-Nothing%20OS-D71920?style=flat-square)](https://nothing.tech)

---

```text
● NOTHING (TORRENT)   [ DOWNLOADING ]   [● RATIO 0.86x · qBittorrent/4.6.5]

Ubuntu 24.04 LTS Desktop (amd64)

[■■■■■■■■■■■■■■■■■■········································]  36.8%

┌ DOWN SPEED: 18.4 MB/s ┐ ┌ STEALTH UP: 194 KB/s (Stealth) ┐ ┌ ETA:   00:04:12 ┐
└ DOWNLOADED: 1.8 / 5.0 ┘ └ REPORTED:   1.5 GB (Simulated) ┘ └ PEERS: 38       ┘

SPEED GRAPH  ▅▆▇██▇▆▅▄▃▂ ▂▃▄▅▆▇

FILES PROGRESS
● 01 - ubuntu-24.04-desktop-amd64.iso     ■■■■■■■····  54.2% (5.0 GB)
✕ 02 - release-notes.txt                  [SKIPPED]

[p] Pause/Resume   [s] Settings   [o] Open Folder   [q] Quit
```

</div>

---

## ⚡ Overview

**NothingTorrent** is an ultra-fast, single-binary BitTorrent client designed for power users who appreciate terminal minimalism, clean aesthetics, and complete privacy control. It provides a tactile TUI experience with instant file selection before downloading, client emulation for tracker compatibility, cold-hardware bandwidth protection, and seamless integration into Linux and Windows environments.

### 📊 Feature Matrix

| Feature | NothingTorrent | Traditional GUI (qBittorrent) | Basic CLI (rtorrent/aria2) |
|---|:---:|:---:|:---:|
| **Memory Footprint** | **~25 MB** | ~150–400 MB | ~15–30 MB |
| **Interactive File Selector** | **Yes (pre-download)** | Yes (heavy popups) | No (manual config) |
| **Nothing OS Aesthetic** | **Yes (dot-matrix/red glyphs)** | No | No |
| **Tracker Client Emulation** | **Built-in (qB/Transmission)** | Third-party hacks | Manual passkeys |
| **Cold-Router / Ratio Keeper** | **Built-in (0.80–0.90x)** | No | External scripts |
| **Desktop / Browser Integration** | **Automatic (one-click)** | Desktop-only | Manual setup |
| **Layout-Agnostic Hotkeys** | **Yes (QWERTY & Cyrillic)** | Varies | Mostly English only |

---

## 🎯 Key Capabilities

- **Interactive File Picker:** Preview torrent metadata and toggle specific files (`Space`) before a single piece is written to disk.
- **Stealth Ratio Guard:** Automatically reports believable upload ratios (`0.80x – 0.90x`) to private trackers with randomized jitter.
- **Cold-Router Mode:** Does not burn CPU cycles or upload bandwidth (`disable_seed: true`), keeping your router and hardware cool.
- **Adaptive Terminal UI:** Scales effortlessly from narrow 50-column split panes up to ultrawide 4K terminals.
- **Cyrillic Keyboard Support:** Works identically on Russian and English layouts without needing to switch layouts in your system.
- **Desktop Auto-Integration:** Intercepts magnet links in Firefox, Chrome, Brave, and `.torrent` files in Dolphin, Nautilus, Thunar, Wofi, Rofi, and Windows Explorer.

---

## ⌨️ Controls & Keybindings

| Key (EN) | Key (RU) | Action | Context |
|---|---|---|---|
| `↑` / `k` | `л` | Move cursor up | File Selector / Settings |
| `↓` / `j` | `о` | Move cursor down | File Selector / Settings |
| `Space` | `Пробел` | Toggle file checkbox / toggle setting | File Selector / Settings |
| `a` | `ф` | Select all files | File Selector |
| `n` | `т` | Deselect all files | File Selector |
| `i` | `ш` | Invert file selection | File Selector |
| `Enter` | `Enter` | Start download / Apply setting | File Selector / Settings |
| `p` | `з` | Pause / Resume active download | Dashboard |
| `s` | `ы` | Open / Close interactive Settings | Dashboard / Global |
| `o` | `щ` | Open download directory in file manager | Dashboard |
| `q` / `Esc` | `й` / `Esc` | Gracefully quit application | Global |

---

## 🛠️ Quick Installation & Setup

### 🪟 Windows (Windows Terminal / PowerShell / CMD)
1. Download `nothing-torrent-windows-amd64.exe` from **[Releases](https://github.com/kaktus-stack/Nothing-torent/releases/latest)**.
2. Run in terminal or double click:
```powershell
# Optional: Register magnet: protocol and .torrent association in Windows Registry (no admin required)
.\nothing-torrent-windows-amd64.exe --install

# Launch with magnet link or file
.\nothing-torrent-windows-amd64.exe "magnet:?xt=urn:btih:..."
```

### 🐧 Linux (Arch / Ubuntu / Debian / Fedora)
```bash
git clone https://github.com/kaktus-stack/Nothing-torent.git
cd Nothing-torent

# Build single static binary (21 MB)
make build

# Install binary & register Linux desktop associations
make install
```

---

## ⚙️ Configuration (`config.json`)

Saved automatically to `~/.config/nothing-torrent/config.json`:

```json
{
  "download_dir": "/home/user/Downloads",
  "default_client": "qbittorrent",
  "stealth_mode": true,
  "target_ratio": 0.85,
  "min_upload_kbps": 90,
  "max_upload_kbps": 220,
  "disable_seed": true
}
```

---

## 📄 License

Licensed under the **GNU General Public License v3.0 (GPL-3.0)**. See [LICENSE](LICENSE) for details. Copyleft guarantees this software remains free and open-source forever.
