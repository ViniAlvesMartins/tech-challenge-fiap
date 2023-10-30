package service

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type PaymentService struct {
	repository port.PaymentRepository

	orderService           port.OrderService
	externalPaymentService port.ExternalPaymentService
}

func NewPaymentService(r port.PaymentRepository, o port.OrderService, e port.ExternalPaymentService) *PaymentService {
	return &PaymentService{
		repository:             r,
		orderService:           o,
		externalPaymentService: e,
	}
}

func (p *PaymentService) Create(payment *entity.Payment) error {
	return p.repository.Create(payment)
}

func (p *PaymentService) PayWithQRCode(id int) error {
	var err error

	order, err := p.orderService.GetById(id)

	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	payment := &entity.Payment{
		Order:  order,
		Type:   enum.PIX,
		Status: enum.CONFIRMED,
		Amount: order.Amount,
	}

	if err = p.Create(payment); err != nil {
		return err
	}

	if err = p.externalPaymentService.PayOrder(*order, enum.PIX); err != nil {
		return err
	}

	if err = p.orderService.SetStatusToReceived(order.ID, enum.RECEIVED); err != nil {
		return err
	}

	return nil
}
