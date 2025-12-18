package stealth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func MoveMouseHuman(page *rod.Page, targetX, targetY float64) {
	fmt.Println("[DEMO] Human-like mouse movement")

	startX := rand.Float64() * 300
	startY := rand.Float64() * 300
	steps := rand.Intn(20) + 20

	for i := 0; i <= steps; i++ {
		progress := float64(i) / float64(steps)

		x := startX + (targetX-startX)*progress + rand.Float64()*3
		y := startY + (targetY-startY)*progress + rand.Float64()*3

		page.Mouse.MoveTo(proto.Point{X: x, Y: y})
		time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
	}
}
