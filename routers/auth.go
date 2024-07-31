package routers

import (
	"blog/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r *fiber.Router) {
	(*r).Post("/login", handlers.Login)
	(*r).Post("/register", handlers.Register)
	(*r).Post("/logout", handlers.Logout)
}
