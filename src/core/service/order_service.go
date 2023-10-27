package service

import (
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"

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

	order.StatusOrder = enum.AWAITING_PAYMENT
	orderNew, err := o.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	return &orderNew, nil
}

func (o *OrderService) GetAll() (*[]entity.Order, error) {
	orders, err := o.orderRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return &orders, nil
}
