# NOTHING (TORRENT) 🔴

> Minimalist BitTorrent TUI client inspired by the **Nothing OS** aesthetic.
> Built with Go, Bubble Tea, and pure BitTorrent internals.

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

[p] Pause/Resume   [s] Toggle Stealth   [o] Open Folder   [q] Quit
```

---

## ✨ Features

- **Nothing OS Aesthetic:**
  - High-contrast monochrome palette (`#0A0A0A` black, `#F2F2F2` white, `#737373` muted).
  - Signature Nothing Red (`#D71920`) accent and indicator glyphs (`●`, `○`, `›`).
  - Dot-matrix segmented progress bars and Unicode Braille speed graphs.
- **Interactive File Selection:**
  - Pick which files to download **before** chunks start writing to disk.
  - Arrow navigation (`↑`/`↓` or `k`/`j`), toggle with `[Space]`, select all (`a`), deselect all (`n`), invert (`i`).
- **Stealth / Cold-Router Mode (Tracker Spoofing):**
  - Keeps ratio safely balanced around **`0.80 – 0.90`** on private trackers.
  - Emulates client fingerprints (default: `qBittorrent/4.6.5` with `-qB4650-...` peer IDs).
  - Modifies tracker announce metrics with realistic speed jitter without burning upload bandwidth or heating up the router.
- **Deep Linux & Desktop Integration:**
  - Auto-registers `.desktop` entry in `~/.local/share/applications/`.
  - Installs custom Nothing SVG app icon in `~/.local/share/icons/`.
  - Automatically hooks into `application/x-bittorrent` and `x-scheme-handler/magnet` via `xdg-mime`.
  - Clicking any torrent file in file managers (Dolphin, Nautilus, Thunar) or magnet link in browsers (Chrome, Firefox, Brave) instantly pops open your favorite terminal (`kitty`, `alacritty`, `ghostty`, etc.) with NothingTorrent running.

---

## 🚀 Installation & Quick Start

### Build from Source (Arch Linux / Linux)
Requirements: `go` 1.22+

```bash
git clone https://github.com/nothing-torrent/nothing-torrent.git
cd nothing-torrent

# Build binary
make build

# Install binary & register Linux MIME associations
./nothing-torrent --install
```

### Usage

```bash
# Open magnet link directly
nothing-torrent "magnet:?xt=urn:btih:..."

# Open .torrent file directly
nothing-torrent ~/Downloads/movie.torrent

# Launch interactive mode (pastes prompt if no args)
nothing-torrent

# Re-register desktop integration
nothing-torrent --install

# Remove desktop integration
nothing-torrent --uninstall
```

---

## ⌨️ Controls

### File Selector Screen
| Key | Action |
|---|---|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Space` | Toggle download checkbox for selected file |
| `a` | Select all files |
| `n` | Deselect all files |
| `i` | Invert selection |
| `Enter` | Confirm selection and start downloading |
| `q` / `Esc` | Quit / Cancel |

### Live Dashboard Screen
| Key | Action |
|---|---|
| `p` | Pause / Resume download |
| `s` | Toggle Stealth Ratio Guard ON / OFF |
| `o` | Open download directory in system file manager |
| `q` / `Ctrl+C` | Gracefully stop and exit |

---

## ⚙️ Configuration

Configuration is automatically saved to `~/.config/nothing-torrent/config.json`:

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

- `target_ratio`: Desired target ratio reported to trackers (default: `0.85`).
- `disable_seed`: If `true`, does not send real chunks to peers (saves router CPU and upload bandwidth while reporting target ratio).
- `default_client`: Spoofed client profile (`qbittorrent`, `transmission`, `deluge`).

---

## 📄 License
MIT License.
