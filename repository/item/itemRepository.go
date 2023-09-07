package item

import "context"

type Item struct {
	ID      string `bson:"_id,omitempty"`
	Name    string `bson:"name,omitempty"`
	Details string `bson:"details,omitempty"`
}

type Repository interface {
	Create(ctx context.Context, item *Item) (*Item, error)
	Read(ctx context.Context, id string) (*Item, error)
	Update(ctx context.Context, item *Item) (*Item, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Item, error)
}
