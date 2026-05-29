package routes

import "github.com/gofiber/fiber/v2"

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
}
