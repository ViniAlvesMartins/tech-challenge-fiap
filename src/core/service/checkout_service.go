package service

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type CheckoutService struct {
	paymentService         port.PaymentService
	orderService           port.OrderService
	externalPaymentService port.ExternalPaymentService
}

func NewCheckoutService(p port.PaymentService, o port.OrderService, e port.ExternalPaymentService) *CheckoutService {
	return &CheckoutService{
		paymentService:         p,
		orderService:           o,
		externalPaymentService: e,
	}
}

func (c *CheckoutService) PayWithQRCode(id int) error {
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
