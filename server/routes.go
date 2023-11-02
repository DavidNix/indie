package server

import (
	"github.com/DavidNix/indie/ent"
	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App, client *ent.Client) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// Example users index
	app.Get("/users", userListHandler(client))
}
