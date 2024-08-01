package routers

import (
	"blog/handlers"
	"github.com/gofiber/fiber/v2"
)

func PostsRouter(r *fiber.Router) {
	(*r).Post("/", handlers.CreatePost)
	(*r).Get("/", handlers.ListPersonalPosts)
	(*r).Get("/all", handlers.ListPosts)

	(*r).Get("/:id", handlers.ViewPost)
	(*r).Patch("/:id", handlers.EditPost)
	(*r).Delete("/:id", handlers.DeletePost)

	(*r).Get("/:postId/comments", handlers.ViewPostComments)
	(*r).Post("/:postId/comments", handlers.AddComment)
	(*r).Patch("/:postId/comments/:commentId", handlers.EditComment)
	(*r).Delete("/:postId/comments/:commentId", handlers.DeleteComment)
}
