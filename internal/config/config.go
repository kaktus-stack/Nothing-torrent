package config

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// ClientProfile defines tracker spoofing parameters
type ClientProfile struct {
	Name        string `json:"name"`
	PeerIDPrefix string `json:"peer_id_prefix"` // e.g. "-qB4650-"
	UserAgent   string `json:"user_agent"`     // e.g. "qBittorrent/4.6.5"
}

var AvailableProfiles = map[string]ClientProfile{
	"qbittorrent": {
		Name:         "qBittorrent 4.6.5",
		PeerIDPrefix: "-qB4650-",
		UserAgent:    "qBittorrent/4.6.5",
	},
	"transmission": {
		Name:         "Transmission 4.0.5",
		PeerIDPrefix: "-TR4050-",
		UserAgent:    "Transmission/4.0.5",
	},
	"deluge": {
		Name:         "Deluge 2.1.1",
		PeerIDPrefix: "-DE2110-",
		UserAgent:    "Deluge/2.1.1",
	},
}

// Config stores user settings
type Config struct {
	DownloadDir   string  `json:"download_dir"`
	DefaultClient string  `json:"default_client"` // "qbittorrent", "transmission", "deluge"
	StealthMode   bool    `json:"stealth_mode"`   // Enable fake ratio reporting
	TargetRatio   float64 `json:"target_ratio"`   // Default: 0.85 (range 0.80 - 0.90)
	MinUploadKBps int     `json:"min_upload_kbps"` // Min simulated upload speed (e.g. 80)
	MaxUploadKBps int     `json:"max_upload_kbps"` // Max simulated upload speed (e.g. 240)
	DisableSeed   bool    `json:"disable_seed"`    // Turn off actual upload to save router/bandwidth
}

// DefaultConfig returns sane defaults
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	dlDir := filepath.Join(home, "Downloads")
	return &Config{
		DownloadDir:   dlDir,
		DefaultClient: "qbittorrent",
		StealthMode:   true,
		TargetRatio:   0.85,
		MinUploadKBps: 90,
		MaxUploadKBps: 220,
		DisableSeed:   true, // True cold-router mode: zero real upload, pure stealth reports
	}
}

// ConfigPath returns path to ~/.config/nothing-torrent/config.json
func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "nothing-torrent", "config.json")
}

// Load loads config from disk or returns default
func Load() (*Config, error) {
	path := ConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		cfg := DefaultConfig()
		_ = cfg.Save()
		return cfg, nil
	}
	cfg := DefaultConfig()
	if err := json.Unmarshal(data, cfg); err != nil {
		return cfg, nil
	}
	return cfg, nil
}

// Save saves config to disk
func (c *Config) Save() error {
	path := ConfigPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GeneratePeerID creates a believable 20-byte peer_id
func (c *Config) GeneratePeerID() string {
	profile, ok := AvailableProfiles[c.DefaultClient]
	if !ok {
		profile = AvailableProfiles["qbittorrent"]
	}

	prefix := profile.PeerIDPrefix
	remaining := 20 - len(prefix)
	if remaining <= 0 {
		return prefix[:20]
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := make([]byte, remaining)
	for i := range suffix {
		suffix[i] = letters[r.Intn(len(letters))]
	}

	return prefix + string(suffix)
}

// GetUserAgent returns the spoofed User-Agent string
func (c *Config) GetUserAgent() string {
	profile, ok := AvailableProfiles[c.DefaultClient]
	if !ok {
		profile = AvailableProfiles["qbittorrent"]
	}
	return profile.UserAgent
}
