package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/food-app-go/pkg/controllers"
)

func SetUpPaymentRoutes(router fiber.Router) {
	router.Post("/initialize", controllers.InitializePayment)
	router.Post("/webhook", controllers.PaymentWebhook)
}
