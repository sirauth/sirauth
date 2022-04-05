package routes

import (
	"sirauth/ent"

	"github.com/gofiber/fiber/v2"
)

type ClientRoutes struct {
	db *ent.Client
}

func NewClientRoutes(db *ent.Client) IRoutes {
	return &ClientRoutes{
		db,
	}
}

func (c *ClientRoutes) AddRoutes(app *fiber.App) {
	app.Get("/clients/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		client, err := c.db.RealmOAuthClient.Get(ctx.Context(), id)
		if err != nil {
			return err
		}

		return ctx.JSON(client)
	})
}
