package messaging

import (
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (m *Service) Send(profileURL, message string) {
	m.Logger.Info().
		Str("profile", profileURL).
		Msg("Simulating message send")

	time.Sleep(1 * time.Second)
}
