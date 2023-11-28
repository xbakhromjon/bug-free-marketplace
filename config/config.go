package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment               string // develop, staging, production
	CtxTimeout                int    // context timeout in seconds
	LogLevel                  string
	HTTPPort                  string
	PostgresHost              string
	PostgresPort              string
	PostgresDatabase          string
	PostgresUser              string
	PostgresPassword          string
	PostgresConnectionTimeOut int // seconds
	PostgresConnectionTry     int
}

// Load loads environment vars and inflates Config
func Load(dotenvPath string) Config {

	err := godotenv.Load(dotenvPath)

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	c := Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", "5005"))
	c.CtxTimeout = cast.ToInt(GetOrReturnDefault("CTX_TIMEOUT", 7))
	// Postgres
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "marketplace_db"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "root"))
	c.PostgresConnectionTimeOut = cast.ToInt(GetOrReturnDefault("POSTGRES_CONNECTION_TIMEOUT", 5))
	c.PostgresConnectionTry = cast.ToInt(GetOrReturnDefault("POSTGRES_CONNECTION_TRY", 10))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
