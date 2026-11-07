package ui

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
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
	StateSettings
	StateError
)

type tickMsg time.Time
type metadataReadyMsg struct{}
type metadataErrMsg struct{ err error }

// Model is the main Bubble Tea application state
type Model struct {
	state          AppState
	prevState      AppState
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

	// Settings state
	settingsCursor SettingsField

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
		speedHistory: make([]int64, 0, 50),
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

// Update handles incoming events and keyboard input with layout-independent key mapping
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
		// Normalize key so that Russian/Cyrillic layout triggers same action as QWERTY
		rawKey := msg.String()
		actionKey := NormalizeKey(rawKey)

		// Global quit
		if rawKey == "ctrl+c" || actionKey == "q" {
			m.engine.Close()
			return m, tea.Quit
		}

		// Settings view handlers
		if m.state == StateSettings {
			switch rawKey {
			case "esc":
				m.state = m.prevState
				_ = m.cfg.Save()
				return m, nil
			}

			switch actionKey {
			case "s", "esc":
				m.state = m.prevState
				_ = m.cfg.Save()
				return m, nil

			case "up", "k":
				if m.settingsCursor > 0 {
					m.settingsCursor--
				}
				return m, nil

			case "down", "j":
				if m.settingsCursor < FieldCount-1 {
					m.settingsCursor++
				}
				return m, nil

			case "left", "h":
				m.adjustSetting(false)
				return m, nil

			case "right", "l", " ":
				m.adjustSetting(true)
				return m, nil

			case "enter":
				m.adjustSetting(true)
				return m, nil
			}
			return m, nil
		}

		// Standard views handlers
		switch rawKey {
		case "up":
			actionKey = "k"
		case "down":
			actionKey = "j"
		case "left":
			actionKey = "h"
		case "right":
			actionKey = "l"
		case "enter":
			actionKey = "enter"
		case " ":
			actionKey = " "
		case "esc":
			actionKey = "esc"
		}

		switch actionKey {
		case "k": // Up
			if m.state == StateFileSelect {
				if m.cursor > 0 {
					m.cursor--
					if m.cursor < m.scrollOffset {
						m.scrollOffset = m.cursor
					}
				}
			}

		case "j": // Down
			if m.state == StateFileSelect {
				if m.cursor < len(m.engine.Files)-1 {
					m.cursor++
					maxVisible := m.height - 10
					if maxVisible < 4 {
						maxVisible = 4
					}
					if m.cursor >= m.scrollOffset+maxVisible {
						m.scrollOffset = m.cursor - maxVisible + 1
					}
				}
			}

		case " ": // Space
			if m.state == StateFileSelect && len(m.engine.Files) > 0 {
				m.engine.Files[m.cursor].Selected = !m.engine.Files[m.cursor].Selected
			}

		case "a": // Select all
			if m.state == StateFileSelect {
				for _, f := range m.engine.Files {
					f.Selected = true
				}
			}

		case "n": // Select none
			if m.state == StateFileSelect {
				for _, f := range m.engine.Files {
					f.Selected = false
				}
			}

		case "i": // Invert selection
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

		case "p": // Pause / Resume
			if m.state == StateDownloading {
				m.isPaused = !m.isPaused
			}

		case "s": // Open Settings
			if m.state == StateDownloading || m.state == StateFileSelect {
				m.prevState = m.state
				m.state = StateSettings
			}

		case "o": // Open download directory
			openDirectory(m.cfg.DownloadDir)
		}
	}

	return m, nil
}

func openDirectory(dir string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", dir)
	case "darwin":
		cmd = exec.Command("open", dir)
	default:
		cmd = exec.Command("xdg-open", dir)
	}
	_ = cmd.Start()
}

func (m *Model) adjustSetting(forward bool) {
	switch m.settingsCursor {
	case FieldStealth:
		m.cfg.StealthMode = !m.cfg.StealthMode

	case FieldTargetRatio:
		if forward {
			m.cfg.TargetRatio += 0.05
			if m.cfg.TargetRatio > 1.50 {
				m.cfg.TargetRatio = 1.50
			}
		} else {
			m.cfg.TargetRatio -= 0.05
			if m.cfg.TargetRatio < 0.50 {
				m.cfg.TargetRatio = 0.50
			}
		}

	case FieldClientProfile:
		profiles := []string{"qbittorrent", "transmission", "deluge"}
		currIdx := 0
		for i, p := range profiles {
			if p == m.cfg.DefaultClient {
				currIdx = i
				break
			}
		}
		if forward {
			currIdx = (currIdx + 1) % len(profiles)
		} else {
			currIdx = (currIdx - 1 + len(profiles)) % len(profiles)
		}
		m.cfg.DefaultClient = profiles[currIdx]

	case FieldDisableSeed:
		m.cfg.DisableSeed = !m.cfg.DisableSeed

	case FieldUploadSpeed:
		if m.cfg.MinUploadKBps <= 60 {
			m.cfg.MinUploadKBps = 90
			m.cfg.MaxUploadKBps = 220
		} else if m.cfg.MinUploadKBps <= 90 {
			m.cfg.MinUploadKBps = 180
			m.cfg.MaxUploadKBps = 350
		} else {
			m.cfg.MinUploadKBps = 50
			m.cfg.MaxUploadKBps = 120
		}
	}
	_ = m.cfg.Save()
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

	case StateSettings:
		return RenderSettings(
			m.cfg,
			m.settingsCursor,
			m.width,
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
