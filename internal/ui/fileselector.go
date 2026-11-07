package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nothing-torrent/nothing-torrent/internal/torrent"
)

// RenderFileSelector renders the file selection screen
func RenderFileSelector(
	files []*torrent.FileEntry,
	cursor int,
	scrollOffset int,
	viewportHeight int,
	torrentName string,
	width int,
) string {
	var b strings.Builder

	// 1. Header
	header := fmt.Sprintf("%s   %s",
		Logo(),
		BadgeStyle.Render("FILE SELECTOR"),
	)
	b.WriteString(header + "\n\n")

	// 2. Torrent Title & info
	title := TitleStyle.Render(torrentName)
	b.WriteString(title + "\n")

	// Calculate counts and sizes
	var totalSize int64
	var selectedSize int64
	var selectedCount int
	for _, f := range files {
		totalSize += f.Length
		if f.Selected {
			selectedCount++
			selectedSize += f.Length
		}
	}

	statsLine := fmt.Sprintf(
		"Total Files: %s   Total Size: %s   Selected: %s (%s)",
		StatValueStyle.Render(fmt.Sprintf("%d", len(files))),
		StatValueStyle.Render(FormatBytes(totalSize)),
		RedDotStyle.Render(fmt.Sprintf("%d/%d", selectedCount, len(files))),
		StatValueStyle.Render(FormatBytes(selectedSize)),
	)
	b.WriteString(statsLine + "\n")
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n\n")

	// 3. File list viewport
	maxVisible := viewportHeight - 10
	if maxVisible < 5 {
		maxVisible = 5
	}

	endIdx := scrollOffset + maxVisible
	if endIdx > len(files) {
		endIdx = len(files)
	}

	for i := scrollOffset; i < endIdx; i++ {
		f := files[i]
		isCursor := (i == cursor)

		var cursorMarker string
		if isCursor {
			cursorMarker = RedDotStyle.Render("› ")
		} else {
			cursorMarker = "  "
		}

		var checkMarker string
		if f.Selected {
			checkMarker = RedDotStyle.Render("[●] ")
		} else {
			checkMarker = lipgloss.NewStyle().Foreground(ColorDark).Render("[○] ")
		}

		sizeStr := StatLabelStyle.Render(fmt.Sprintf("(%s)", FormatBytes(f.Length)))
		filename := f.DisplayPath

		// Truncate path if too wide
		availWidth := width - 30
		if availWidth > 20 && len(filename) > availWidth {
			filename = "..." + filename[len(filename)-availWidth:]
		}

		var line string
		if isCursor {
			nameStyled := lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render(filename)
			line = fmt.Sprintf("%s%s%s %s", cursorMarker, checkMarker, nameStyled, sizeStr)
			line = lipgloss.NewStyle().Background(ColorDarker).Width(width).Render(line)
		} else {
			var nameStyled string
			if f.Selected {
				nameStyled = lipgloss.NewStyle().Foreground(ColorWhite).Render(filename)
			} else {
				nameStyled = lipgloss.NewStyle().Foreground(ColorMuted).Render(filename)
			}
			line = fmt.Sprintf("%s%s%s %s", cursorMarker, checkMarker, nameStyled, sizeStr)
		}

		b.WriteString(line + "\n")
	}

	// 4. Scroll indicator
	if len(files) > maxVisible {
		scrollInfo := fmt.Sprintf("▲ ▼ [%d-%d of %d]", scrollOffset+1, endIdx, len(files))
		b.WriteString("\n" + lipgloss.NewStyle().Foreground(ColorMuted).Render(scrollInfo) + "\n")
	} else {
		b.WriteString("\n")
	}

	// 5. Footer hotkeys
	b.WriteString(lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("─", width)) + "\n")
	footer := fmt.Sprintf(
		"%s %s   %s %s   %s %s   %s %s   %s %s",
		HelpKeyStyle.Render("[↑/↓]"), HelpDescStyle.Render("Navigate"),
		HelpKeyStyle.Render("[Space]"), HelpDescStyle.Render("Toggle"),
		HelpKeyStyle.Render("[a]"), HelpDescStyle.Render("All"),
		HelpKeyStyle.Render("[n]"), HelpDescStyle.Render("None"),
		BadgeRedStyle.Render("[Enter] START DOWNLOAD"),
	)
	b.WriteString(footer + "\n")

	return b.String()
}
