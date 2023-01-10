package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	AUTH_ADDRESS string
	// SERVER_ADDRESS string

	DB_USER_NAME string
	DB_PASSWORD  string
	DB_SSLMODE   string
	DB_HOST      string
	DB_PORT      string
	DB_NAME      string
	SIGN         string
}

func init() {
	godotenv.Load()
}

func NewConfigs() *Configs {
	return &Configs{
		AUTH_ADDRESS: os.Getenv("AUTH_ADDRESS"),
		// SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		DB_USER_NAME: os.Getenv("DB_USER_NAME"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_NAME:      os.Getenv("DB_NAME"),
		SIGN:         os.Getenv("SIGN"),
	}
}
