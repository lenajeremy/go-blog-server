package main

import (
	"blog/database"
	"blog/routers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	database.ConnectDB(database.ConnectDBConfig{
		MakeMigrations: true,
	})

	app := fiber.New()

	routers.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
