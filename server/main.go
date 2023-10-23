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

	app.Get("/images", func(c *fiber.Ctx) error {
		images := client.GetDockerImages()
		return c.JSON(images)
	})

	app.Static("/", "./pages")
	app.Listen(":8989")
}
