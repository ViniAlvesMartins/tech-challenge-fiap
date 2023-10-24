package port

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
)

type OrderRepository interface {
	Create(order domain.Order) (domain.Order, error)
}

type OrderService interface {
	Create(order domain.Order) (*domain.Order, error)
}
