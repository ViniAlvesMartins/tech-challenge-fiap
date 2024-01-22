package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/enum"
)

type PaymentUseCase struct {
	repository contract.PaymentRepository
}

func NewPaymentUseCase(r contract.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		repository: r,
	}
}

func (p *PaymentUseCase) Create(payment *entity.Payment) error {
	return p.repository.Create(payment)
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
