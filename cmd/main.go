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
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New()
	logg.Info().Msg("LinkedIn Automation POC started")

	if !scheduler.Allowed(cfg.StartHour, cfg.EndHour) {
		logg.Warn().Msg("Outside allowed demo window")
		return
	}

	br, err := browser.New(cfg.Headless)
	if err != nil {
		log.Fatal(err)
	}

	// Warm-up behavior
	br.Page.MustNavigate("https://www.linkedin.com")
	br.Page.MustWaitLoad()

	stealth.Think(800, 1200)
	stealth.MoveMouseHuman(br.Page, 500, 350)
	stealth.RandomScroll(br.Page)

	authService := auth.Service{Page: br.Page, Logger: logg}
	searchService := search.Service{Logger: logg}
	connectService := connect.Service{Logger: logg}

	if err := authService.Login(cfg.Email, cfg.Password); err != nil {
		logg.Error().Err(err).Msg("Login demo failed")
		return
	}

	profiles := searchService.CollectProfiles()

	for _, p := range profiles {
		logg.Info().Str("profile", p).Msg("Processing profile")
		stealth.Think(1200, 2000)
		connectService.Send(p)
		time.Sleep(2 * time.Second)
	}

	logg.Info().Msg("Demo execution completed successfully")
}
