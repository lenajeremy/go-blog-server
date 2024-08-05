package database

import (
	"blog/config"
	"blog/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type ConnectDBConfig struct {
	MakeMigrations bool
}

func ConnectDB(dbConfig ConnectDBConfig) {
	config.LoadEnv()

	var (
		DbName = config.GetConfig("PGDATABASE")
		DbHost = config.GetConfig("PGHOST")
		DbPass = config.GetConfig("PGPASSWORD")
		DbPort = config.GetConfig("PGPORT")
		DbUser = config.GetConfig("PGUSER")
	)

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPass, DbPort, DbName)
	pg := postgres.Open(dsn)

	DB, err = gorm.Open(pg, &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		log.Println("Connected to database successfully")
	}

	if dbConfig.MakeMigrations {
		runMigrations()
	}
}

func runMigrations() {
	err := DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Comment{}, &models.Post{})
	log.Println("Successfully created database migrations")

	if err != nil {
		log.Fatal(err.Error())
	}
}
