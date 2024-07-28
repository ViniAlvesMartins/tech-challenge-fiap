package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	"log/slog"
	"slices"
)

type PaymentStatusUpdateMessage struct {
	OrderId int                `json:"order_id"`
	Status  enum.PaymentStatus `json:"status"`
}

type PaymentStatusUpdateHandler struct {
	logger       *slog.Logger
	orderUseCase contract.OrderUseCase
}

func NewPaymentStatusUpdateHandler(o contract.OrderUseCase, l *slog.Logger) *PaymentStatusUpdateHandler {
	return &PaymentStatusUpdateHandler{logger: l, orderUseCase: o}
}

func (f *PaymentStatusUpdateHandler) Handle(ctx context.Context, b []byte) error {
	var message PaymentStatusUpdateMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	if !slices.Contains([]enum.PaymentStatus{enum.PaymentStatusCanceled, enum.PaymentStatusConfirmed}, message.Status) {
		f.logger.Info(fmt.Sprintf("Received event with status %s. Ignoring message...", message.Status))
		return nil
	}

	// get order to check if status is valid for payment approval/cancel
	order, err := f.orderUseCase.GetById(message.OrderId)
	if err != nil {
		f.logger.Error(
			fmt.Sprintf("error getting order details: [order id: %d]", message.OrderId),
			slog.Any("error", err.Error()),
		)
		return err
	}

	if order == nil {
		f.logger.Warn(fmt.Sprintf("order not found: [order id: %d]", message.OrderId))
		return nil
	}

	// cannot cancel/approve payment of FINISHED/PAID/CANCELED order
	if slices.Contains([]enum.StatusOrder{enum.OrderStatusPaid, enum.OrderStatusFinished, enum.OrderStatusCanceled}, order.OrderStatus) {
		f.logger.Warn(fmt.Sprintf("cannot change payment status [payment status: %s] of order [order status: %s]", message.Status, order.OrderStatus))
		return nil
	}

	return f.updateOrderStatus(message.OrderId, message.Status)
}

func (f *PaymentStatusUpdateHandler) updateOrderStatus(id int, s enum.PaymentStatus) error {
	status := enum.OrderStatusPaid

	if s == enum.PaymentStatusCanceled {
		status = enum.OrderStatusCanceled
	}

	return f.orderUseCase.UpdateStatusById(id, status)
}
