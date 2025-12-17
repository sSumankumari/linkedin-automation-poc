package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func TypeHuman(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(rand.Intn(120)+50) * time.Millisecond)
	}
}
