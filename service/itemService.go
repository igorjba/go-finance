package service

import (
	"context"
	"errors"

	"github.com/igorjba/go-test-back/repository/item"
)

type ItemService struct {
	Repo item.Repository
}

func NewItemService(repo item.Repository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) Create(ctx context.Context, item *item.Item) error {
	_, err := s.Repo.Create(ctx, item)
	return err
}

func (s *ItemService) GetByID(ctx context.Context, id string) (interface{}, error) {
	return s.Repo.Read(ctx, id)
}

func (s *ItemService) GetAll(ctx context.Context) ([]interface{}, error) {
	items, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}
	interfaceSlice := make([]interface{}, len(items))
	for i, item := range items {
		interfaceSlice[i] = item
	}
	return interfaceSlice, nil
}

func (s *ItemService) Update(ctx context.Context, id string, item *item.Item) error {
	// Verificar se o item existe antes de atualizar
	if _, err := s.Repo.Read(ctx, id); err != nil {
		return errors.New("Item não encontrado")
	}
	_, err := s.Repo.Update(ctx, item)
	return err
}

func (s *ItemService) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
