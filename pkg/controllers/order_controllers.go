package controllers

import "github.com/gofiber/fiber/v2"

func UpdateOrderStatus(c *fiber.Ctx) error {
	return c.SendString("order-1")
}

func GetOrderById(c *fiber.Ctx) error {
	return c.SendString("order-1")
}

func CreateOrder(c *fiber.Ctx) error {
	return c.SendString("order-1")
}
