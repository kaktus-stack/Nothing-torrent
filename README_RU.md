<div align="center">

# NOTHING (TORRENT) 🔴

```text
 █▄ █ █▀█ ▀█▀ █ █ █ █▄ █ █▀▀   ( ▀█▀ █▀█ █▀█ █▀█ █▀▀ █▄ █ ▀█▀ )
 █ ▀█ █▄█  █  █▀█ █ █ ▀█ █▄█    █  █▄█ █▀▄ █▀▄ ██▄ █ ▀█  █  
```

**Минималистичный BitTorrent TUI клиент в эстетике Nothing OS для Linux и Windows.**

<br>

<!-- Панель переключения языка -->
<p>
  <a href="README.md"><img src="https://img.shields.io/badge/Switch_to-English-white?style=for-the-badge&logo=readme&logoColor=black" alt="English"></a>
  <a href="README_RU.md"><img src="https://img.shields.io/badge/Язык-Русский-D71920?style=for-the-badge&logo=readme&logoColor=white" alt="Русский"></a>
</p>

<!-- Кнопка перехода к релизу -->
<p>
  <a href="https://github.com/kaktus-stack/Nothing-torent/releases/latest">
    <img src="https://img.shields.io/badge/⬇️_СКАЧАТЬ_СВЕЖИЙ_РЕЛИЗ-v1.0.1-D71920?style=for-the-badge&logo=github&logoColor=white" height="38" alt="Скачать релиз">
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

## ⚡ О проекте

**NothingTorrent** — это сверхбыстрый консольный BitTorrent-клиент, написанный на Go в минималистичной стилистике **Nothing OS**. Проект создан для тех, кому надоели тяжёлые монструозные клиенты с веб-мордами, жрущие по 400 МБ оперативки, но при этом нужен комфортный интерактивный выбор файлов стрелочками, живая псевдографика и бесшовная системная интеграция.

### 📊 Сравнение возможностей

| Функция | NothingTorrent | Обычные GUI (qBittorrent) | Консольные CLI (rtorrent/aria2) |
|---|:---:|:---:|:---:|
| **Потребление RAM** | **~25 МБ** | ~150–400 МБ | ~15–30 МБ |
| **Выбор файлов до скачивания** | **Да (интерактивный список)** | Да (всплывающие окна) | Нет (ручная правка) |
| **Эстетика Nothing OS** | **Да (точечная матрица / глифы)** | Нет | Нет |
| **Маскировка под клиенты** | **Встроена (qB/Transmission)** | Сторонние патчи | Ручные ключи |
| **Стелс-рейтинг (холодный роутер)** | **Встроен (0.80–0.90x)** | Нет | Внешние скрипты |
| **Интеграция с проводником/браузером** | **Автоматическая (в 1 клик)** | Только рабочий стол | Ручной xdg-setup |
| **Независимость от раскладки** | **Да (работает на RU и EN)** | Зависит от ОС | В основном только EN |

---

## 🎯 Ключевые фичи

- **Интерактивный селектор файлов:** подгружает метаданные торрента из роя и позволяет стрелочками и пробелом (`Space`) выбрать только нужные серии сериала или отдельные файлы **до** начала записи на диск.
- **Стелс-режим «Холодный роутер» (Ratio Guard):** эмулирует `qBittorrent/4.6.5` или `Transmission/4.0.5`, поддерживая безопасный рейтинг на закрытых трекерах в диапазоне **`0.80x – 0.90x`** с рандомизированным джиттером.
- **Нулевая нагрузка на железо:** физическая отдача блоков в сеть отключена (`disable_seed: true`), твой интернет-канал свободен, а роутер не греется.
- **Адаптивный TUI-интерфейс:** красиво масштабируется от узких 50-колоночных тайлинговых сплитов в Hyprland/i3 до широких 4K экранов.
- **Поддержка русской раскладки:** все хоткеи работают одинаково, переключать язык на клавиатуре не нужно (`й` = `q`, `ы` = `s`, `з` = `p`, `щ` = `o` и т.д.).
- **Интеграция в Linux и Windows:** перехватывает клики по `magnet:` ссылкам в Firefox, Chrome, Edge, Brave и файлы `.torrent` в Dolphin, Nautilus, Thunar, Wofi, Rofi и Проводнике Windows.

---

## ⌨️ Управление и горячие клавиши

| Клавиша (RU) | Клавиша (EN) | Действие | Где работает |
|---|---|---|---|
| `л` / `↑` | `k` / `↑` | Переместить курсор вверх | Выбор файлов / Настройки |
| `о` / `↓` | `j` / `↓` | Переместить курсор вниз | Выбор файлов / Настройки |
| `Пробел` | `Space` | Включить/выключить файл / переключить опцию | Выбор файлов / Настройки |
| `ф` | `a` | Выбрать все файлы | Выбор файлов |
| `т` | `n` | Снять выбор со всех файлов | Выбор файлов |
| `ш` | `i` | Инвертировать выбор | Выбор файлов |
| `Enter` | `Enter` | Запустить закачку / применить настройку | Выбор файлов / Настройки |
| `з` | `p` | Пауза / Возобновление закачки | Дашборд |
| `ы` | `s` | Открыть / Закрыть меню настроек | Дашборд / Везде |
| `щ` | `o` | Открыть папку со скачанными файлами | Дашборд |
| `й` / `Esc` | `q` / `Esc` | Выход из программы | Везде |

---

## 🛠️ Установка и запуск

### 🪟 Windows (Windows Terminal / PowerShell / CMD)
1. Скачай `nothing-torrent-windows-amd64.exe` со страницы **[Релизов](https://github.com/kaktus-stack/Nothing-torent/releases/latest)**.
2. Запусти в терминале или кликни дважды:
```powershell
# Регистрация magnet: ссылок и .torrent файлов в реестре Windows (без прав админа)
.\nothing-torrent-windows-amd64.exe --install

# Запуск с передачей ссылки или торрент-файла
.\nothing-torrent-windows-amd64.exe "magnet:?xt=urn:btih:..."
```

### 🐧 Linux (Arch / Ubuntu / Debian / Fedora)
```bash
git clone https://github.com/kaktus-stack/Nothing-torent.git
cd Nothing-torent

# Сборка статичного бинарника (21 МБ)
make build

# Установка бинарника и регистрация в Linux десктопе
make install
```

---

## ⚙️ Конфигурация (`config.json`)

Автоматически сохраняется в `~/.config/nothing-torrent/config.json`:

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
| `download_dir` | `~/Downloads` | Директория для сохранения файлов |
| `default_client` | `"qbittorrent"` | Профиль маскировки (`qbittorrent`, `transmission`, `deluge`) |
| `stealth_mode` | `true` | Включение автоматического поддержания рейтинга на трекерах |
| `target_ratio` | `0.85` | Целевой рейтинг, сообщаемый трекеру (диапазон 0.80–0.90x) |
| `disable_seed` | `true` | Режим холодного роутера (нулевая физическая отдача блоков) |
| `min_upload_kbps` | `90` | Нижний порог симулируемой скорости отдачи |
| `max_upload_kbps` | `220` | Верхний порог симулируемой скорости отдачи |

---

## 📄 Лицензия

Проект распространяется под свободной копилефт-лицензией **GNU General Public License v3.0 (GPL-3.0)**. Полный текст доступен в файле [LICENSE](LICENSE). Копилефт гарантирует, что код навсегда останется открытым и свободным для сообщества.
