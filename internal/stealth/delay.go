package stealth

import (
	"fmt"
	"math/rand"
	"time"
)

func Think(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs) + minMs
	fmt.Printf("[DEMO] Human thinking delay: %d ms\n", delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
