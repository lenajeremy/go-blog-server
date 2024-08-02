package handlers

import (
	"blog/database"
	"blog/models"
	"blog/utils"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	if err := database.DB.Preload("Comments").Order("created_at DESC").Find(&posts).Error; err != nil {
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

	var posts []models.Post

	if err := database.DB.Preload("Comments").Order("created_at DESC").Where("author_id = ?", user.ID).Find(&posts).Error; err != nil {
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
	type editFormV struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Content  string `json:"content"`
	}

	postId := c.Params("id")
	var formV editFormV
	var post models.Post
	var err error
	var user models.User

	if err = c.BodyParser(&formV); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	user, err = utils.GetUserFromContext(c)

	if err != nil {
		log.Panic("This shouldn't happen, user should be logged in")
	}

	err = database.DB.Where("id = ? and author_id = ?", postId, user.ID).First(&post).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Blog with id %s does not exist", postId),
				"data":    nil,
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
				"data":    nil,
			})
		}
	}

	log.Println(formV)

	if formV.Title != "" {
		post.Title = formV.Title
	}

	if formV.Subtitle != "" {
		post.SubTitle = formV.Subtitle
	}

	if formV.Content != "" {
		post.Content = formV.Content
	}

	err = database.DB.Save(&post).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully edited post",
		"data":    post,
	})
}

func DeletePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Post

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Fatal("this should not happen! User should be logged in!")
	}

	err = database.DB.Where("id = ? and author_id", postId, user.ID).First(&post).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	err = database.DB.Delete(&post).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully deleted post",
		"data":    post,
	})
}

func ViewPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	if err := database.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Can't find post with ID:" + id,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "returning post",
		"data":    post,
	})
}

func AddComment(c *fiber.Ctx) error {
	type CommentFV struct {
		Content string `json:"content"`
	}

	postId := c.Params("postId")
	var commFV CommentFV
	var post models.Post
	var user models.User
	var err error

	user, err = utils.GetUserFromContext(c)

	if err != nil {
		log.Fatal("This should not happen, user should be logged in")
	}

	if err = c.BodyParser(&commFV); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	err = database.DB.Where("id = ?", postId).First(&post).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	comment := models.Comment{
		Content:         commFV.Content,
		PostCommentedOn: post.ID,
		AuthorID:        user.ID,
	}

	if err = database.DB.Create(comment).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    comment,
		"message": "Successfully commented on post",
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
	comments := make([]models.Comment, 0)

	err := database.DB.Where("post_commented_on = ?", postId).Order("created_at DESC").Find(&comments)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully retrieved comment",
		"data": fiber.Map{
			"comments": comments,
		},
	})
}
