package authentication

import (
	"log"
	"strings"
	"time"

	"github.com/Muhammad5943/go-fiber-gorm/database"
	"github.com/Muhammad5943/go-fiber-gorm/models/entity"
	"github.com/Muhammad5943/go-fiber-gorm/models/request"
	"github.com/Muhammad5943/go-fiber-gorm/utils"
	"github.com/dgrijalva/jwt-go"
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
		log.Print("Error: ", errValidate)
		return c.Status(400).JSON(&fiber.Map{
			"message": "You might have bad request",
			"error":   strings.Split(errValidate.Error(), "Key"),
		})
	}

	// chack available user
	var user entity.User
	err := database.DB.First(&user, "email=?", login.Email).Error
	if err != nil {
		log.Print("Error: ", err)
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

	// Generate JSON Web Token
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// GenerateToken
	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Print("Error: ", errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "failed generate the token",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"token": token,
	})
}
