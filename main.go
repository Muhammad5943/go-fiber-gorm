package main

import (
	"log"

	"github.com/Muhammad5943/go-fiber-gorm/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//Loading Environmental Variable
	loadenv()

	app := fiber.New()

	// Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
