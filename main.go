package main

import (
	"github.com/gofiber/fiber/v2"
)

type Song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var songs = []Song{
	{ID: "1", Title: "Imagine", Artist: "John Lennon"},
	{ID: "2", Title: "Hey Jude", Artist: "The Beatles"},
	{ID: "3", Title: "Bohemian Rhapsody", Artist: "Queen"},
}

func main() {
	app := fiber.New()

	app.Get("/songs", func(c *fiber.Ctx) error {
		return c.JSON(songs)
	})

	app.Post("/songs", func(c *fiber.Ctx) error {
		newSong := new(Song)

		if err := c.BodyParser(newSong); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format salah"})
		}

		songs = append(songs, *newSong)
		return c.Status(201).JSON(newSong)
	})

	app.Delete("/songs/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, song := range songs {
			if song.ID == id {
				songs = append(songs[:i], songs[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}

		return c.Status(404).JSON(fiber.Map{"message": "Lagu ga nemu"})
	})

	app.Listen(":3000")
}
