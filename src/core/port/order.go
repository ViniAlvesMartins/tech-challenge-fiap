package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	GetAll() ([]entity.Order, error)
	GetById(id int) (*entity.Order, error)
}

type OrderService interface {
	Create(order entity.Order) (*entity.Order, error)
	GetAll() (*[]entity.Order, error)
	GetById(id int) (*entity.Order, error)
}
