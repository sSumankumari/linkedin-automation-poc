package search

import (
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	Logger zerolog.Logger
}

func (s *Service) CollectProfiles() []string {
	s.Logger.Info().Msg("Simulating search and profile collection")
	time.Sleep(2 * time.Second)

	return []string{
		"https://linkedin.com/in/example1",
		"https://linkedin.com/in/example2",
	}
}
