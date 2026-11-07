package desktop

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const LauncherScriptContent = `#!/usr/bin/env bash
# NothingTorrent Universal Terminal Launcher

# Find binary
BIN="$HOME/.local/bin/nothing-torrent"
if [ ! -x "$BIN" ]; then
    BIN="$(command -v nothing-torrent 2>/dev/null)"
fi

if [ -z "$BIN" ] || [ ! -x "$BIN" ]; then
    notify-send "NothingTorrent" "Error: nothing-torrent binary not found in ~/.local/bin or PATH" 2>/dev/null
    exit 1
fi

# If already inside an interactive terminal, run directly
if [ -t 0 ] && [ -t 1 ]; then
    exec "$BIN" "$@"
fi

# Detect installed terminal emulator
CANDIDATES=(
    "$TERMINAL"
    kitty
    ghostty
    alacritty
    foot
    wezterm
    konsole
    gnome-terminal
    xfce4-terminal
    termite
    urxvt
    rxvt
    x-terminal-emulator
    xterm
)

for term in "${CANDIDATES[@]}"; do
    [ -z "$term" ] && continue
    TERM_BIN="$(command -v "$term" 2>/dev/null)"
    [ -z "$TERM_BIN" ] && continue

    TERM_NAME="$(basename "$TERM_BIN")"
    case "$TERM_NAME" in
        kitty)
            exec "$TERM_BIN" --title "NothingTorrent" "$BIN" "$@"
            ;;
        alacritty)
            exec "$TERM_BIN" --title "NothingTorrent" -e "$BIN" "$@"
            ;;
        foot)
            exec "$TERM_BIN" --title "NothingTorrent" "$BIN" "$@"
            ;;
        ghostty)
            exec "$TERM_BIN" --title="NothingTorrent" -e "$BIN" "$@"
            ;;
        wezterm)
            exec "$TERM_BIN" start -- "$BIN" "$@"
            ;;
        gnome-terminal)
            exec "$TERM_BIN" -- "$BIN" "$@"
            ;;
        xfce4-terminal|konsole|xterm|rxvt|urxvt|termite|x-terminal-emulator)
            exec "$TERM_BIN" -e "$BIN" "$@"
            ;;
        *)
            exec "$TERM_BIN" -e "$BIN" "$@"
            ;;
    esac
done

# Fallback notification if no terminal exists
if command -v notify-send >/dev/null 2>&1; then
    notify-send -u critical "NothingTorrent" "Could not find a terminal emulator (kitty, alacritty, foot, ghostty, etc.)"
fi
exit 1
`

const DesktopFileContent = `[Desktop Entry]
Type=Application
Name=NothingTorrent
GenericName=BitTorrent Client
Comment=Minimalist Nothing OS styled BitTorrent TUI Client
Exec=%s %%U
Icon=nothing-torrent
Terminal=false
Categories=Network;FileTransfer;P2P;
MimeType=application/x-bittorrent;x-scheme-handler/magnet;
StartupNotify=false
Keywords=torrent;bittorrent;magnet;p2p;download;nothing;
`

const SvgIconContent = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128" width="128" height="128">
  <rect width="128" height="128" rx="28" fill="#0D0D0D"/>
  <circle cx="64" cy="64" r="50" fill="none" stroke="#262626" stroke-width="2" stroke-dasharray="4,4"/>
  
  <circle cx="44" cy="40" r="5" fill="#F0F0F0"/>
  <circle cx="44" cy="52" r="5" fill="#F0F0F0"/>
  <circle cx="44" cy="64" r="5" fill="#F0F0F0"/>
  <circle cx="44" cy="76" r="5" fill="#F0F0F0"/>
  <circle cx="44" cy="88" r="5" fill="#F0F0F0"/>
  
  <circle cx="56" cy="56" r="5" fill="#F0F0F0"/>
  <circle cx="68" cy="68" r="5" fill="#F0F0F0"/>
  <circle cx="76" cy="76" r="5" fill="#F0F0F0"/>
  
  <circle cx="84" cy="40" r="5" fill="#F0F0F0"/>
  <circle cx="84" cy="52" r="5" fill="#F0F0F0"/>
  <circle cx="84" cy="64" r="5" fill="#F0F0F0"/>
  <circle cx="84" cy="76" r="5" fill="#F0F0F0"/>
  <circle cx="84" cy="88" r="5" fill="#F0F0F0"/>

  <circle cx="98" cy="30" r="7" fill="#D71920"/>
