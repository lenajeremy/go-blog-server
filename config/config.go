package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Failed to load .env file")
		}
	}
}

func GetConfig(key string) string {
	if val, okay := os.LookupEnv(key); okay {
		log.Printf("environment variable `%s` is available", key)
		return val
	} else {
		log.Printf("environment variable `%s` is not available", key)
	}

	return ""
}
