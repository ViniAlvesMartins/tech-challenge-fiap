package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"time"
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

func (o *OrderService) Create(order domain.Order) (*domain.Order, error) {
	order.CreatedAt = time.Now()
	order.StatusOrder = "WAITING"
	orderNew, err := o.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	return &orderNew, nil
}

func (o *OrderService) Find() (*[]domain.Order, error) {
	orders, err := o.orderRepository.Find()

	if err != nil {
		return nil, err
	}

	return &orders, nil
}
