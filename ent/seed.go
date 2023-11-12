package ent

import (
	"context"
)

// Seed adds sample data to the database.
func Seed(ctx context.Context, c *Client) error {
	_, err := c.User.CreateBulk(
		c.User.Create().SetName("Alice"),
		c.User.Create().SetName("Bob"),
		c.User.Create().SetName("Zod"),
	).Save(ctx)
	return err
}
