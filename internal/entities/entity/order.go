package entity

import (
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
)

type Order struct {
	ID          int              `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientId    *int             `json:"client_id"`
	StatusOrder enum.StatusOrder `json:"status_order"`
	Amount      float32          `json:"amount"`
	CreatedAt   time.Time        `json:"created_at,omitempty"`
	Products    []*Product       `json:"products" gorm:"many2many:orders_products"`
}

func (o *Order) GetAmount() float32 {
	var amount float32

	for _, p := range o.Products {
		amount += p.Price
	}

	return amount
}
