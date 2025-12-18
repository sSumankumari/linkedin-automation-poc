package stealth

import (
	"fmt"
	"math/rand"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	scroll := rand.Intn(400) + 200
	fmt.Printf("[DEMO] Scrolling page by %dpx\n", scroll)

	page.MustEval(`y => window.scrollBy(0, y)`, scroll)
}
