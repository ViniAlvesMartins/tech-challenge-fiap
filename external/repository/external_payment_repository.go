package repository

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type ExternalPaymentRepository struct {
}

func NewExternalPaymentRepository() *ExternalPaymentRepository {
	return &ExternalPaymentRepository{}
}

func (e *ExternalPaymentRepository) Create(p entity.ExternalPayment) error {
	return nil
}
