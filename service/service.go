package service

type Service interface {
	Create(item interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	Update(id string, item interface{}) error
	Delete(id string) error
}
