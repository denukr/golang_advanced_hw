package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db Db
}

type Db struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Db{
			DSN: os.Getenv("DSN"),
		},
	}
}
