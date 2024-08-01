package handlers

import (
	"blog/database"
	"blog/models"
	"blog/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"firstName": "jeremiah",
	})
}

func Register(c *fiber.Ctx) error {
	type RFValues struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	db := database.DB
	var user models.User
	var uProf models.Profile
	var rBody RFValues

	if err := c.BodyParser(&rBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "Missing form values. Please fill all the fields",
		})
	}

	if rBody.LastName == "" || rBody.FirstName == "" || rBody.Email == "" || rBody.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "Missing form values. Please fill all the fields",
		})
	}

	passwordHash, err := hashPassword(rBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash user password",
			"data":    nil,
			"success": false,
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		user = models.User{
			Email:         rBody.Email,
			Password:      passwordHash,
			EmailVerified: false,
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		uProf = models.Profile{
			FirstName: &rBody.FirstName,
			LastName:  &rBody.LastName,
			UserId:    user.ID,
		}

		if err := tx.Create(&uProf).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to register user",
			"data":    err.Error(),
			"success": false,
		})
	}

	publicUser, err := utils.StructToMap(models.PublicUser{
		ID:        user.ID.String(),
		FirstName: *uProf.FirstName,
		LastName:  *uProf.LastName,
		Email:     user.Email,
		Verified:  user.EmailVerified,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User has been created successfully",
		"data":    publicUser,
	})
}

func hashPassword(password string) (string, error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passByte), err
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"loggingOut": true,
	})
}
