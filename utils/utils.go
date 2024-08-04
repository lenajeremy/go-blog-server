package utils

import (
	"blog/database"
	"blog/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func StructToMap(str any) (data map[string]any, err error) {
	data = map[string]any{}

	// convert the struct to json byte
	jsonByte, err := json.Marshal(str)

	// return if there's any error while converting
	if err != nil {
		return
	}

	// convert the jsonByte to a map[string]any
	err = json.Unmarshal(jsonByte, &data)

	return
}

func GetUserFromContext(c *fiber.Ctx) (user models.User, err error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid, _ := uuid.Parse(claims["id"].(string))
	err = database.DB.First(&user, uid).Error

	return
}

func RespondSuccess(c *fiber.Ctx, message string, data any) error {
	return c.JSON(fiber.Map{
		"data":    data,
		"message": message,
		"success": true,
	})
}

func RespondError(c *fiber.Ctx, errorCode int, errorMessage string) error {
	return c.Status(errorCode).JSON(fiber.Map{
		"data":    nil,
		"message": errorMessage,
		"success": false,
	})
}
