package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func New(headless bool) (*Browser, error) {
	u := launcher.New().
		Headless(headless).
		Leakless(false).
		Set("disable-blink-features", "AutomationControlled").
		MustLaunch()


	b := rod.New().
		ControlURL(u).
		MustConnect()

	p := b.MustPage()

	return &Browser{
		Browser: b,
		Page:    p,
	}, nil
}