</svg>
`

// Install registers nothing-torrent as default desktop app and torrent handler
func Install() error {
	if runtime.GOOS == "windows" {
		return installWindows()
	}
	return installLinux()
}

// Uninstall removes system registrations
func Uninstall() error {
	if runtime.GOOS == "windows" {
		return uninstallWindows()
	}
	return uninstallLinux()
}

// IsInstalled checks if desktop / protocol handler entry already exists
func IsInstalled() bool {
	if runtime.GOOS == "windows" {
		return isInstalledWindows()
	}
	return isInstalledLinux()
}

func installLinux() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home: %w", err)
	}

	targetBinDir := filepath.Join(home, ".local", "bin")
	_ = os.MkdirAll(targetBinDir, 0755)

	// 1. Locate current binary and copy to ~/.local/bin/nothing-torrent
	currentExec, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	currentExec, _ = filepath.EvalSymlinks(currentExec)

	targetBin := filepath.Join(targetBinDir, "nothing-torrent")
	if currentExec != targetBin {
		src, err := os.Open(currentExec)
		if err == nil {
			defer src.Close()
			dst, err := os.OpenFile(targetBin, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
			if err == nil {
				_, _ = io.Copy(dst, src)
				dst.Close()
			}
		}
	}

	// 2. Install universal launcher wrapper script
	launcherPath := filepath.Join(targetBinDir, "nothing-torrent-launcher")
	if err := os.WriteFile(launcherPath, []byte(LauncherScriptContent), 0755); err != nil {
		return fmt.Errorf("failed to write launcher script: %w", err)
	}
	_ = os.Chmod(launcherPath, 0755)

	// 3. Install Icon
	iconDir := filepath.Join(home, ".local", "share", "icons", "hicolor", "scalable", "apps")
	_ = os.MkdirAll(iconDir, 0755)
	iconPath := filepath.Join(iconDir, "nothing-torrent.svg")
	_ = os.WriteFile(iconPath, []byte(SvgIconContent), 0644)

	// 4. Install Desktop file with Terminal=false and launcher script
	appsDir := filepath.Join(home, ".local", "share", "applications")
	_ = os.MkdirAll(appsDir, 0755)
	desktopPath := filepath.Join(appsDir, "nothing-torrent.desktop")
	desktopData := fmt.Sprintf(DesktopFileContent, launcherPath)
	if err := os.WriteFile(desktopPath, []byte(desktopData), 0644); err != nil {
		return fmt.Errorf("failed to write desktop file: %w", err)
	}

	// 5. Update desktop database
	if updateDB, err := exec.LookPath("update-desktop-database"); err == nil {
		_ = exec.Command(updateDB, appsDir).Run()
	}

	// 6. Register MIME types
	if xdgMime, err := exec.LookPath("xdg-mime"); err == nil {
		_ = exec.Command(xdgMime, "default", "nothing-torrent.desktop", "application/x-bittorrent").Run()
		_ = exec.Command(xdgMime, "default", "nothing-torrent.desktop", "x-scheme-handler/magnet").Run()
	}

	return nil
}

func uninstallLinux() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	desktopPath := filepath.Join(home, ".local", "share", "applications", "nothing-torrent.desktop")
	_ = os.Remove(desktopPath)

	launcherPath := filepath.Join(home, ".local", "bin", "nothing-torrent-launcher")
	_ = os.Remove(launcherPath)

	iconPath := filepath.Join(home, ".local", "share", "icons", "hicolor", "scalable", "apps", "nothing-torrent.svg")
	_ = os.Remove(iconPath)

	appsDir := filepath.Join(home, ".local", "share", "applications")
	if updateDB, err := exec.LookPath("update-desktop-database"); err == nil {
		_ = exec.Command(updateDB, appsDir).Run()
	}
	return nil
}

func isInstalledLinux() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	desktopPath := filepath.Join(home, ".local", "share", "applications", "nothing-torrent.desktop")
	_, err = os.Stat(desktopPath)
	return err == nil
}

// Windows Registry Protocol and File Association Handlers
func installWindows() error {
	currentExec, err := os.Executable()
	if err != nil {
		return err
	}

	execCmd := fmt.Sprintf("\"%s\" \"%%1\"", currentExec)

	// 1. Register magnet protocol handler in HKCU (no admin rights needed)
	_ = exec.Command("reg", "add", `HKCU\Software\Classes\magnet`, "/ve", "/t", "REG_SZ", "/d", "URL:Magnet Protocol", "/f").Run()
	_ = exec.Command("reg", "add", `HKCU\Software\Classes\magnet`, "/v", "URL Protocol", "/t", "REG_SZ", "/d", "", "/f").Run()
	if err := exec.Command("reg", "add", `HKCU\Software\Classes\magnet\shell\open\command`, "/ve", "/t", "REG_SZ", "/d", execCmd, "/f").Run(); err != nil {
		return fmt.Errorf("failed to register magnet protocol: %w", err)
	}

	// 2. Register .torrent file association in HKCU
	_ = exec.Command("reg", "add", `HKCU\Software\Classes\.torrent`, "/ve", "/t", "REG_SZ", "/d", "NothingTorrent", "/f").Run()
	_ = exec.Command("reg", "add", `HKCU\Software\Classes\NothingTorrent`, "/ve", "/t", "REG_SZ", "/d", "NothingTorrent File", "/f").Run()
	_ = exec.Command("reg", "add", `HKCU\Software\Classes\NothingTorrent\shell\open\command`, "/ve", "/t", "REG_SZ", "/d", execCmd, "/f").Run()

	return nil
}

func uninstallWindows() error {
	_ = exec.Command("reg", "delete", `HKCU\Software\Classes\magnet\shell\open\command`, "/f").Run()
	_ = exec.Command("reg", "delete", `HKCU\Software\Classes\.torrent`, "/f").Run()
	_ = exec.Command("reg", "delete", `HKCU\Software\Classes\NothingTorrent`, "/f").Run()
	return nil
}

func isInstalledWindows() bool {
	cmd := exec.Command("reg", "query", `HKCU\Software\Classes\magnet\shell\open\command`)
	return cmd.Run() == nil
}
