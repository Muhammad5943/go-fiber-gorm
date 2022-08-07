package middleware

import (
	"log"

	"github.com/Muhammad5943/go-fiber-gorm/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	// log.Print("token ", token)
	if token == "" {
		log.Print("unauthenticated")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	_, err := utils.VerifyToken(token)
	log.Print(utils.VerifyToken(token))
	if err != nil {
		log.Print("Error: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated, token not match",
		})
	}

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}
