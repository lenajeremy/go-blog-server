package handlers

import "github.com/gofiber/fiber/v2"

func CreatePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"creating": false,
	})
}

func EditPost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"editing": true,
	})
}

func DeletePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"deleting": true,
	})
}

func ViewPost(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"viewing": true,
		"id":      id,
	})
}

func ListPosts(c *fiber.Ctx) error {
	page, limit, category, tags :=
		c.Query("page"),
		c.Query("limit"),
		c.Query("category"),
		c.Query("tags")

	return c.JSON(fiber.Map{
		"lists":    []any{},
		"page":     page,
		"limit":    limit,
		"category": category,
		"tags":     tags,
	})
}

func AddComment(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"adding": true,
	})
}

func EditComment(c *fiber.Ctx) error {
	commentId := c.Params("id")
	return c.JSON(fiber.Map{
		"commentId": commentId,
		"editing":   true,
	})
}

func DeleteComment(c *fiber.Ctx) error {
	commentId := c.Params("id")
	return c.JSON(fiber.Map{
		"commentId": commentId,
		"deleting":  true,
	})
}

func ViewPostComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	return c.JSON(fiber.Map{
		"postId":   postId,
		"comments": []any{},
	})
}
