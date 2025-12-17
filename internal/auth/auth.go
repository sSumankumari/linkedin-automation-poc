package auth

import (
	"errors"

	"github.com/go-rod/rod"
	"github.com/rs/zerolog"
)

type Service struct {
	Page   *rod.Page
	Logger zerolog.Logger
}

func (a *Service) Login(email, password string) error {
	a.Logger.Info().Msg("Starting login flow")

	a.Page.MustNavigate("https://www.linkedin.com/login")
	a.Page.MustWaitLoad()

	if a.Page.MustHas("input[name=session_key]") == false {
		return errors.New("login form not detected")
	}

	return nil
}
