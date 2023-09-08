package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Connect() (*mongo.Client, error)
}
