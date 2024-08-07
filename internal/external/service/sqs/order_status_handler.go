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

type OrderStatusUpdateMessage struct {
	OrderId int              `json:"order_id"`
	Status  enum.StatusOrder `json:"status"`
}

type OrderStatusUpdateHandler struct {
	logger       *slog.Logger
	orderUseCase contract.OrderUseCase
}

func NewOrderStatusUpdateHandler(o contract.OrderUseCase, l *slog.Logger) *OrderStatusUpdateHandler {
	return &OrderStatusUpdateHandler{logger: l, orderUseCase: o}
}

func (f *OrderStatusUpdateHandler) Handle(ctx context.Context, b []byte) error {
	var message OrderStatusUpdateMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	if !slices.Contains([]enum.StatusOrder{enum.OrderStatusReceived, enum.OrderStatusReady, enum.OrderStatusFinished, enum.OrderStatusPreparing}, message.Status) {
		f.logger.Info(fmt.Sprintf("Received event with status %s. Ignoring message...", message.Status))
		return nil
	}

	return f.orderUseCase.UpdateStatusById(message.OrderId, message.Status)
}
