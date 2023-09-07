package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "../go-test/api"
	"../go-test/config"
)

func main() {
	config.Init()

	app := fiber.New()

	app.Use(logger.New())

	api.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
