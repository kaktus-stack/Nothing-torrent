package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
	"github.com/nothing-torrent/nothing-torrent/internal/desktop"
	"github.com/nothing-torrent/nothing-torrent/internal/torrent"
	"github.com/nothing-torrent/nothing-torrent/internal/ui"
)

func main() {
	// If launched from GUI/desktop without a terminal attached (Linux only),
	// spawn ourselves inside the user's terminal emulator
	if runtime.GOOS != "windows" && !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsTerminal(os.Stdout.Fd()) {
		home, _ := os.UserHomeDir()
		launcher := filepath.Join(home, ".local", "bin", "nothing-torrent-launcher")
		if _, err := os.Stat(launcher); err == nil {
			cmd := exec.Command(launcher, os.Args[1:]...)
			_ = cmd.Start()
			return
		}
	}

	installFlag := flag.Bool("install", false, "Register nothing-torrent as default desktop app and torrent handler")
	uninstallFlag := flag.Bool("uninstall", false, "Remove desktop app registration")
	flag.Parse()

	args := flag.Args()

	// Handle install subcommand or flag
	if *installFlag || (len(args) > 0 && args[0] == "install") {
		handleInstall()
		return
	}

	// Handle uninstall subcommand or flag
	if *uninstallFlag || (len(args) > 0 && args[0] == "uninstall") {
		handleUninstall()
		return
	}

	// Auto-install desktop entry if running for the first time
	if !desktop.IsInstalled() {
		_ = desktop.Install()
	}

	// Determine torrent source (magnet link or file)
	var source string
	if len(args) > 0 {
		source = args[0]
	} else {
		source = promptForTorrent()
		if source == "" {
			fmt.Println("No torrent provided. Exiting.")
			return
		}
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		return
	}

	// Initialize Torrent Engine
	engine, err := torrent.NewEngine(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize engine: %v\n", err)
		return
	}
	defer engine.Close()

	// Load torrent
	if err := engine.LoadTorrent(source); err != nil {
		fmt.Printf("Error loading torrent: %v\n", err)
		return
	}

	// Launch Bubble Tea TUI
	appModel := ui.NewModel(cfg, engine, source)
	p := tea.NewProgram(appModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}
}

func handleInstall() {
	fmt.Println()
	fmt.Println(ui.Logo())
	fmt.Println(lipgloss.NewStyle().Foreground(ui.ColorMuted).Render("Installing Linux desktop integration..."))

	if err := desktop.Install(); err != nil {
		fmt.Println(ui.BadgeRedStyle.Render("FAILED:"), err)
		return
	}

	fmt.Println()
	fmt.Println(ui.RedDotStyle.Render("●") + " Desktop entry installed: " + lipgloss.NewStyle().Foreground(ui.ColorWhite).Render("~/.local/share/applications/nothing-torrent.desktop"))
	fmt.Println(ui.RedDotStyle.Render("●") + " Universal launcher:      " + lipgloss.NewStyle().Foreground(ui.ColorWhite).Render("~/.local/bin/nothing-torrent-launcher"))
	fmt.Println(ui.RedDotStyle.Render("●") + " High-res icon installed: " + lipgloss.NewStyle().Foreground(ui.ColorWhite).Render("~/.local/share/icons/hicolor/scalable/apps/nothing-torrent.svg"))
	fmt.Println(ui.RedDotStyle.Render("●") + " Binary registered at:    " + lipgloss.NewStyle().Foreground(ui.ColorWhite).Render("~/.local/bin/nothing-torrent"))
	fmt.Println(ui.RedDotStyle.Render("●") + " Default handler set for: " + lipgloss.NewStyle().Foreground(ui.ColorWhite).Render("application/x-bittorrent & magnet: links"))
	fmt.Println()
	fmt.Println(ui.BadgeStyle.Render("READY"), "Compatible with Wofi, Rofi, Dolphin, Nautilus, Thunar, Chrome, Firefox, etc.!")
	fmt.Println()
}

func handleUninstall() {
	if err := desktop.Uninstall(); err != nil {
		fmt.Println(ui.BadgeRedStyle.Render("FAILED:"), err)
		return
	}
	fmt.Println(ui.Logo())
	fmt.Println("Desktop integration removed.")
}

func promptForTorrent() string {
	fmt.Println()
	fmt.Println(ui.Logo() + "   " + ui.BadgeStyle.Render("READY"))
	fmt.Println(lipgloss.NewStyle().Foreground(ui.ColorMuted).Render("No torrent specified. Paste magnet link or path to .torrent file:"))
	fmt.Print(ui.RedDotStyle.Render("› "))

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
