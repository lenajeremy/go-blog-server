package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func GetConfig(key string) string {
	if key, okay := os.LookupEnv(key); okay {
		return key
	}

	return ""
}
