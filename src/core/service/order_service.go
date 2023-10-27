package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"log/slog"
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type OrderService struct {
	orderRepository port.OrderRepository
	logger          *slog.Logger
}

func NewOrderService(orderRepository port.OrderRepository, logger *slog.Logger) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		logger:          logger,
	}
}

func (o *OrderService) Create(order entity.Order) (*entity.Order, error) {
	order.CreatedAt = time.Now()
	order.StatusOrder = "WAITING"
	orderNew, err := o.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	return &orderNew, nil
}

func (o *OrderService) Find() (*[]entity.Order, error) {
	orders, err := o.orderRepository.Find()

	if err != nil {
		return nil, err
	}

	return &orders, nil
}
