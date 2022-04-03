package routes

import "github.com/gofiber/fiber/v2"

type IRoutes interface {
	AddRoutes(app *fiber.App)
}
