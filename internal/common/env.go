package common

import (
	"github.com/joho/godotenv"
	"log"
)

func SetUpEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
