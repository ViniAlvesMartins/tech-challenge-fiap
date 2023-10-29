package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type PaymentService struct {
	repository port.PaymentRepository
}

func NewPaymentService(r port.PaymentRepository) *PaymentService {
	return &PaymentService{
		repository: r,
	}
}

func (p *PaymentService) Create(payment *entity.Payment) error {
	return p.repository.Create(payment)
}

func (p *PaymentService) PayWithQRCode(order *entity.Order) error {
	payment := &entity.Payment{
		Order:  order,
		Type:   enum.PIX,
		Status: enum.CONFIRMED,
		Amount: order.Amount,
	}

	return p.Create(payment)
}
