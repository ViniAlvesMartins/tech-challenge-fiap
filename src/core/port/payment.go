package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type PaymentRepository interface {
	Create(payment *entity.Payment) error
}

type PaymentService interface {
	Create(payment *entity.Payment) error
}

type PaymentExternal interface {
	PayWithQrCode()
}
