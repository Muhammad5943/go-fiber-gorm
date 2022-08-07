package main

import (
	"github.com/Muhammad5943/go-fiber-gorm/database"
	"github.com/Muhammad5943/go-fiber-gorm/database/migration"
	"github.com/Muhammad5943/go-fiber-gorm/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	database.DatabaseInit()
	// Run migration
	migration.RunMigration()

	// Initial fiberApp
	app := fiber.New()

	// Initial Route
	route.RouteInit(app)

	// Running App
	app.Listen(":3000")
}
