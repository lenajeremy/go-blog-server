package main

import (
	"blog/database"
	"blog/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {

	database.ConnectDB(database.ConnectDBConfig{
		MakeMigrations: true,
	})

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("static", "public")

	routers.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
