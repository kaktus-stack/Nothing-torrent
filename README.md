# NOTHING (TORRENT) 🔴

<div align="center">

```text
 █▄ █ █▀█ ▀█▀ █ █ █ █▄ █ █▀▀   ( ▀█▀ █▀█ █▀█ █▀█ █▀▀ █▄ █ ▀█▀ )
 █ ▀█ █▄█  █  █▀█ █ █ ▀█ █▄█    █  █▄█ █▀▄ █▀▄ ██▄ █ ▀█  █  
```

**Minimalist BitTorrent TUI Client inspired by the Nothing OS aesthetic.**  
*Минималистичный BitTorrent TUI клиент в эстетике Nothing OS для Linux.*

---

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-D71920?style=flat-square)](LICENSE)
[![Platform: Linux](https://img.shields.io/badge/Platform-Linux%20%2F%20Arch-black?style=flat-square&logo=archlinux)](https://archlinux.org)
[![UI: Bubble Tea](https://img.shields.io/badge/UI-Bubble%20Tea-F0F0F0?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Style: Nothing OS](https://img.shields.io/badge/Style-Nothing%20OS-D71920?style=flat-square)](https://nothing.tech)

[**[ 🇬🇧 English Documentation ](#-english)**] &nbsp;&nbsp;|&nbsp;&nbsp; [**[ 🇷🇺 Документация на русском ](#-русский)**]

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

<a name="-english"></a>
# 🇬🇧 English

## ⚡ Overview

**NothingTorrent** is an ultra-fast, single-binary BitTorrent client designed for power users who appreciate terminal minimalism, clean aesthetics, and complete privacy control. It provides a tactile TUI experience with instant file selection before downloading, client emulation for tracker compatibility, cold-hardware bandwidth protection, and seamless integration into Linux desktop environments.

### 📊 Feature Matrix

| Feature | NothingTorrent | Traditional GUI (qBittorrent) | Basic CLI (rtorrent/aria2) |
|---|:---:|:---:|:---:|
| **Memory Footprint** | **~25 MB** | ~150–400 MB | ~15–30 MB |
| **Interactive File Selector** | **Yes (pre-download)** | Yes (heavy popups) | No (manual config) |
| **Nothing OS Aesthetic** | **Yes (dot-matrix/red glyphs)** | No | No |
| **Tracker Client Emulation** | **Built-in (qB/Transmission)** | Third-party hacks | Manual passkeys |
| **Cold-Router / Ratio Keeper** | **Built-in (0.80–0.90x)** | No | External scripts |
| **Linux MIME / Browser Hook** | **Automatic (one-click)** | Desktop-only | Manual xdg setup |
| **Layout-Agnostic Hotkeys** | **Yes (QWERTY & Cyrillic)** | Varies | Mostly English only |

---

## 🎯 Key Capabilities

- **Interactive File Picker:** Preview torrent metadata and toggle specific files (`Space`) before a single piece is written to disk.
- **Stealth Ratio Guard:** Automatically reports believable upload ratios (`0.80x – 0.90x`) to private trackers with randomized jitter.
- **Cold-Router Mode:** Does not burn CPU cycles or upload bandwidth (`disable_seed: true`), keeping your router and hardware cool.
- **Adaptive Terminal UI:** Scales effortlessly from narrow 50-column split panes up to ultrawide 4K terminals.
- **Cyrillic Keyboard Support:** Works identically on Russian and English layouts without needing to switch layouts in your system.
- **Desktop Auto-Integration:** Intercepts magnet links in Firefox, Chrome, Brave, and `.torrent` files in Dolphin, Nautilus, Thunar, Wofi, and Rofi.

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

## 🛠️ Quick Installation

```bash
# Clone the repository
git clone https://github.com/kaktus-stack/Nothing-torent.git
cd Nothing-torent

# Build single static binary (22 MB)
make build

# Install binary & register Linux desktop associations
make install
```

---

<br>

---

<a name="-русский"></a>
# 🇷🇺 Русский

## ⚡ О проекте

**NothingTorrent** — это легковесный терминальный BitTorrent-клиент, написанный на Go в минималистичной эстетике **Nothing OS**. Проект создан для тех, кто устал от тяжёлых монструозных GUI-клиентов, жрущих сотни мегабайт оперативной памяти, но хочет иметь удобный интерактивный выбор файлов стрелочками, красивую псевдографику и бесшовную интеграцию в систему.

### 🌟 Главные фичи

1. **Эстетика Nothing OS:**
   - Высококонтрастная монохромная палитра (`#0A0A0A` черный, `#F2F2F2` белый, `#737373` приглушенный).
   - Фирменный акцентный красный Nothing Red (`#D71920`) и глифы (`●`, `○`, `›`).
   - Точечные сегментированные прогресс-бары и графики скорости на Unicode Braille.
2. **Интерактивный селектор файлов:**
   - Сначала подгружает метаданные из роя (с анимацией точечной матрицы `[ · · · ]`).
   - Позволяет стрелочками и пробелом выбрать только нужные серии сериала или отдельные файлы из раздачи.
3. **Стелс-режим «Холодный роутер» (Ratio Guard):**
   - Эмулирует популярные клиенты (`qBittorrent/4.6.5`, `Transmission/4.0.5`, `Deluge/2.1.1`) на уровне User-Agent и Peer-ID.
   - Поддерживает безопасный рейтинг на закрытых трекерах в районе **`0.80x – 0.90x`** с реалистичным джиттером.
   - Физическая раздача блоков отключена: интернет-канал свободен, роутер не перегревается.
4. **Независимость от раскладки:**
   - Все горячие клавиши работают одинаково как на английской, так и на русской раскладке (`й` = `q`, `ы` = `s`, `з` = `p`, `щ` = `o` и т.д.).
5. **Бесшовная интеграция в Linux / Arch:**
   - Автоматически находит твой эмулятор терминала (`kitty`, `alacritty`, `foot`, `ghostty`, `wezterm`, `gnome-terminal` и др.).
   - Перехватывает ссылки `magnet:` из браузеров и файлы `.torrent` из проводников и лаунчеров (Wofi, Rofi).

---

## ⚙️ Конфигурация (`config.json`)

Настройки автоматически создаются в `~/.config/nothing-torrent/config.json`:

```json
{
  "download_dir": "/home/username/Downloads",
  "default_client": "qbittorrent",
  "stealth_mode": true,
  "target_ratio": 0.85,
  "min_upload_kbps": 90,
  "max_upload_kbps": 220,
  "disable_seed": true
}
```

| Параметр | По умолчанию | Описание |
|---|---|---|
| `download_dir` | `~/Downloads` | Папка для сохранения скачанных файлов |
| `default_client` | `"qbittorrent"` | Профиль маскировки (`qbittorrent`, `transmission`, `deluge`) |
| `stealth_mode` | `true` | Включение автоматического поддержания рейтинга на трекерах |
| `target_ratio` | `0.85` | Целевой рейтинг, сообщаемый трекеру (0.80–0.90x) |
| `disable_seed` | `true` | Режим холодного роутера (нулевая реальная отдача блоков) |
| `min_upload_kbps` | `90` | Нижний порог симулируемой скорости отдачи в аннонсах |
| `max_upload_kbps` | `220` | Верхний порог симулируемой скорости отдачи в аннонсах |

---

## 📄 Лицензия

Распространяется под свободной лицензией [MIT](LICENSE).
