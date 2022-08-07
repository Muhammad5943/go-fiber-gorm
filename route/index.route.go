package route

import (
	"github.com/Muhammad5943/go-fiber-gorm/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	/* Test */
	r.Get("/test", controllers.Test)

	api := r.Group("/api") // /api
	v1 := api.Group("/v1") // /api/v1

	/* User */
	v1.Get("/users", controllers.GetUsers)
	v1.Post("/users", controllers.CreateUser)
	v1.Get("/user/user_id=:userId", controllers.GetUserById)
	v1.Put("/user/user_id=:userId", controllers.UpdateUser)
	v1.Delete("/user/user_id=:userId", controllers.DeleteUser)
}
