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
	payment.Status = enum.CONFIRMED

	return p.repository.Create(payment)
}
