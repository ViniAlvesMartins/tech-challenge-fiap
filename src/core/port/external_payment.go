package port

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
)

type ExternalPaymentService interface {
	PayOrder(order entity.Order, paymentType enum.PaymentType) error
	Create(payment entity.ExternalPayment) error
}

type ExternalPaymentRepository interface {
	Create(payment entity.ExternalPayment) error
}
