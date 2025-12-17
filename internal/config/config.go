package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email       string
	Password    string
	Headless    bool
	DailyLimit  int
	StartHour   int
	EndHour     int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		Email:      os.Getenv("LINKEDIN_EMAIL"),
		Password:   os.Getenv("LINKEDIN_PASSWORD"),
		Headless:   os.Getenv("HEADLESS") == "true",
		DailyLimit: 10,
		StartHour:  10,
		EndHour:    18,
	}, nil
}
