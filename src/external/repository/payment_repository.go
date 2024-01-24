package repository

import (
	"errors"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewPaymentRepository(db *gorm.DB, log *slog.Logger) *PaymentRepository {
	return &PaymentRepository{
		db:  db,
		log: log,
	}
}

func (p *PaymentRepository) Create(payment *entity.Payment) error {
	return p.db.Create(payment).Error
}

func (p *PaymentRepository) GetLastPaymentStatus(orderId int) (*entity.Payment, error) {
	var payment entity.Payment

	result := p.db.Order("payments.created_at asc").Where("payments.order_id= ?", orderId).Find(payment)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &payment, nil
}
