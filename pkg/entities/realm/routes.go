package realm

import (
	"sirauth/ent"
	"sirauth/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	db *ent.Client
}

func NewRoutes(db *ent.Client) routes.IRoutes {
	return &Route{
		db,
	}
}

func (c *Route) AddRoutes(app *fiber.App) {
	app.Get("/realms/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		if id == 0 {
			return ctx.Status(400).JSON(map[string]string{
				"error": "invalid id",
			})
		}
		entity, err := c.db.Realm.Get(ctx.Context(), id)
		if err != nil {
			return err
		}

		return ctx.JSON(entity)
	})
}
