package desktop

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const DesktopFileContent = `[Desktop Entry]
Type=Application
Name=NothingTorrent
GenericName=BitTorrent Client
Comment=Minimalist Nothing OS styled BitTorrent TUI Client
Exec=%s %%U
Icon=nothing-torrent
Terminal=true
Categories=Network;FileTransfer;P2P;
MimeType=application/x-bittorrent;x-scheme-handler/magnet;
StartupNotify=false
`

const SvgIconContent = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128" width="128" height="128">
  <rect width="128" height="128" rx="28" fill="#0D0D0D"/>
  <circle cx="64" cy="64" r="50" fill="none" stroke="#262626" stroke-width="2" stroke-dasharray="4,4"/>
  
  <!-- Dot matrix letter N -->
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

  <!-- Iconic Nothing Red Dot -->
  <circle cx="98" cy="30" r="7" fill="#D71920"/>
</svg>
`

// DetectTerminal returns available terminal emulator
func DetectTerminal() string {
	if term := os.Getenv("TERMINAL"); term != "" {
		if path, err := exec.LookPath(term); err == nil {
			return path
		}
	}
	candidates := []string{
		"ghostty",
		"kitty",
		"alacritty",
		"foot",
		"wezterm",
		"gnome-terminal",
		"xfce4-terminal",
		"konsole",
		"x-terminal-emulator",
		"xterm",
	}
	for _, c := range candidates {
		if path, err := exec.LookPath(c); err == nil {
			return path
		}
	}
	return ""
}

// Install registers nothing-torrent as desktop app and default torrent handler
func Install() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home: %w", err)
	}

	// 1. Locate current binary or target location
	currentExec, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	currentExec, _ = filepath.EvalSymlinks(currentExec)

	targetBinDir := filepath.Join(home, ".local", "bin")
	_ = os.MkdirAll(targetBinDir, 0755)
	targetBin := filepath.Join(targetBinDir, "nothing-torrent")

	// Copy binary to ~/.local/bin/nothing-torrent if not running from there
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

	execCommand := targetBin
	if _, err := os.Stat(targetBin); err != nil {
		execCommand = currentExec
	}

	// If the user's desktop environment doesn't handle Terminal=true well,
	// we check if a known terminal exists to wrap it
	term := DetectTerminal()
	var finalExec string
	if term != "" && (filepath.Base(term) == "kitty" || filepath.Base(term) == "alacritty" || filepath.Base(term) == "foot" || filepath.Base(term) == "ghostty") {
		finalExec = fmt.Sprintf("%s -e %s", term, execCommand)
	} else {
		finalExec = execCommand
	}

	// 2. Install Icon
	iconDir := filepath.Join(home, ".local", "share", "icons", "hicolor", "scalable", "apps")
	_ = os.MkdirAll(iconDir, 0755)
	iconPath := filepath.Join(iconDir, "nothing-torrent.svg")
	_ = os.WriteFile(iconPath, []byte(SvgIconContent), 0644)

	// 3. Install Desktop file
	appsDir := filepath.Join(home, ".local", "share", "applications")
	_ = os.MkdirAll(appsDir, 0755)
	desktopPath := filepath.Join(appsDir, "nothing-torrent.desktop")
	desktopData := fmt.Sprintf(DesktopFileContent, finalExec)
	if err := os.WriteFile(desktopPath, []byte(desktopData), 0644); err != nil {
		return fmt.Errorf("failed to write desktop file: %w", err)
	}

	// 4. Update desktop database
	if updateDB, err := exec.LookPath("update-desktop-database"); err == nil {
		_ = exec.Command(updateDB, appsDir).Run()
	}

	// 5. Register MIME types
	if xdgMime, err := exec.LookPath("xdg-mime"); err == nil {
		_ = exec.Command(xdgMime, "default", "nothing-torrent.desktop", "application/x-bittorrent").Run()
		_ = exec.Command(xdgMime, "default", "nothing-torrent.desktop", "x-scheme-handler/magnet").Run()
	}

	return nil
}

// Uninstall removes desktop file and icon
func Uninstall() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	desktopPath := filepath.Join(home, ".local", "share", "applications", "nothing-torrent.desktop")
	_ = os.Remove(desktopPath)

	iconPath := filepath.Join(home, ".local", "share", "icons", "hicolor", "scalable", "apps", "nothing-torrent.svg")
	_ = os.Remove(iconPath)

	appsDir := filepath.Join(home, ".local", "share", "applications")
	if updateDB, err := exec.LookPath("update-desktop-database"); err == nil {
		_ = exec.Command(updateDB, appsDir).Run()
	}
	return nil
}

// IsInstalled checks if desktop entry already exists
func IsInstalled() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	desktopPath := filepath.Join(home, ".local", "share", "applications", "nothing-torrent.desktop")
	_, err = os.Stat(desktopPath)
	return err == nil
}
