package entity

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"github.com/google/uuid"
)

type ExternalPayment struct {
	ID      *uuid.UUID       `json:"id"`
	OrderID int              `json:"-"`
	Order   *Order           `json:"order"`
	Type    enum.PaymentType `json:"type"`
	Amount  float32          `json:"amount"`
}
