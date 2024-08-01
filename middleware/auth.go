package middleware

import (
	"blog/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
)

var Protected = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		Key: []byte(config.GetConfig("JWT_SECRET")),
	},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, please login before you can access this endpoint",
			"data":    nil,
		})
	},
})
