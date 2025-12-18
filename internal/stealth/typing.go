package stealth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func TypeHuman(el *rod.Element, text string) {
	fmt.Println("[DEMO] Human-like typing started")
	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(rand.Intn(120)+60) * time.Millisecond)
	}
}
