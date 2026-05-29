package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/food-app-go/pkg/controllers"
)

func SetUpUserRoutes(router fiber.Router) {
	router.Post("/register", controllers.CreateUser)
	router.Post("/login", controllers.LoginUser)
	router.Get("/profile", controllers.GetUserProfile)
}
