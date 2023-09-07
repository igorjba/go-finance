package main

import (
	"context"
	"os"

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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	db := client.Database(os.Getenv("MONGO_DB_NAME"))

	itemRepo := item.NewMongoRepository(db, "nome_da_colecao")
	itemService := service.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	app := api.InitServer()

	routes.SetItemRoutes(app, itemHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	api.RunServer(app, ":"+port)
}
