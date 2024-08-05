package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {

}

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Failed to load .env file")
		}
	}
}

func GetConfig(key string) string {
	if key, okay := os.LookupEnv(key); okay {
		return key
	}

	return ""
}
