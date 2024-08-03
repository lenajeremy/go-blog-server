package routers

import (
	"blog/middleware"
	"github.com/gofiber/fiber/v2"
)

// RouteSetupFunc Takes a router for a particular path and assigns the endpoints
// and handlers for that router.
type RouteSetupFunc = func(r *fiber.Router)

// RouteConfig Defines the structure of the routes. It should include a path,
// a function RouteSetupFunc and a boolean to indicate whether the
// route should be protected or not.
type RouteConfig struct {
	Path           string
	RouteSetupFunc RouteSetupFunc
	RequiresAuth   bool
}

func SetupRouter(app *fiber.App) {
	appRoutes := []RouteConfig{
		{"/auth", AuthRouter, false},
		{"/post", PostsRouter, true},
	}

	//app.Use(cors.New())
	//app.Use(logger.New())

	apiRoute := app.Group("/api")

	for _, rc := range appRoutes {
		var pathRouter fiber.Router = apiRoute.Group(rc.Path)

		if rc.RequiresAuth {
			pathRouter.Use(middleware.Protected())
		}

		rc.RouteSetupFunc(&pathRouter)
	}
}
