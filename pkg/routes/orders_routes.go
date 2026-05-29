package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/food-app-go/pkg/controllers"
)

func SetUpOrderRoutes(router fiber.Router) {
	router.Post("/", controllers.CreateOrder)
	router.Get("/:id", controllers.GetOrderById)
	router.Patch("/:id/status", controllers.UpdateOrderStatus)
}
