package messaging

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (m *Service) Send(profileURL, message string) {
	m.Logger.Info().
		Str("profile", profileURL).
		Msg("[DEMO] Message send simulated")

	fmt.Println("[DEMO] Message NOT sent (simulation only)")
	time.Sleep(1 * time.Second)
}
