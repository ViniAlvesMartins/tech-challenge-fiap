package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type ExternalPaymentService interface {
	PayOrder(order entity.Order, paymentType enum.PaymentType) error
	Create(payment entity.ExternalPayment) error
}

type ExternalPaymentRepository interface {
	Create(payment entity.ExternalPayment) error
}
