package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Verify *VerifyConfig
}

type VerifyConfig struct {
	SenderEmail   string
	RecieverEmail string
	Password      string
	Address       string
}

func Load() *Config {
	godotenv.Load(".env")
	return &Config{
		Verify: &VerifyConfig{
			SenderEmail:   os.Getenv("SenderEmail"),
			RecieverEmail: os.Getenv("RecieverEmail"),
			Password:      os.Getenv("Password"),
			Address:       os.Getenv("Address"),
		},
	}
}
