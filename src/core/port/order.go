package port

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	GetAll() ([]entity.Order, error)
	GetById(id int) (*entity.Order, error)
	SetStatusToReceived(id int, status enum.StatusOrder) error
}

type OrderService interface {
	Create(order entity.Order, products []*entity.Product) (*entity.Order, error)
	GetAll() (*[]entity.Order, error)
	GetById(id int) (*entity.Order, error)
	SetStatusToReceived(id int, status enum.StatusOrder) error
}
