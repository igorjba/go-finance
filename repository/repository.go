package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository defines the interface for any repository.
type Repository interface {
	Connect() (*mongo.Client, error)
}
