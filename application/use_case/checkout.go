package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"log/slog"
)

type CheckoutUseCase struct {
	logger                 *slog.Logger
	paymentService         contract.PaymentUseCase
	orderService           contract.OrderUseCase
	externalPaymentService contract.ExternalPaymentService
}

func NewCheckoutUseCase(l *slog.Logger, p contract.PaymentUseCase, o contract.OrderUseCase, e contract.ExternalPaymentService) *CheckoutUseCase {
	return &CheckoutUseCase{
		logger:                 l,
		paymentService:         p,
		orderService:           o,
		externalPaymentService: e,
	}
}

func (c *CheckoutUseCase) PayWithQRCode(id int) error {
	var err error

	order, err := c.orderService.GetById(id)

	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	if err = c.paymentService.PayWithQRCode(order); err != nil {
		return err
	}

	if err = c.externalPaymentService.PayOrder(*order, enum.PIX); err != nil {
		return err
	}

	if err = c.orderService.SetStatusToReceived(order.ID, enum.RECEIVED); err != nil {
		return err
	}

	return nil
}
