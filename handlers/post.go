package handlers

import (
	"blog/database"
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreatePost(c *fiber.Ctx) error {
	type PostFV struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Content  string `json:"content"`
	}

	var post models.Post
	var user models.User
	var err error
	var rBody PostFV
	db := database.DB

	if err := c.BodyParser(&rBody); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	if rBody.Subtitle == "" || rBody.Title == "" || rBody.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Missing important values",
		})
	}

	user, err = utils.GetUserFromContext(c)

	if err != nil {
		log.Fatal("This should not happen, user should be logged in!")
	}

	post = models.Post{
		Title:    rBody.Title,
		SubTitle: rBody.Subtitle,
		Content:  rBody.Content,
		AuthorID: user.ID,
		Comments: make([]models.Comment, 0),
	}

	err = db.Create(&post).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": fmt.Sprintf("Unable to create post due to %s", err.Error()),
		})
	}

	if p, err := utils.StructToMap(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	} else {
		return c.JSON(fiber.Map{
			"success": true,
			"data":    p,
			"message": "Successfully created post",
		})
	}
}

func ListPosts(c *fiber.Ctx) error {
	page, limit := c.Query("page"), c.Query("limit")
	var posts []models.Post

	if err := database.DB.Preload("Comments").Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"data":    posts,
		"success": true,
		"message": "Successfully retried posts",
		"page":    page,
		"limit":   limit,
		"total":   len(posts),
	})
}

func ListPersonalPosts(c *fiber.Ctx) error {
	page, limit := c.Query("page"), c.Query("limit")
	user, err := utils.GetUserFromContext(c)

	if err != nil {
		log.Panic("this should not happen. there should be a user here")
	}

	log.Println(user)

	posts := []models.Post{}

	if err := database.DB.Preload("Comments").Where("author_id = ?", user.ID).Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"data":    posts,
		"success": true,
		"message": "Successfully retrieved user posts",
		"total":   len(posts),
		"page":    page,
		"limit":   limit,
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
