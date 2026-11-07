package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
)

type SettingsField int

const (
	FieldStealth SettingsField = iota
	FieldTargetRatio
	FieldClientProfile
	FieldDisableSeed
	FieldUploadSpeed
	FieldDownloadDir
	FieldCount
)

// RenderSettings renders the interactive Nothing-styled Settings menu
func RenderSettings(cfg *config.Config, selectedField SettingsField, width int) string {
	var b strings.Builder

	// Header
	header := fmt.Sprintf("%s   %s", Logo(), BadgeRedStyle.Render("SETTINGS"))
	b.WriteString(header + "\n\n")

	b.WriteString(TitleStyle.Render("SYSTEM & STEALTH CONFIGURATION") + "\n")
	b.WriteString(lipgloss.NewStyle().Foreground(ColorMuted).Render("All changes are automatically saved to disk") + "\n")
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n\n")

	// Render items
	fields := []struct {
		id    SettingsField
		label string
		val   string
		desc  string
	}{
		{
			id:    FieldStealth,
			label: "Stealth Ratio Guard",
			val: func() string {
				if cfg.StealthMode {
					return RedDotStyle.Render("[●] ENABLED (Stealth Active)")
				}
				return lipgloss.NewStyle().Foreground(ColorDark).Render("[○] DISABLED")
			}(),
			desc: "Fakes upload stats on tracker announce requests to maintain target ratio",
		},
		{
			id:    FieldTargetRatio,
			label: "Target Ratio",
			val:   fmt.Sprintf("◀  %s  ▶", RedDotStyle.Render(fmt.Sprintf("%.2fx", cfg.TargetRatio))),
			desc:  "Safe range: 0.80x – 0.90x avoids download speed penalties & tracker bans",
		},
		{
			id:    FieldClientProfile,
			label: "Client Emulation",
			val:   fmt.Sprintf("◀  %s  ▶", StatValueStyle.Render(config.AvailableProfiles[cfg.DefaultClient].Name)),
			desc:  "Spoofs Peer-ID and HTTP User-Agent to pass tracker whitelists",
		},
		{
			id:    FieldDisableSeed,
			label: "Cold-Router Mode",
			val: func() string {
				if cfg.DisableSeed {
					return RedDotStyle.Render("[●] ACTIVE (Zero Upload / Cold Hardware)")
				}
				return StatValueStyle.Render("[○] OFF (Uploads actual chunks)")
			}(),
			desc: "Does not seed chunks to peers; keeps router cold and saves bandwidth",
		},
		{
			id:    FieldUploadSpeed,
			label: "Simulated Upload Speed",
			val:   fmt.Sprintf("◀  %d – %d KB/s  ▶", cfg.MinUploadKBps, cfg.MaxUploadKBps),
			desc:  "Reported speed variation jitter so upload graph looks realistic to trackers",
		},
		{
			id:    FieldDownloadDir,
			label: "Download Directory",
			val:   StatValueStyle.Render(cfg.DownloadDir),
			desc:  "Files will be saved into this directory",
		},
	}

	for _, f := range fields {
		isSelected := (f.id == selectedField)

		var cursor string
		if isSelected {
			cursor = RedDotStyle.Render("› ")
		} else {
			cursor = "  "
		}

		labelStyled := lipgloss.NewStyle().Bold(isSelected).Foreground(ColorWhite).Render(fmt.Sprintf("%-25s", f.label))
		line := fmt.Sprintf("%s%s %s", cursor, labelStyled, f.val)

		if isSelected {
			line = lipgloss.NewStyle().Background(ColorDarker).Width(width).Render(line)
			b.WriteString(line + "\n")
			descStyled := lipgloss.NewStyle().Foreground(ColorMuted).Render(fmt.Sprintf("     %s", f.desc))
			b.WriteString(descStyled + "\n\n")
		} else {
			b.WriteString(line + "\n\n")
		}
	}

	// Footer hotkeys
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	footer := fmt.Sprintf(
		"%s %s   %s %s   %s   %s %s",
		HelpKeyStyle.Render("[↑/↓]"), HelpDescStyle.Render("Navigate"),
		HelpKeyStyle.Render("[←/→ / Space]"), HelpDescStyle.Render("Change"),
		BadgeStyle.Render("[Esc / s] Save & Back"),
		HelpKeyStyle.Render("[q]"), HelpDescStyle.Render("Quit"),
	)
	b.WriteString(footer + "\n")

	return b.String()
}
