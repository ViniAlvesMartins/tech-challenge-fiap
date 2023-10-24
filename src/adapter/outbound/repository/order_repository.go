package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"gorm.io/gorm"
	"log/slog"
)

type OrderRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewOrderRepository(db *gorm.DB, logger *slog.Logger) *OrderRepository {
	return &OrderRepository{
		db:     db,
		logger: logger,
	}
}

func (o *OrderRepository) Create(order domain.Order) (domain.Order, error) {

	if result := o.db.Create(&order); result.Error != nil {
		o.logger.Error("result.Error")
		return order, errors.New("create order from repository has failed")
	}

	return order, nil
}
