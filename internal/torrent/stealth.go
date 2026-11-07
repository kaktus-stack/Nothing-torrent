package torrent

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/nothing-torrent/nothing-torrent/internal/config"
)

// StealthManager handles client spoofing and ratio protection
type StealthManager struct {
	mu           sync.Mutex
	cfg          *config.Config
	startTime    time.Time
	lastReport   time.Time
	fakeUploaded int64
	lastDown     int64
	rnd          *rand.Rand
}

// NewStealthManager creates a new stealth tracker manager
func NewStealthManager(cfg *config.Config) *StealthManager {
	return &StealthManager{
		cfg:       cfg,
		startTime: time.Now(),
		rnd:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// InterceptRequest modifies tracker announce HTTP requests
func (sm *StealthManager) InterceptRequest(req *http.Request) error {
	if req == nil || req.URL == nil {
		return nil
	}

	// 1. Spoof User-Agent for all tracker communications
	req.Header.Set("User-Agent", sm.cfg.GetUserAgent())
	req.Header.Set("Accept-Encoding", "gzip")

	// 2. Check if this is an announce request
	if !strings.Contains(req.URL.Path, "announce") && !strings.Contains(req.URL.RawQuery, "info_hash") {
		return nil
	}

	if !sm.cfg.StealthMode {
		return nil
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	query := req.URL.Query()
	downStr := query.Get("downloaded")
	if downStr == "" {
		return nil
	}

	downloaded, err := strconv.ParseInt(downStr, 10, 64)
	if err != nil || downloaded <= 0 {
		return nil
	}

	sm.lastDown = downloaded

	// Target ratio with a small realistic jitter (+/- 2.5%)
	// so ratio stays between ~0.82 and 0.88 instead of exact fixed float
	jitter := 0.975 + sm.rnd.Float64()*0.05
	targetRatio := sm.cfg.TargetRatio * jitter
	if targetRatio < 0.78 {
		targetRatio = 0.80
	} else if targetRatio > 0.92 {
		targetRatio = 0.90
	}

	calcUploaded := int64(float64(downloaded) * targetRatio)
	if calcUploaded > sm.fakeUploaded {
		sm.fakeUploaded = calcUploaded
	}

	// Replace the 'uploaded' query param with our stealth stat
	query.Set("uploaded", strconv.FormatInt(sm.fakeUploaded, 10))

	// Ensure spoofed peer_id if needed
	if query.Get("peer_id") == "" {
		query.Set("peer_id", sm.cfg.GeneratePeerID())
	}

	req.URL.RawQuery = query.Encode()
	sm.lastReport = time.Now()

	return nil
}

// GetCurrentStats returns live stealth statistics
func (sm *StealthManager) GetCurrentStats(realDownloaded int64) (fakeUploaded int64, currentRatio float64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if realDownloaded <= 0 {
		return 0, 0.0
	}

	ratio := sm.cfg.TargetRatio
	uploaded := int64(float64(realDownloaded) * ratio)
	if uploaded < sm.fakeUploaded {
		uploaded = sm.fakeUploaded
	}

	calculatedRatio := float64(uploaded) / float64(realDownloaded)
	return uploaded, calculatedRatio
}

// GetSimulatedUploadSpeed calculates a realistic upload speed display (KB/s)
func (sm *StealthManager) GetSimulatedUploadSpeed(isDownloading bool) int {
	if !sm.cfg.StealthMode || !isDownloading {
		return 0
	}
	min := sm.cfg.MinUploadKBps
	max := sm.cfg.MaxUploadKBps
	if max <= min {
		max = min + 50
	}
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return min + sm.rnd.Intn(max-min+1)
}

// FormatRatio returns formatted ratio string (e.g. "0.85x")
func FormatRatio(ratio float64) string {
	if ratio <= 0 {
		return "0.00x"
	}
	return fmt.Sprintf("%.2fx", ratio)
}
