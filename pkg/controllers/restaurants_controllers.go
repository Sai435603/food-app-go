package controllers

import "github.com/gofiber/fiber/v2"

func GetRestaurantMenu(c *fiber.Ctx) error {
	return c.SendString("Restaurant")

}

func GetRestaurantByID(c *fiber.Ctx) error {
	return c.SendString("Restaurant")

}

func GetAllRestaurants(c *fiber.Ctx) error {
	return c.SendString("Restaurant")
}
