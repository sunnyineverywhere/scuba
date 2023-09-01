package main

import (
	"fmt"
	"scuba/client"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		containers := client.GetContainers()
		for _, container := range containers {
			fmt.Println("Container ID:", container["Id"])
			fmt.Println("Container Name:", container["Names"])
			fmt.Println("Container Image:", container["Image"])
		}
		return c.JSON(containers)
	})

	
	app.Listen(":8989")
}
