package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {
	if result := o.db.Create(&order).Model(&order).Preload("Products").Where("id= ?", order.ID).First(&order); result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

func (o *OrderRepository) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	results := o.db.Raw("select * from ze_burguer.orders where not status_order = 'FINISHED' order by case when status_order = 'READY' then 1 when status_order = 'PREPARING' then 2 when status_order = 'RECEIVED' then 3 else 4 end asc, created_at asc").Find(&orders)
	if results.Error != nil {
		return orders, results.Error
	}

	return orders, nil
}

func (o *OrderRepository) GetById(id int) (*entity.Order, error) {
	var order entity.Order

	result := o.db.Model(&order).Preload("Products").Where("id= ?", id).First(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &order, nil
}

func (o *OrderRepository) UpdateStatusById(id int, status enum.StatusOrder) error {
	return o.db.Model(&entity.Order{}).Where("id = ?", id).Update("status_order", status).Error
}
