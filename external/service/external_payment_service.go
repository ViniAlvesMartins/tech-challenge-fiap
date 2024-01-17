package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
)

type ExternalPayment struct {
	repository contract.ExternalPaymentRepository
}

func NewExternalPayment(r contract.ExternalPaymentRepository) *ExternalPayment {
	return &ExternalPayment{
		repository: r,
	}
}

func (e *ExternalPayment) Create(p entity.ExternalPayment) error {
	return e.repository.Create(p)
}

func (e *ExternalPayment) PayOrder(order entity.Order, paymentType enum.PaymentType) error {
	payment := entity.ExternalPayment{
		Order:  &order,
		Type:   paymentType,
		Amount: order.Amount,
	}

	return e.Create(payment)
}
