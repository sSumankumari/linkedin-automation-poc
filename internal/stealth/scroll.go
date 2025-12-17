package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
)

// RandomScroll simulates natural scrolling behavior
func RandomScroll(page *rod.Page) {
	scroll := rand.Intn(500) + 200

	page.MustEval(`
		(scrollY) => {
			window.scrollBy(0, scrollY);
		}
	`, scroll)
}
