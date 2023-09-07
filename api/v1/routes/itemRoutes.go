package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igorjba/go-test-back/api/v1/handlers"
	"github.com/igorjba/go-test-back/api/v1/middlewares" // Importando os middlewares
)

func SetItemRoutes(app *fiber.App, handler *handlers.ItemHandler) {
	itemGroup := app.Group("/items", middlewares.AuthMiddleware) // Adicionando middleware de autenticação

	itemGroup.Get("/", handler.GetAllItems)
	itemGroup.Get("/:id", handler.GetItemByID)
	itemGroup.Post("/", handler.CreateItem)
	itemGroup.Put("/:id", handler.UpdateItem)
	itemGroup.Delete("/:id", handler.DeleteItem)
}
