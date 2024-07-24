package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
)

type OrderUseCase struct {
	orderRepository contract.OrderRepository
}

func NewOrderUseCase(orderRepository contract.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
	}
}

func (o *OrderUseCase) Create(order entity.Order) (*entity.Order, error) {
	order.StatusOrder = enum.OrderStatusAwaitingPayment
	order.Amount = order.GetAmount()

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

func (o *OrderUseCase) UpdateStatusById(id int, status enum.StatusOrder) error {
	err := o.orderRepository.UpdateStatusById(id, status)
	if err != nil {
		return err
	}

	return nil
}
