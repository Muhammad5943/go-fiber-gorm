package route

import (
	"github.com/Muhammad5943/go-fiber-gorm/config"
	"github.com/Muhammad5943/go-fiber-gorm/controllers"
	"github.com/Muhammad5943/go-fiber-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	/* Static/ Public Folder */
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	/* Test */
	r.Get("/test", controllers.Test)

	api := r.Group("/api") // /api
	v1 := api.Group("/v1") // /api/v1

	/* User */
	v1.Get("/users", middleware.AuthMiddleware, controllers.GetUsers)
	v1.Post("/users", controllers.CreateUser)
	v1.Get("/user/user_id=:userId", controllers.GetUserById)
	v1.Put("/user/user_id=:userId", controllers.UpdateUser)
	v1.Delete("/user/user_id=:userId", controllers.DeleteUser)
}
