package connect

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (c *Service) Send(profileURL string) {
	c.Logger.Info().
		Str("profile", profileURL).
		Msg("[DEMO] Connection request simulated")

	fmt.Println("[DEMO] Connect button NOT clicked (ToS safe)")
	time.Sleep(1 * time.Second)
}
