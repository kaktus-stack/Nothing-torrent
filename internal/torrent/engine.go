package torrent

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/nothing-torrent/nothing-torrent/internal/config"
)

// FileEntry represents a file within the torrent for UI selection
type FileEntry struct {
	Index          int
	DisplayPath    string
	Length         int64
	BytesCompleted int64
	Selected       bool
	rawFile        *torrent.File
}

// TorrentEngine encapsulates the torrent client and active torrent
type TorrentEngine struct {
	mu           sync.RWMutex
	client       *torrent.Client
	currTorrent  *torrent.Torrent
	cfg          *config.Config
	Stealth      *StealthManager

	Name         string
	TotalLength  int64
	Files        []*FileEntry
	InfoReceived bool

	lastBytes    int64
	lastTime     time.Time
	CurrentSpeed int64 // bytes/sec
}

// NewEngine initializes the BitTorrent engine with stealth profile
func NewEngine(cfg *config.Config) (*TorrentEngine, error) {
	stealth := NewStealthManager(cfg)

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = cfg.DownloadDir
	clientConfig.PeerID = cfg.GeneratePeerID()
	clientConfig.HTTPUserAgent = cfg.GetUserAgent()
	clientConfig.NoUpload = cfg.DisableSeed

	// Hook announce requests for stealth manipulation
	clientConfig.HttpRequestDirector = stealth.InterceptRequest

	cl, err := torrent.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create torrent client: %w", err)
	}

	return &TorrentEngine{
		client:   cl,
		cfg:      cfg,
		Stealth:  stealth,
		lastTime: time.Now(),
	}, nil
}

// LoadTorrent opens a magnet link or .torrent file
func (te *TorrentEngine) LoadTorrent(source string) error {
	te.mu.Lock()
	defer te.mu.Unlock()

	var t *torrent.Torrent
	var err error

	source = strings.TrimSpace(source)
	if strings.HasPrefix(source, "magnet:") {
		t, err = te.client.AddMagnet(source)
	} else {
		// Treat as file path
		cleanPath := filepath.Clean(source)
		if strings.HasPrefix(cleanPath, "~/") {
			home, _ := os.UserHomeDir()
			cleanPath = filepath.Join(home, cleanPath[2:])
		}
		mi, loadErr := metainfo.LoadFromFile(cleanPath)
		if loadErr != nil {
			return fmt.Errorf("failed to open torrent file: %w", loadErr)
		}
		t, err = te.client.AddTorrent(mi)
	}

	if err != nil {
		return fmt.Errorf("failed to add torrent: %w", err)
	}

	te.currTorrent = t
	te.Name = t.Name()
	return nil
}

// WaitForMetadata blocks until torrent info is available or context is cancelled
func (te *TorrentEngine) WaitForMetadata(ctx context.Context) error {
	if te.currTorrent == nil {
		return fmt.Errorf("no active torrent")
	}

	select {
	case <-te.currTorrent.GotInfo():
		te.mu.Lock()
		defer te.mu.Unlock()

		te.InfoReceived = true
		te.Name = te.currTorrent.Name()
		te.TotalLength = te.currTorrent.Length()

		rawFiles := te.currTorrent.Files()
		te.Files = make([]*FileEntry, len(rawFiles))
		for i, f := range rawFiles {
			te.Files[i] = &FileEntry{
				Index:          i,
				DisplayPath:    f.DisplayPath(),
				Length:         f.Length(),
				BytesCompleted: f.BytesCompleted(),
				Selected:       true, // default: select all
				rawFile:        f,
			}
		}
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

// ApplySelection configures which files will be downloaded
func (te *TorrentEngine) ApplySelection() {
	te.mu.Lock()
	defer te.mu.Unlock()

	if te.currTorrent == nil || len(te.Files) == 0 {
		return
	}

	for _, fe := range te.Files {
		if fe.rawFile == nil {
			continue
		}
		if fe.Selected {
			fe.rawFile.SetPriority(torrent.PiecePriorityNormal)
			fe.rawFile.Download()
		} else {
			fe.rawFile.SetPriority(torrent.PiecePriorityNone)
			fe.rawFile.Cancel()
		}
	}
}

// UpdateStats recalculates download speed and file progresses
func (te *TorrentEngine) UpdateStats() (completed int64, totalSelected int64, speed int64, peers int) {
	te.mu.Lock()
	defer te.mu.Unlock()

	if te.currTorrent == nil {
		return 0, 0, 0, 0
	}

	var totalComp int64
	var totalSel int64

	for _, fe := range te.Files {
		if fe.rawFile != nil {
			fe.BytesCompleted = fe.rawFile.BytesCompleted()
			if fe.Selected {
				totalComp += fe.BytesCompleted
				totalSel += fe.Length
			}
		}
	}

	now := time.Now()
	elapsed := now.Sub(te.lastTime).Seconds()
	if elapsed >= 0.8 {
		delta := totalComp - te.lastBytes
		if delta >= 0 {
			te.CurrentSpeed = int64(float64(delta) / elapsed)
		}
		te.lastBytes = totalComp
		te.lastTime = now
	}

	conns := te.currTorrent.PeerConns()
	peers = len(conns)

	return totalComp, totalSel, te.CurrentSpeed, peers
}

// Close shuts down the torrent client
func (te *TorrentEngine) Close() {
	te.mu.Lock()
	defer te.mu.Unlock()
	if te.client != nil {
		te.client.Close()
	}
}
