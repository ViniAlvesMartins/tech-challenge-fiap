package repository

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/entity"
	"gorm.io/gorm"
	"log/slog"
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
