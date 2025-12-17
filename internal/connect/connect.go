package connect

import (
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (c *Service) Send(profileURL string) {
	c.Logger.Info().
		Str("profile", profileURL).
		Msg("Simulating connection request")

	time.Sleep(1 * time.Second)
}
