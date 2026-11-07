package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
	"github.com/nothing-torrent/nothing-torrent/internal/torrent"
)

// RenderDashboard renders the live download dashboard with responsive layout
func RenderDashboard(
	torrentName string,
	files []*torrent.FileEntry,
	completedBytes int64,
	totalSelectedBytes int64,
	speed int64,
	peers int,
	speedHistory []int64,
	cfg *config.Config,
	stealth *torrent.StealthManager,
	isPaused bool,
	width int,
	viewportHeight int,
) string {
	var b strings.Builder

	if width < 40 {
		width = 40
	}

	// Calculate overall ratio and progress
	var progressRatio float64
	if totalSelectedBytes > 0 {
		progressRatio = float64(completedBytes) / float64(totalSelectedBytes)
	}

	fakeUploaded, currentRatio := stealth.GetCurrentStats(completedBytes)
	simUploadSpeed := stealth.GetSimulatedUploadSpeed(speed > 0 && !isPaused)

	// Calculate ETA
	var etaSeconds int64
	if speed > 0 && totalSelectedBytes > completedBytes {
		etaSeconds = (totalSelectedBytes - completedBytes) / speed
	}

	// 1. Header with Nothing logo and stealth status
	statusBadge := BadgeStyle.Render("DOWNLOADING")
	if isPaused {
		statusBadge = BadgeRedStyle.Render("PAUSED")
	} else if progressRatio >= 1.0 {
		statusBadge = BadgeStyle.Render("COMPLETED")
	}

	var stealthBadge string
	if cfg.StealthMode {
		clientProfile := cfg.GetUserAgent()
		if width < 70 {
			// Compact badge for narrow terminals
			stealthBadge = fmt.Sprintf("[%s %s]", RedDotStyle.Render("●"), torrent.FormatRatio(currentRatio))
		} else {
			stealthBadge = fmt.Sprintf(
				"[%s %s · %s]",
				RedDotStyle.Render("●"),
				lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render(fmt.Sprintf("RATIO %s", torrent.FormatRatio(currentRatio))),
				lipgloss.NewStyle().Foreground(ColorMuted).Render(clientProfile),
			)
		}
	} else {
		stealthBadge = lipgloss.NewStyle().Foreground(ColorMuted).Render("[STEALTH: OFF]")
	}

	header := fmt.Sprintf("%s   %s   %s", Logo(), statusBadge, stealthBadge)
	b.WriteString(header + "\n\n")

	// 2. Torrent Title (truncated if wider than terminal)
	displayTitle := torrentName
	if len(displayTitle) > width-4 {
		displayTitle = displayTitle[:width-7] + "..."
	}
	b.WriteString(TitleStyle.Render(displayTitle) + "\n\n")

	// 3. Main Progress Bar (Adaptive width)
	barWidth := width - 16
	if barWidth > 65 {
		barWidth = 65
	} else if barWidth < 12 {
		barWidth = 12
	}
	progressBar := RenderProgressBar(barWidth, progressRatio)
	b.WriteString(progressBar + "\n\n")

	// 4. Telemetry Grid (Responsive: 3 cards if wide, 2 or stacked if narrow)
	downSpeedStr := FormatSpeed(speed)
	upSpeedStr := fmt.Sprintf("%d KB/s (Stealth)", simUploadSpeed)
	downDataStr := fmt.Sprintf("%s / %s", FormatBytes(completedBytes), FormatBytes(totalSelectedBytes))
	upDataStr := fmt.Sprintf("%s (Sim)", FormatBytes(fakeUploaded))

	if width >= 78 {
		// Wide screen: 3 horizontal cards
		col1 := fmt.Sprintf("DOWN SPEED: %s\nDOWNLOADED: %s",
			StatValueStyle.Render(downSpeedStr),
			StatValueStyle.Render(downDataStr),
		)
		col2 := fmt.Sprintf("STEALTH UP: %s\nREPORTED:   %s",
			RedDotStyle.Render(upSpeedStr),
			StatValueStyle.Render(upDataStr),
		)
		col3 := fmt.Sprintf("ETA:        %s\nPEERS:      %s",
			StatValueStyle.Render(FormatETA(etaSeconds)),
			StatValueStyle.Render(fmt.Sprintf("%d", peers)),
		)
		grid := lipgloss.JoinHorizontal(lipgloss.Top,
			CardStyle.Render(col1),
			CardStyle.Render(col2),
			CardStyle.Render(col3),
		)
		b.WriteString(grid + "\n\n")
	} else {
		// Compact / Narrow screen: clean tabular telemetry
		line1 := fmt.Sprintf(
			"DOWN: %s (%s)   PEERS: %s",
			StatValueStyle.Render(downSpeedStr),
			StatLabelStyle.Render(downDataStr),
			StatValueStyle.Render(fmt.Sprintf("%d", peers)),
		)
		line2 := fmt.Sprintf(
			"STEALTH UP: %s (%s)   ETA: %s",
			RedDotStyle.Render(upSpeedStr),
			StatLabelStyle.Render(upDataStr),
			StatValueStyle.Render(FormatETA(etaSeconds)),
		)
		b.WriteString(line1 + "\n" + line2 + "\n\n")
	}

	// 5. Sparkline Speed History
	if len(speedHistory) > 0 && width >= 45 {
		sparkLen := width - 18
		if sparkLen > 40 {
			sparkLen = 40
		}
		spark := RenderSparkline(speedHistory, sparkLen)
		b.WriteString(fmt.Sprintf("SPEED GRAPH  %s\n\n", spark))
	}

	// 6. Selected Files Progress
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render("FILES PROGRESS") + "\n\n")

	maxVisibleFiles := viewportHeight - 22
	if maxVisibleFiles < 3 {
		maxVisibleFiles = 3
	}

	shown := 0
	for _, f := range files {
		if !f.Selected {
			continue
		}
		if shown >= maxVisibleFiles {
			remaining := 0
			for _, rest := range files[shown:] {
				if rest.Selected {
					remaining++
				}
			}
			if remaining > 0 {
				b.WriteString(lipgloss.NewStyle().Foreground(ColorMuted).Render(fmt.Sprintf("... and %d more files", remaining)) + "\n")
			}
			break
		}

		var fileRatio float64
		if f.Length > 0 {
			fileRatio = float64(f.BytesCompleted) / float64(f.Length)
		}

		var fileIcon string
		if fileRatio >= 1.0 {
			fileIcon = RedDotStyle.Render("●")
		} else if fileRatio > 0 {
			fileIcon = lipgloss.NewStyle().Foreground(ColorWhite).Render("›")
		} else {
			fileIcon = lipgloss.NewStyle().Foreground(ColorDark).Render("○")
		}

		filename := f.DisplayPath

		if width >= 70 {
			// Wide layout: includes mini progress bar
			maxNameLen := width - 42
			if maxNameLen > 15 && len(filename) > maxNameLen {
				filename = "..." + filename[len(filename)-maxNameLen:]
			}
			miniBar := RenderMiniBar(10, fileRatio)
			line := fmt.Sprintf(
				"%s %-32s %s %5.1f%% (%s)",
				fileIcon,
				filename,
				miniBar,
				fileRatio*100,
				FormatBytes(f.Length),
			)
			b.WriteString(line + "\n")
		} else {
			// Compact layout: no mini bar, clean percent & size
			maxNameLen := width - 25
			if maxNameLen > 10 && len(filename) > maxNameLen {
				filename = "..." + filename[len(filename)-maxNameLen:]
			}
			line := fmt.Sprintf(
				"%s %-20s %5.1f%% (%s)",
				fileIcon,
				filename,
				fileRatio*100,
				FormatBytes(f.Length),
			)
			b.WriteString(line + "\n")
		}
		shown++
	}

	// 7. Footer hotkeys (Responsive wrapping)
	b.WriteString("\n" + lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	if width >= 65 {
		footer := fmt.Sprintf(
			"%s %s   %s %s   %s %s   %s %s",
			HelpKeyStyle.Render("[p]"), HelpDescStyle.Render("Pause/Resume"),
			HelpKeyStyle.Render("[s]"), HelpDescStyle.Render("Settings"),
			HelpKeyStyle.Render("[o]"), HelpDescStyle.Render("Open Folder"),
			HelpKeyStyle.Render("[q]"), HelpDescStyle.Render("Quit"),
		)
		b.WriteString(footer + "\n")
	} else {
		footer := fmt.Sprintf(
			"%s %s  %s %s  %s %s",
			HelpKeyStyle.Render("[p]"), HelpDescStyle.Render("Pause"),
			HelpKeyStyle.Render("[s]"), HelpDescStyle.Render("Settings"),
			HelpKeyStyle.Render("[q]"), HelpDescStyle.Render("Quit"),
		)
		b.WriteString(footer + "\n")
	}

	return b.String()
}
