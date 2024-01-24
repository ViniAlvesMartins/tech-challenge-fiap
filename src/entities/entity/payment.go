package entity

import (
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type Payment struct {
	ID        int                `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int                `json:"-"`
	Order     *Order             `json:"order" gorm:"foreignKey:OrderID;references:ID"`
	Type      enum.PaymentType   `json:"type"`
	Status    enum.PaymentStatus `json:"status"`
	Amount    float32            `json:"amount"`
	CreatedAt *time.Time         `json:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
}
