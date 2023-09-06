package main

import (
	"scuba/client"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()

	app.Get("/containers", func(c *fiber.Ctx) error {
		containers := client.GetContainers()
		return c.JSON(containers)
	})

	app.Static("/", "./pages")
	app.Listen(":8989")
}
