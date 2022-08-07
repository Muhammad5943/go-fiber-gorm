package authentication

import (
	"strings"

	"github.com/Muhammad5943/go-fiber-gorm/database"
	"github.com/Muhammad5943/go-fiber-gorm/models/entity"
	"github.com/Muhammad5943/go-fiber-gorm/models/request"
	"github.com/Muhammad5943/go-fiber-gorm/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "register",
	})
}

func Login(c *fiber.Ctx) error {
	login := new(request.LoginRequest)
	if err := c.BodyParser(login); err != nil {
		return err
	}

	// validate login request
	validate := validator.New()
	errValidate := validate.Struct(login)
	if errValidate != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "You might have bad request",
			"error":   strings.Split(errValidate.Error(), "Key"),
		})
	}

	// chack available user
	var user entity.User
	err := database.DB.First(&user, "email=?", login.Email).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "wrong credential",
		})
	}

	// check validation password
	isValid := utils.CheckPassword(login.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "you input wrong password",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"token": "secret",
	})
}
