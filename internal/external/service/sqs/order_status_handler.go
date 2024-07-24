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
	orderUseCase contract.OrderUseCase
	logger       *slog.Logger
}

func NewOrderStatusUpdateHandler(l *slog.Logger, o contract.OrderUseCase) *OrderStatusUpdateHandler {
	return &OrderStatusUpdateHandler{logger: l, orderUseCase: o}
}

func (f *OrderStatusUpdateHandler) Handle(ctx context.Context, b []byte) error {
	var message OrderStatusUpdateMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	if !slices.Contains([]enum.StatusOrder{enum.OrderStatusFinished, enum.OrderStatusPreparing}, message.Status) {
		f.logger.Info(fmt.Sprintf("Received event with status %s. Ignoring message...", message.Status))
		return nil
	}

	//add use case method to change order status

	return nil
}
