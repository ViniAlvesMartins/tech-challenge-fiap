package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderFromEntity(t *testing.T) {
	t.Run("convert entity to dto successfully", func(t *testing.T) {
		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			OrderStatus: enum.OrderStatusAwaitingPayment,
			Amount:      123.45,
			Products: []*entity.Product{
				{
					ID:          1,
					ProductName: "Test Product 1",
					Description: "Test Product 1",
					Price:       10.5,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					ProductName: "Test Product 2",
					Description: "Test Product 2",
					Price:       10.5,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		dto := OrderFromEntity(order)

		assert.IsType(t, OrderDto{}, dto)
		assert.Equal(t, 2, len(dto.Products))
	})
}

func TestOrderListFromEntity(t *testing.T) {
	t.Run("convert entity order list to dto list successfully", func(t *testing.T) {
		orders := []entity.Order{
			{
				ID:          1,
				ClientId:    nil,
				OrderStatus: enum.OrderStatusAwaitingPayment,
				Amount:      123.45,
				Products: []*entity.Product{
					{
						ID:          1,
						ProductName: "Test Product 1",
						Description: "Test Product 1",
						Price:       10.5,
						CategoryId:  1,
						Active:      true,
					},
					{
						ID:          2,
						ProductName: "Test Product 2",
						Description: "Test Product 2",
						Price:       10.5,
						CategoryId:  1,
						Active:      true,
					},
				},
			},
		}

		dto := OrderListFromEntity(orders)
		assert.Equal(t, 1, len(dto))
	})
}
