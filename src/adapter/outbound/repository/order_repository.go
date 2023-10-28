package repository

import (
	"errors"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"gorm.io/gorm"
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

func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {
	if result := o.db.Create(&order); result.Error != nil {
		o.logger.Error("result.Error")
		return order, errors.New("create order from repository has failed")
	}

	return order, nil
}

func (o *OrderRepository) GetAll() ([]entity.Order, error) {
	var orders []entity.Order

	results := o.db.Order("orders.created_at asc").Preload("Products").Not("orders.status_order= ?", enum.FINISHED).Find(&orders)

	if results.Error != nil {
		o.logger.Error("result.Error")
		return orders, errors.New("find orders from repository has failed")
	}

	return orders, nil
}

func (o *OrderRepository) GetById(id int) (*entity.Order, error) {
	var order entity.Order

	result := o.db.Model(&order).Where("id= ?", id).Find(&order)

	if result.Error != nil {
		o.logger.Error("get order by id (%s) from repository has failed", id)
		return nil, errors.New("get order by id from repository has failed")
	}

	return &order, nil
}
