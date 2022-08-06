package route

import (
	"github.com/Muhammad5943/go-fiber-gorm/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/test", controllers.Test)
}
