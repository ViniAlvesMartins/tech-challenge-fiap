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

func (o *OrderRepository) AnonymizeOrderClient(clientID int) error {
	return o.db.Model(&entity.Order{}).Where("client_id = ?", clientID).Update("client_id", nil).Error
}

func (o *OrderRepository) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	results := o.db.Raw("select * from ze_burguer.orders where not order_status = 'FINISHED' order by case when order_status = 'READY' then 1 when order_status = 'PREPARING' then 2 when order_status = 'RECEIVED' then 3 else 4 end asc, created_at asc").Find(&orders)
	if results.Error != nil {
		return orders, results.Error
	}

	return orders, nil
}

func (o *OrderRepository) GetByStatus(status enum.StatusOrder) ([]*entity.Order, error) {
	var orders []*entity.Order

	result := o.db.Model(orders).Where("order_status = ?", status)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return orders, nil
}

func (o *OrderRepository) CancelExpiredOrders(threshold int) error {
	results := o.db.Exec("update ze_burguer.orders set order_status = ? "+
		"WHERE (DATE_PART('Day', timezone('utc', now()) - orders.created_at)) * 24 + (DATE_PART('Hour',  timezone('utc', now())) - DATE_PART('Hour', orders.created_at)) >= ? "+
		"AND order_status = ?", enum.OrderStatusFinished, threshold, enum.OrderStatusAwaitingPayment)
	if results.Error != nil {
		return results.Error
	}

	return nil
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
	return o.db.Model(&entity.Order{}).Where("id = ?", id).Update("order_status", status).Error
}
