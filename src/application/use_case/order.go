package use_case

import (
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type OrderUseCase struct {
	orderRepository contract.OrderRepository
	logger          *slog.Logger
}

func NewOrderUseCase(orderRepository contract.OrderRepository, logger *slog.Logger) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
		logger:          logger,
	}
}

func (o *OrderUseCase) Create(order entity.Order, products []*entity.Product) (*entity.Order, error) {
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

func (o *OrderUseCase) GetAll() (*[]entity.Order, error) {
	orders, err := o.orderRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (o *OrderUseCase) GetById(id int) (*entity.Order, error) {
	order, err := o.orderRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderUseCase) SetStatusToReceived(id int, status enum.StatusOrder) error {
	return o.orderRepository.SetStatusToReceived(id, status)
}
