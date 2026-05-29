package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/")

	user := api.Group("/users")
	SetUpUserRoutes(user)
	restaurent := api.Group("/restaurents")
	SetUpRestaurentRoutes(restaurent)
	orders := api.Group("/orders")
	SetUpOrderRoutes(orders)
	payments := api.Group("/payments")
	SetUpPaymentRoutes(payments)

	log.Println("All routes successfully registered!")
}
