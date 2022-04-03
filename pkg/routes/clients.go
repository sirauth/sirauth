package routes

import (
	"sirauth/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClientRoutes struct {
	db *pgxpool.Pool
}

func NewClientRoutes(db *pgxpool.Pool) IRoutes {
	return &ClientRoutes{
		db,
	}
}

func (c *ClientRoutes) AddRoutes(app *fiber.App) {
	app.Get("/clients/:clientId", func(ctx *fiber.Ctx) error {
		clientId := ctx.Params("clientId")
		client, err := entities.GetByClientID(c.db, clientId)
		if err != nil {
			return err
		}

		return ctx.JSON(client)
	})
}
