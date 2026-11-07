package ui

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
	"github.com/nothing-torrent/nothing-torrent/internal/torrent"
)

type AppState int

const (
	StateResolving AppState = iota
	StateFileSelect
	StateDownloading
	StateError
)

type tickMsg time.Time
type metadataReadyMsg struct{}
type metadataErrMsg struct{ err error }

// Model is the main Bubble Tea application state
type Model struct {
	state          AppState
	cfg            *config.Config
	engine         *torrent.TorrentEngine
	source         string
	err            error

	// Window dimensions
	width          int
	height         int

	// Resolving animation
	dotCount       int

	// File selection state
	cursor         int
	scrollOffset   int

	// Download state
	completedBytes int64
	totalBytes     int64
	speed          int64
	peers          int
	speedHistory   []int64
	isPaused       bool
}

// NewModel creates the Bubble Tea model
func NewModel(cfg *config.Config, engine *torrent.TorrentEngine, source string) Model {
	return Model{
		state:        StateResolving,
		cfg:          cfg,
		engine:       engine,
		source:       source,
		width:        80,
		height:       24,
		speedHistory: make([]int64, 0, 40),
	}
}

// Init starts resolving metadata
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForMetadataCmd(),
		m.tickCmd(),
	)
}

func (m Model) tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) waitForMetadataCmd() tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()

		if err := m.engine.WaitForMetadata(ctx); err != nil {
			return metadataErrMsg{err: err}
		}
		return metadataReadyMsg{}
	}
}

// Update handles incoming events and keyboard input
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case metadataReadyMsg:
		m.state = StateFileSelect
		m.cursor = 0
		m.scrollOffset = 0
		return m, nil

	case metadataErrMsg:
		m.state = StateError
		m.err = msg.err
		return m, nil

	case tickMsg:
		m.dotCount = (m.dotCount + 1) % 4
		if m.state == StateDownloading && !m.isPaused {
			comp, total, spd, prs := m.engine.UpdateStats()
			m.completedBytes = comp
			m.totalBytes = total
			m.speed = spd
			m.peers = prs

			m.speedHistory = append(m.speedHistory, spd)
			if len(m.speedHistory) > 50 {
				m.speedHistory = m.speedHistory[1:]
			}
		}
		return m, m.tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.engine.Close()
			return m, tea.Quit

		case "up", "k":
			if m.state == StateFileSelect {
				if m.cursor > 0 {
					m.cursor--
					if m.cursor < m.scrollOffset {
						m.scrollOffset = m.cursor
					}
				}
			}

		case "down", "j":
			if m.state == StateFileSelect {
				if m.cursor < len(m.engine.Files)-1 {
					m.cursor++
					maxVisible := m.height - 10
					if maxVisible < 5 {
						maxVisible = 5
					}
					if m.cursor >= m.scrollOffset+maxVisible {
						m.scrollOffset = m.cursor - maxVisible + 1
					}
				}
			}

		case " ":
			if m.state == StateFileSelect && len(m.engine.Files) > 0 {
				m.engine.Files[m.cursor].Selected = !m.engine.Files[m.cursor].Selected
			}

		case "a":
			if m.state == StateFileSelect {
				for _, f := range m.engine.Files {
					f.Selected = true
				}
			}

		case "n":
			if m.state == StateFileSelect {
				for _, f := range m.engine.Files {
					f.Selected = false
				}
			}

		case "i":
			if m.state == StateFileSelect {
				for _, f := range m.engine.Files {
					f.Selected = !f.Selected
				}
			}

		case "enter":
			if m.state == StateFileSelect {
				m.engine.ApplySelection()
				m.state = StateDownloading
				return m, nil
			}

		case "p":
			if m.state == StateDownloading {
				m.isPaused = !m.isPaused
			}

		case "s":
			if m.state == StateDownloading {
				m.cfg.StealthMode = !m.cfg.StealthMode
				_ = m.cfg.Save()
			}

		case "o":
			_ = exec.Command("xdg-open", m.cfg.DownloadDir).Start()
		}
	}

	return m, nil
}

// View renders the screen based on state
func (m Model) View() string {
	switch m.state {
	case StateResolving:
		dots := strings.Repeat("● ", m.dotCount) + strings.Repeat("○ ", 3-m.dotCount)
		header := fmt.Sprintf("%s   %s", Logo(), BadgeStyle.Render("RESOLVING"))
		body := fmt.Sprintf(
			"\n\nConnecting to DHT and Trackers...\n\n%s\n\n%s\n\n%s",
			lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render("Fetching torrent metadata:"),
			RedDotStyle.Render("[ "+dots+"]"),
			lipgloss.NewStyle().Foreground(ColorMuted).Render("Press [q] to cancel"),
		)
		return header + body

	case StateFileSelect:
		return RenderFileSelector(
			m.engine.Files,
			m.cursor,
			m.scrollOffset,
			m.height,
			m.engine.Name,
			m.width,
		)

	case StateDownloading:
		return RenderDashboard(
			m.engine.Name,
			m.engine.Files,
			m.completedBytes,
			m.totalBytes,
			m.speed,
			m.peers,
			m.speedHistory,
			m.cfg,
			m.engine.Stealth,
			m.isPaused,
			m.width,
			m.height,
		)

	case StateError:
		header := fmt.Sprintf("%s   %s", Logo(), BadgeRedStyle.Render("ERROR"))
		body := fmt.Sprintf(
			"\n\n%s\n\n%s",
			lipgloss.NewStyle().Bold(true).Foreground(ColorWhite).Render(fmt.Sprintf("Failed: %v", m.err)),
			lipgloss.NewStyle().Foreground(ColorMuted).Render("Press [q] to quit"),
		)
		return header + body
	}

	return ""
}
