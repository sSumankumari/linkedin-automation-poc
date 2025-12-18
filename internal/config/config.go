package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Email      string
	Password   string
	Headless   bool
	DailyLimit int
	StartHour  int
	EndHour    int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	startHour := 0
	endHour := 23

	if v := os.Getenv("START_HOUR"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			startHour = parsed
		}
	}

	if v := os.Getenv("END_HOUR"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			endHour = parsed
		}
	}

	return &Config{
		Email:      os.Getenv("LINKEDIN_EMAIL"),
		Password:   os.Getenv("LINKEDIN_PASSWORD"),
		Headless:   os.Getenv("HEADLESS") == "true",
		DailyLimit: 10,
		StartHour:  startHour,
		EndHour:    endHour,
	}, nil
}
