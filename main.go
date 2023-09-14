package main

import (
	"context"
	"fmt"
	"log"

	"github.com/igorjba/go-test-back/api"
	"github.com/igorjba/go-test-back/api/v1/handlers"
	"github.com/igorjba/go-test-back/api/v1/routes"
	"github.com/igorjba/go-test-back/config"
	"github.com/igorjba/go-test-back/repository/item"
	"github.com/igorjba/go-test-back/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.Init()

	client, err := setupMongoDB()
	if err != nil {
		log.Fatalf("Erro ao conectar com MongoDB: %s", err)
	}

	db := client.Database(config.AppConfig.MongoDatabase)

	itemRepo := item.NewMongoRepository(db, "nome_da_colecao")
	itemService := service.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	app := api.InitServer()
	routes.SetItemRoutes(app, itemHandler)

	port := config.AppConfig.ServerPort
	api.RunServer(app, fmt.Sprintf(":%s", port))
}

func setupMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURI)
	return mongo.Connect(context.TODO(), clientOptions)
}
