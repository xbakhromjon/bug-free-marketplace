package config

import (
	"os"
)

type Config struct {
	JwtSecretKey string
	RpcPort      string
}

func NewConfig() *Config {
	var config Config

	config.JwtSecretKey = getEnv("JWT_SECRET_KEY", "")
	config.RpcPort = getEnv("RPC_PORT", "5006")

	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}
