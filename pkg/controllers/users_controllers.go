package controllers

import "github.com/gofiber/fiber/v2"

func GetUserProfile(c *fiber.Ctx) error {
	return c.SendString("User")

}
func LoginUser(c *fiber.Ctx) error {
	return c.SendString("User")

}
func CreateUser(c *fiber.Ctx) error {
	return c.SendString("User")
}
