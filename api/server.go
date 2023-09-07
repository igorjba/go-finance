package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/igorjba/go-test-back/api/v1/middlewares"
)

func InitServer() *fiber.App {
	app := fiber.New()

	app.Use(
		recover.New(),
		logger.New(),
		middlewares.ErrorHandlerMiddleware,
	)

	return app
}

func RunServer(app *fiber.App, port string) {
	log.Printf("Servidor rodando na porta %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s", err.Error())
	}
}
