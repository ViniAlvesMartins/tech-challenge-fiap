package repository

import (
	"errors"
	"fmt"
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

func (p *PaymentRepository) Create(payment entity.Payment) (entity.Payment, error) {
	result := p.db.Create(&payment)

	if result.Error != nil {
		fmt.Println(result)
		p.log.Error("result.Error")
		return payment, errors.New("create payment from repository has failed")
	}

	return payment, nil
}

func (p *PaymentRepository) GetLastPaymentStatus(orderId int) (*entity.Payment, error) {
	var payment entity.Payment

	result := p.db.Order("payments.created_at desc").Where("payments.order_id= ?", orderId).Find(&payment)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &payment, nil
}
