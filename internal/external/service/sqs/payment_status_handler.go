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

	if !slices.Contains([]enum.PaymentStatus{enum.PaymentStatusCanceled}, message.Status) {
		f.logger.Info(fmt.Sprintf("Received event with status %s. Ignoring message...", message.Status))
		return nil
	}

	return f.orderUseCase.UpdateStatusById(message.OrderId, enum.OrderStatusCanceled)
}
