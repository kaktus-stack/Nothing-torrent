package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// FormatBytes formats bytes into human-readable B, KB, MB, GB
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	units := []string{"KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.1f %s", float64(b)/float64(div), units[exp])
}

// FormatSpeed formats bytes/sec into speed string
func FormatSpeed(bytesPerSec int64) string {
	if bytesPerSec <= 0 {
		return "0.0 KB/s"
	}
	return fmt.Sprintf("%s/s", FormatBytes(bytesPerSec))
}

// FormatETA formats seconds into HH:MM:SS or MM:SS
func FormatETA(seconds int64) string {
	if seconds <= 0 || seconds > 86400*30 { // > 30 days or 0
		return "--:--:--"
	}
	d := time.Duration(seconds) * time.Second
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if h > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%02d:%02d", m, s)
}

// RenderProgressBar creates a Nothing-styled dot-matrix progress bar
// e.g. [■■■■■■■■■····················] 38.5%
func RenderProgressBar(width int, ratio float64) string {
	if width < 10 {
		width = 10
	}
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1.0 {
		ratio = 1.0
	}

	filled := int(float64(width) * ratio)
	empty := width - filled

	filledPart := lipgloss.NewStyle().Foreground(ColorRed).Render(strings.Repeat("■", filled))
	emptyPart := lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("·", empty))
	percentStr := fmt.Sprintf(" %5.1f%%", ratio*100)

	return "[" + filledPart + emptyPart + "]" + lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render(percentStr)
}

// RenderMiniBar creates a compact progress indicator for individual files
func RenderMiniBar(width int, ratio float64) string {
	if width < 5 {
		width = 5
	}
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1.0 {
		ratio = 1.0
	}
	filled := int(float64(width) * ratio)
	empty := width - filled
	f := lipgloss.NewStyle().Foreground(ColorWhite).Render(strings.Repeat("■", filled))
	e := lipgloss.NewStyle().Foreground(ColorDark).Render(strings.Repeat("·", empty))
	return f + e
}

// RenderSparkline renders a Unicode braille/block speed sparkline
func RenderSparkline(values []int64, maxLen int) string {
	if len(values) == 0 {
		return ""
	}
	if len(values) > maxLen {
		values = values[len(values)-maxLen:]
	}

	var maxVal int64 = 1
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}

	blocks := []rune{' ', '▂', '▃', '▄', '▅', '▆', '▇', '█'}
	var sb strings.Builder
	for _, v := range values {
		idx := int((float64(v) / float64(maxVal)) * float64(len(blocks)-1))
		if idx < 0 {
			idx = 0
		} else if idx >= len(blocks) {
			idx = len(blocks) - 1
		}
		sb.WriteRune(blocks[idx])
	}
	return lipgloss.NewStyle().Foreground(ColorWhite).Render(sb.String())
}

// Logo renders the Nothing (Torrent) header logo
func Logo() string {
	dot := RedDotStyle.Render("●")
	brand := lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render("NOTHING")
	subtitle := lipgloss.NewStyle().Foreground(ColorMuted).Render("(TORRENT)")
	return fmt.Sprintf("%s %s %s", dot, brand, subtitle)
}
