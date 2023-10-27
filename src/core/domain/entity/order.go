package entity

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"time"
)

type Order struct {
	ID          int              `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientId    int              `json:"client_id"`
	StatusOrder enum.StatusOrder `json:"status_order"`
	Amount      float64          `json:"amount"`
	CreatedAt   time.Time        `json:"created_at,omitempty"`
	Products    []*Product       `json:"products" gorm:"many2many:orders_products"`
}
