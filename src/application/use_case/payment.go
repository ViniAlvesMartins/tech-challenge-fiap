package use_case

import (
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	response_payment_service "github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/modules/response/payment_service"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type PaymentUseCase struct {
	repository             contract.PaymentRepository
	externalPaymentService contract.ExternalPaymentService
	logger                 *slog.Logger
}

func NewPaymentUseCase(r contract.PaymentRepository, e contract.ExternalPaymentService, logger *slog.Logger) *PaymentUseCase {
	return &PaymentUseCase{
		repository:             r,
		externalPaymentService: e,
		logger:                 logger,
	}
}

func (p *PaymentUseCase) Create(payment *entity.Payment) error {
	p.repository.Create(*payment)
	return nil
}

func (p *PaymentUseCase) GetLastPaymentStatus(orderId int) (enum.PaymentStatus, error) {

	payment, err := p.repository.GetLastPaymentStatus(orderId)

	if err != nil {
		return payment.Status, err
	}

	p.logger.Info("teste", payment.Status)

	if payment.Status == "" {
		return enum.PENDING, nil
	}

	return payment.Status, nil
}

func (p *PaymentUseCase) CreateQRCode(order *entity.Order) (*response_payment_service.CreateQRCode, error) {
	payment := &entity.Payment{
		Order:  order,
		Type:   enum.QRCODE,
		Status: enum.PENDING,
		Amount: order.Amount,
	}

	p.Create(payment)

	p.logger.Info("tests", payment)

	qrCode, _ := p.externalPaymentService.CreateQRCode(*payment)

	return &qrCode, nil
}

func (p *PaymentUseCase) PaymentNotification() error {
	return nil
}

/*
	if err = p.orderUseCase.SetStatusToReceived(order.ID, enum.RECEIVED); err != nil {
		return err
*/
