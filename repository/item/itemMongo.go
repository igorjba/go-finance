package item

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collectionName string) Repository {
	return &mongoRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *mongoRepository) Create(ctx context.Context, item *Item) (*Item, error) {
	_, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *mongoRepository) Read(ctx context.Context, id string) (*Item, error) {
	var item Item
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *mongoRepository) Update(ctx context.Context, item *Item) (*Item, error) {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": item.ID}, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *mongoRepository) List(ctx context.Context) ([]*Item, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*Item
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}
