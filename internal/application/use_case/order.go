package use_case

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type OrderUseCase struct {
	orderRepository contract.OrderRepository
	productUseCase  contract.ProductUseCase
	snsService      contract.SnsService
}

func NewOrderUseCase(o contract.OrderRepository, p contract.ProductUseCase, s contract.SnsService) *OrderUseCase {
	return &OrderUseCase{
		orderRepository: o,
		productUseCase:  p,
		snsService:      s,
	}
}

func (o *OrderUseCase) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
	for i, p := range order.Products {
		product, err := o.productUseCase.GetById(p.ID)

		if err != nil {
			return nil, ErrProductNotFound
		}

		order.Products[i].Price = product.Price
	}

	order.OrderStatus = enum.OrderStatusAwaitingPayment
	order.Amount = order.GetAmount()

	orderNew, err := o.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	if err = o.snsService.SendMessage(ctx, orderNew); err != nil {
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

func (o *OrderUseCase) GetByStatus(status enum.StatusOrder) ([]*entity.Order, error) {
	return o.orderRepository.GetByStatus(status)
}

func (o *OrderUseCase) CancelExpiredOrders(threshold int) error {
	return o.orderRepository.CancelExpiredOrders(threshold)
}

func (o *OrderUseCase) AnonymizeOrderClient(clientID int) error {
	return o.orderRepository.AnonymizeOrderClient(clientID)
}
