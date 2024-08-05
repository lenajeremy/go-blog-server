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
	if key, okay := os.LookupEnv(key); okay {
		log.Printf("environment variable `%s` is available", key)
		return key
	} else {
		log.Printf("environment variable `%s` is not available", key)
	}

	return ""
}

// postgresql://postgres:NYelohjGWNgOPVzmcJFHWfbIYbIwgRTy@viaduct.proxy.rlwy.net:26538/railway
