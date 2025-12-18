package auth

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/rs/zerolog"

	"linkedin-automation/internal/stealth"
)

type Service struct {
	Page   *rod.Page
	Logger zerolog.Logger
}

func (a *Service) Login(email, password string) error {
	a.Logger.Info().Msg("[DEMO] Starting login simulation")

	a.Page.MustNavigate("https://www.linkedin.com/login")
	a.Page.MustWaitLoad()

	emailEl := a.Page.MustElement(`#username`)
	passEl := a.Page.MustElement(`#password`)

	stealth.Think(800, 1500)
	stealth.TypeHuman(emailEl, email)

	stealth.Think(500, 1000)
	stealth.TypeHuman(passEl, password)

	fmt.Println("[DEMO] Typed dummy credentials using human typing")

	stealth.Think(1000, 2000)
	fmt.Println("[DEMO] Login button NOT clicked (ToS safe)")

	a.Logger.Info().Msg("[DEMO] Login flow demonstrated successfully")
	return nil
}
