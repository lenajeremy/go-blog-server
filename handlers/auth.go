package handlers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"firstName": "jeremiah",
	})
}

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"lastName": "lena",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"loggingOut": true,
	})
}
