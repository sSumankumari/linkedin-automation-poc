package browser

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation/internal/stealth"
)

type Browser struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func New(headless bool) (*Browser, error) {
	// ✅ CRITICAL: Disable leakless to avoid Windows Defender panic
	_ = os.Setenv("ROD_LEAKLESS", "false")

	u := launcher.New().
		Headless(headless).
		Leakless(false).
		Set("disable-blink-features", "AutomationControlled").
		Set("disable-infobars").
		Set("disable-dev-shm-usage").
		Set("no-sandbox").
		MustLaunch()

	println("[DEMO] Chrome launcher started")

	b := rod.New().
		ControlURL(u).
		MustConnect()

	println("[DEMO] Connected to Chrome instance")

	p := b.MustPage()

	// 🎭 Random viewport (demo stealth)
	width := rand.Intn(200) + 1200
	height := rand.Intn(200) + 700

	p.MustSetViewport(width, height, 1, false)

	println("[DEMO] Viewport randomized:", width, "x", height)

	// 🧬 Fingerprint masking (SAFE DEMO)
	stealth.MaskFingerprint(p)

	println("[DEMO] Fingerprint masking applied")

	time.Sleep(2 * time.Second)

	return &Browser{
		Browser: b,
		Page:    p,
	}, nil
}
