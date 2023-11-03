package server

import (
	"github.com/DavidNix/indie/ent"
	"github.com/gofiber/fiber/v2"
)

func userListHandler(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		users, err := client.User.Query().All(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve users",
			})
		}

		// TODO: Don't return the raw entities. Only done here for convenience.
		return c.JSON(users)
	}
}
