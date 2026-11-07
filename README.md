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

[p] Pause/Resume   [s] Settings   [o] Open Folder   [q] Quit
```

---

## ✨ Features

- **Nothing OS Aesthetic:**
  - High-contrast monochrome palette (`#0A0A0A` black, `#F2F2F2` white, `#737373` muted).
  - Signature Nothing Red (`#D71920`) accent and indicator glyphs (`●`, `○`, `›`).
  - Dot-matrix segmented progress bars and Unicode Braille speed graphs.
- **Responsive Layout:**
  - Automatically adapts to terminal width (from small split-panes <60 cols to ultra-wide displays).
  - Clean card grids on wide terminals; compact stacked telemetry on small screens.
- **Layout-Agnostic Controls (Russian / Cyrillic Support):**
  - All hotkeys work seamlessly regardless of keyboard layout (`й` = `q`, `ы` = `s`, `з` = `p`, `щ` = `o`, `ф` = `a`, `т` = `n`, `л`/`о` = `k`/`j`).
- **Interactive File Selection:**
  - Pick which files to download **before** chunks start writing to disk.
  - Arrow navigation (`↑`/`↓` or `k`/`j`), toggle with `[Space]`, select all (`a`), deselect all (`n`), invert (`i`).
- **In-App Settings Screen:**
  - Press `[s]` anywhere to open the live settings menu.
  - Adjust target ratio, cycle spoofed client profiles, toggle cold-router mode and simulated upload speed presets.
- **Stealth / Cold-Router Mode (Tracker Spoofing):**
  - Keeps ratio safely balanced around **`0.80 – 0.90`** on private trackers.
  - Emulates client fingerprints (default: `qBittorrent/4.6.5` with `-qB4650-...` peer IDs).
  - Modifies tracker announce metrics with realistic speed jitter without burning upload bandwidth or heating up the router.
- **Universal Linux & Desktop Integration:**
  - Auto-detects terminal emulator (`kitty`, `alacritty`, `foot`, `ghostty`, `wezterm`, etc.).
  - Runs with Wofi, Rofi, Dolphin, Nautilus, Thunar, Chrome, Firefox.
  - Sets MIME associations via `xdg-mime` for `application/x-bittorrent` and `x-scheme-handler/magnet`.

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

> 💡 *All keybindings work on both English and Russian keyboard layouts.*

### File Selector Screen
| Key (EN / RU) | Action |
|---|---|
| `↑` / `k` (л) | Move cursor up |
| `↓` / `j` (о) | Move cursor down |
| `Space` | Toggle download checkbox for selected file |
| `a` (ф) | Select all files |
| `n` (т) | Deselect all files |
| `i` (ш) | Invert selection |
| `Enter` | Confirm selection and start downloading |
| `q` (й) / `Esc` | Quit / Cancel |

### Live Dashboard Screen
| Key (EN / RU) | Action |
|---|---|
| `p` (з) | Pause / Resume download |
| `s` (ы) | Open Settings menu |
| `o` (щ) | Open download directory in system file manager |
| `q` (й) / `Ctrl+C` | Gracefully stop and exit |

### Settings Menu
| Key (EN / RU) | Action |
|---|---|
| `↑` / `↓` / `k` / `j` | Navigate setting fields |
| `←` / `→` / `h` / `l` / `Space` | Adjust / cycle value |
| `Esc` / `s` (ы) | Save changes and return to dashboard |
| `q` (й) | Quit client |

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

---

## 📄 License
MIT License.
