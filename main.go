package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"open-music/config"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	app.Listen(":3000")
}
