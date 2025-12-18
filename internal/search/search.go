package search

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (s *Service) CollectProfiles() []string {
	s.Logger.Info().Msg("[DEMO] Simulating LinkedIn search")
	time.Sleep(2 * time.Second)

	fmt.Println("[DEMO] Profiles parsed from DOM (simulated)")

	return []string{
		"https://linkedin.com/in/demo-profile-1",
		"https://linkedin.com/in/demo-profile-2",
	}
}
