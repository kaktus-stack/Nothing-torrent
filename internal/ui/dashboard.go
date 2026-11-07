package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
	"github.com/nothing-torrent/nothing-torrent/internal/torrent"
)

// RenderDashboard renders the live download dashboard
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
		stealthBadge = fmt.Sprintf(
			"[%s %s · %s]",
			RedDotStyle.Render("●"),
			lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render(fmt.Sprintf("RATIO %s", torrent.FormatRatio(currentRatio))),
			lipgloss.NewStyle().Foreground(ColorMuted).Render(clientProfile),
		)
	} else {
		stealthBadge = lipgloss.NewStyle().Foreground(ColorMuted).Render("[STEALTH: OFF]")
	}

	header := fmt.Sprintf("%s   %s   %s", Logo(), statusBadge, stealthBadge)
	b.WriteString(header + "\n\n")

	// 2. Torrent Title
	b.WriteString(TitleStyle.Render(torrentName) + "\n\n")

	// 3. Main Progress Bar
	barWidth := width - 18
	if barWidth > 60 {
		barWidth = 60
	} else if barWidth < 20 {
		barWidth = 20
	}
	progressBar := RenderProgressBar(barWidth, progressRatio)
	b.WriteString(progressBar + "\n\n")

	// 4. Telemetry Grid (2 columns)
	downSpeedStr := FormatSpeed(speed)
	upSpeedStr := fmt.Sprintf("%d KB/s (Stealth)", simUploadSpeed)
	downDataStr := fmt.Sprintf("%s / %s", FormatBytes(completedBytes), FormatBytes(totalSelectedBytes))
	upDataStr := fmt.Sprintf("%s (Simulated)", FormatBytes(fakeUploaded))

	col1 := fmt.Sprintf(
		"DOWN SPEED: %s\nDOWNLOADED: %s",
		StatValueStyle.Render(downSpeedStr),
		StatValueStyle.Render(downDataStr),
	)

	col2 := fmt.Sprintf(
		"STEALTH UP: %s\nREPORTED:   %s",
		RedDotStyle.Render(upSpeedStr),
		StatValueStyle.Render(upDataStr),
	)

	col3 := fmt.Sprintf(
		"ETA:        %s\nPEERS:      %s",
		StatValueStyle.Render(FormatETA(etaSeconds)),
		StatValueStyle.Render(fmt.Sprintf("%d", peers)),
	)

	grid := lipgloss.JoinHorizontal(lipgloss.Top,
		CardStyle.Render(col1),
		CardStyle.Render(col2),
		CardStyle.Render(col3),
	)
	b.WriteString(grid + "\n\n")

	// 5. Sparkline Speed History
	if len(speedHistory) > 0 {
		spark := RenderSparkline(speedHistory, width-20)
		b.WriteString(fmt.Sprintf("SPEED GRAPH  %s\n\n", spark))
	}

	// 6. Selected Files Progress
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render("FILES PROGRESS") + "\n\n")

	maxVisibleFiles := viewportHeight - 22
	if maxVisibleFiles < 4 {
		maxVisibleFiles = 4
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

		miniBar := RenderMiniBar(12, fileRatio)
		filename := f.DisplayPath
		maxNameLen := width - 42
		if maxNameLen > 15 && len(filename) > maxNameLen {
			filename = "..." + filename[len(filename)-maxNameLen:]
		}

		line := fmt.Sprintf(
			"%s %-35s %s %5.1f%% (%s)",
			fileIcon,
			filename,
			miniBar,
			fileRatio*100,
			FormatBytes(f.Length),
		)
		b.WriteString(line + "\n")
		shown++
	}

	// 7. Footer hotkeys
	b.WriteString("\n" + lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	footer := fmt.Sprintf(
		"%s %s   %s %s   %s %s   %s %s",
		HelpKeyStyle.Render("[p]"), HelpDescStyle.Render("Pause/Resume"),
		HelpKeyStyle.Render("[s]"), HelpDescStyle.Render("Toggle Stealth"),
		HelpKeyStyle.Render("[o]"), HelpDescStyle.Render("Open Folder"),
		HelpKeyStyle.Render("[q]"), HelpDescStyle.Render("Quit"),
	)
	b.WriteString(footer + "\n")

	return b.String()
}
