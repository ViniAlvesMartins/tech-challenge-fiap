//go:generate mockgen -destination=mock/service.go -source=service.go -package=mock
package contract

import (
	response_payment_service "github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/modules/response/payment_service"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
)

type ExternalPaymentService interface {
	CreateQRCode(payment entity.Payment) (response_payment_service.CreateQRCode, error)
}
