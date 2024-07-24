package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	"log/slog"
	"slices"
)

type PaymentStatusUpdateMessage struct {
	OrderId int                `json:"order_id"`
	Status  enum.PaymentStatus `json:"status"`
}

type PaymentStatusUpdateHandler struct {
	logger *slog.Logger
}

func NewPaymentStatusUpdateHandler(l *slog.Logger) *PaymentStatusUpdateHandler {
	return &PaymentStatusUpdateHandler{logger: l}
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

	//add use case method to cancel order

	return nil
}
