package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// MoveMouseHuman simulates gradual human-like mouse movement
func MoveMouseHuman(page *rod.Page, targetX, targetY float64) {
	steps := rand.Intn(15) + 15

	startX := rand.Float64() * 200
	startY := rand.Float64() * 200

	for i := 0; i <= steps; i++ {
		progress := float64(i) / float64(steps)

		x := startX + (targetX-startX)*progress + rand.Float64()*2
		y := startY + (targetY-startY)*progress + rand.Float64()*2

		page.Mouse.MoveTo(proto.Point{
			X: x,
			Y: y,
		})

		time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
	}
}
