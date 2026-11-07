package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Nothing OS Palette
	ColorRed       = lipgloss.Color("#D71920") // Signature Nothing Red
	ColorWhite     = lipgloss.Color("#F2F2F2") // High-contrast Stark White
	ColorMuted     = lipgloss.Color("#737373") // Muted Subtitle Gray
	ColorDark      = lipgloss.Color("#242424") // Subtle Border Dark Gray
	ColorDarker    = lipgloss.Color("#141414") // Card Background
	ColorBlack     = lipgloss.Color("#0A0A0A") // Deep Background
	ColorHighlight = lipgloss.Color("#FFFFFF") // Pure White

	// Typography & Container Styles
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorWhite).
			Background(ColorDarker).
			Padding(0, 1)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorWhite)

	RedDotStyle = lipgloss.NewStyle().
			Foreground(ColorRed).
			Bold(true)

	BadgeStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Background(ColorDark).
			Padding(0, 1)

	BadgeRedStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Background(ColorRed).
			Bold(true).
			Padding(0, 1)

	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorDark).
			Padding(1, 2)

	StatValueStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorWhite)

	StatLabelStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	SelectedFileStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorWhite).
				Background(ColorDarker)

	UnselectedFileStyle = lipgloss.NewStyle().
				Foreground(ColorMuted)

	HelpKeyStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Bold(true)

	HelpDescStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)
)
