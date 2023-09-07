package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igorjba/go-test-back/api/v1/handlers"
)

func SetRoutes(app *fiber.App, handler *handlers.ItemHandler) {
	SetItemRoutes(app, handler)
}
