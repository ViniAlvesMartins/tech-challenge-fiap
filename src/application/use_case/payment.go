package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type PaymentUseCase struct {
	repository             contract.PaymentRepository
	externalPaymentService contract.ExternalPaymentService
	orderUseCase           contract.OrderUseCase
}

func NewPaymentUseCase(r contract.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		repository: r,
	}
}

func (p *PaymentUseCase) Create(payment *entity.Payment) error {
	return p.repository.Create(payment)
}

func (p *PaymentUseCase) GetLastPaymentStatus(orderId int) (*enum.PaymentStatus, error) {

	payment, err := p.repository.GetLastPaymentStatus(orderId)

	if err != nil {
		return &payment.Status, err
	}

	return &payment.Status, nil
}

func (p *PaymentUseCase) PayWithQRCode(order *entity.Order) error {
	payment := &entity.Payment{
		Order:  order,
		Type:   enum.PIX,
		Status: enum.CONFIRMED,
		Amount: order.Amount,
	}

	return p.Create(payment)
}

func (p *PaymentUseCase) Checkout(id int) error {
	var err error

	order, err := p.orderUseCase.GetById(id)

	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	if err = p.PayWithQRCode(order); err != nil {
		return err
	}

	if err = p.externalPaymentService.PayOrder(*order, enum.PIX); err != nil {
		return err
	}

	if err = p.orderUseCase.SetStatusToReceived(order.ID, enum.RECEIVED); err != nil {
		return err
	}

	return nil
}
