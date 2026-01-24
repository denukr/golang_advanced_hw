package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Verify *VerifyConfig
}

type VerifyConfig struct {
	Email    string
	Password string
	Address  string
}

func Load() *Config {
	godotenv.Load(".env")
	return &Config{
		Verify: &VerifyConfig{
			Email:    os.Getenv("Email"),
			Password: os.Getenv("Password"),
			Address:  os.Getenv("Address"),
		},
	}
}
