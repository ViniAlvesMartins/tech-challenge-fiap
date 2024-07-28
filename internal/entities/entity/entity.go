package entity

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	"time"
)

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Client struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Cpf   int    `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CategoryId  int     `json:"category_id"`
	Active      bool    `json:"active"`
}

type Order struct {
	ID          int              `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientId    *int             `json:"client_id"`
	OrderStatus enum.StatusOrder `json:"order_status"`
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
