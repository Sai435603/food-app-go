package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hey, Welcome to Food App")
	app := fiber.New()
	log.Fatal(app.Listen(":8000"))
}
