package main

import (
	"log"
	"time"

	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/config"
	"linkedin-automation/internal/connect"
	"linkedin-automation/internal/logger"
	"linkedin-automation/internal/scheduler"
	"linkedin-automation/internal/search"
	"linkedin-automation/internal/stealth"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New()
	logg.Info().Msg("Application started")

	// Schedule check
	if !scheduler.Allowed(cfg.StartHour, cfg.EndHour) {
		logg.Warn().Msg("Outside allowed execution window. Exiting safely.")
		return
	}

	// Launch browser
	br, err := browser.New(cfg.Headless)
	if err != nil {
		log.Fatal(err)
	}

	// Initial human-like behavior
	br.Page.MustNavigate("https://example.com")
	br.Page.MustWaitLoad()

	stealth.Think(800, 1500)
	stealth.MoveMouseHuman(br.Page, 400, 300)
	stealth.RandomScroll(br.Page)

	// Initialize services
	authService := auth.Service{Page: br.Page, Logger: logg}
	searchService := search.Service{Logger: logg}
	connectService := connect.Service{Logger: logg}

	// Login with detection
	if err := authService.Login(cfg.Email, cfg.Password); err != nil {
		logg.Error().Err(err).Msg("Login failed or checkpoint detected")
		return
	}

	stealth.Think(1000, 2000)

	// Search simulation
	profiles := searchService.CollectProfiles()
	logg.Info().Int("count", len(profiles)).Msg("Profiles collected")

	// Connection workflow
	for _, p := range profiles {
		logg.Info().Str("profile", p).Msg("Processing profile")

		stealth.Think(1500, 3000)
		stealth.MoveMouseHuman(br.Page, 600, 350)

		connectService.Send(p)

		// Cooldown between actions
		time.Sleep(2 * time.Second)
	}

	logg.Info().Msg("Execution completed")
}
