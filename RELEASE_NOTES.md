# NOTHING (TORRENT) — Initial Public Release 🔴

```text
 █▄ █ █▀█ ▀█▀ █ █ █ █▄ █ █▀▀   ( ▀█▀ █▀█ █▀█ █▀█ █▀▀ █▄ █ ▀█▀ )
 █ ▀█ █▄█  █  █▀█ █ █ ▀█ █▄█    █  █▄█ █▀▄ █▀▄ ██▄ █ ▀█  █  
```

> **Minimalist BitTorrent TUI Client inspired by the Nothing OS aesthetic.**  
> *Available as standalone native binaries for Windows, Linux, and macOS.*

---

### ✨ Highlights & Features

- **🔴 Nothing OS Visual Aesthetic:**
  - High-contrast monochrome interface with dot-matrix styling.
  - Iconic Nothing Red (`#D71920`) indicator dots and status glyphs.
  - Segmented dot progress bars (`■■■■■······`) and real-time Unicode Braille speed graphs.
- **📁 Interactive Pre-Download File Selector:**
  - View torrent metadata before anything is written to disk.
  - Navigate with arrows (`↑`/`↓`), toggle individual files with `[Space]`, select all (`a`), deselect all (`n`), invert (`i`).
- **🥷 Stealth / Cold-Router Mode (Tracker Spoofing):**
  - Keeps tracker ratio safely balanced around **`0.80x – 0.90x`** with realistic speed jitter.
  - Client emulation for `qBittorrent/4.6.5`, `Transmission/4.0.5`, and `Deluge/2.1.1` to pass strict tracker whitelists.
  - Physical chunk uploading disabled (`disable_seed: true`) to keep your router cool and your bandwidth free.
- **⌨️ Layout-Agnostic Controls (Russian & English):**
  - All hotkeys work seamlessly regardless of active keyboard layout (`й` = `q`, `ы` = `s`, `з` = `p`, `щ` = `o`, etc.).
- **⚙️ In-App Settings Screen:**
  - Press `[s]` anywhere to adjust target ratio, cycle client profiles, and toggle stealth modes.
- **🐧🪟 Deep System Integration:**
  - **Linux:** Automatic desktop launcher with terminal detection (`kitty`, `alacritty`, `foot`, `ghostty`, etc.) and `xdg-mime` associations for `application/x-bittorrent` and `x-scheme-handler/magnet`.
  - **Windows:** Zero-admin registry integration via `nothing-torrent.exe --install` for automatic handling of magnet links in Chrome, Edge, and Firefox.

---

### 📦 Download Assets

| Platform | File | Architecture |
|---|---|---|
| **Windows** | `nothing-torrent-windows-amd64.exe` | 64-bit (x86_64) |
| **Linux** | `nothing-torrent-linux-amd64` | 64-bit (x86_64) |
| **Linux (ARM)** | `nothing-torrent-linux-arm64` | 64-bit ARM (Raspberry Pi, aarch64) |
| **macOS (Apple Silicon)** | `nothing-torrent-darwin-arm64` | Apple Silicon (M1/M2/M3/M4) |
| **macOS (Intel)** | `nothing-torrent-darwin-amd64` | 64-bit Intel |

---

### 📜 License
Licensed under the **GNU General Public License v3.0 (GPL-3.0)**. Copyleft ensures the code stays free and open-source forever.
