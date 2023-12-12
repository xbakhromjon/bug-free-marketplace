package config

import (
	"os"
)

type Config struct {
	JwtSecretKey string
}

func NewConfig() *Config {
	var config Config

	config.JwtSecretKey = getEnv("JWT_SECRET_KEY", "")

	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}
