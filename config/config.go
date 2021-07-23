package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DEFAULT_PORT = "8000"
)

type Config struct {
	ServerPort string
}

func init() {
	godotenv.Load()
}

func GetConfig() *Config {
	return &Config{
		ServerPort: getPort(),
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return ":" + port
}
