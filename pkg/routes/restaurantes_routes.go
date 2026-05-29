package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/food-app-go/pkg/controllers"
)

func SetUpRestaurentRoutes(router fiber.Router) {
	router.Get("/", controllers.GetAllRestaurants)
	router.Get("/:id", controllers.GetRestaurantByID)
	router.Get("/:id/menu", controllers.GetRestaurantMenu)
}
