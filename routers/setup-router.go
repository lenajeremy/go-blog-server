package routers

import "github.com/gofiber/fiber/v2"

type RouteSetupHandler = func(r *fiber.Router)

func SetupRouter(app *fiber.App) {

	appRoutes := map[string]RouteSetupHandler{
		"/auth": AuthRouter,
		"/post": PostsRouter,
	}

	apiRoute := app.Group("/api")

	for path, Router := range appRoutes {
		pathRoute := apiRoute.Group(path)
		Router(&pathRoute)
	}
}
