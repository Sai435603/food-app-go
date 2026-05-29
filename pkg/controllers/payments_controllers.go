package controllers

import "github.com/gofiber/fiber/v2"

func PaymentWebhook(c *fiber.Ctx) error {
	return c.SendString("payments")
}

func InitializePayment(c *fiber.Ctx) error {
	return c.SendString("payments")
}
