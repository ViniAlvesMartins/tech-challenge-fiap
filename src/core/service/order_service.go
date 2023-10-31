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

func (o *OrderService) Create(order entity.Order, products []*entity.Product) (*entity.Order, error) {
	order.StatusOrder = enum.AWAITING_PAYMENT

	var amount float32
	amount = 0

	for _, prod := range products {
		amount += prod.Price
	}

	order.SetAmount(amount)

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

func (o *OrderService) GetById(id int) (*entity.Order, error) {
	order, err := o.orderRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderService) SetStatusToReceived(id int, status enum.StatusOrder) error {
	return o.orderRepository.SetStatusToReceived(id, status)
}
