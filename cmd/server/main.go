package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/food-app-go/pkg/config"
	"github.com/sai435603/food-app-go/pkg/routes"
)

func main() {
	fmt.Println("Hey, Welcome to Food App")
	app := fiber.New()
	config.GetDb()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
