package handlers

import (
	"blog/config"
	"blog/database"
	"blog/models"
	"blog/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

func Login(c *fiber.Ctx) error {
	type LoginFV struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var rBody LoginFV
	var user models.User
	db := database.DB

	if err := c.BodyParser(&rBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing important values",
			"data":    nil,
		})
	}

	log.Println(rBody)

	if rBody.Password == "" || rBody.Email == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Missing important values",
			"data":    nil,
		})
	}

	if err := db.Where("email = ?", rBody.Email).Find(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if ok := verifyPassword(user.Password, rBody.Password); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid login credentials",
			"data":    nil,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.GetConfig("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": map[string]string{
			"token": t,
		},
		"message": "Successfully logged in!",
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

func verifyPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"loggingOut": true,
	})
}
