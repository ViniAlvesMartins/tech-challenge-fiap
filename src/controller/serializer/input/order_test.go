package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderDto_ConvertToEntity(t *testing.T) {
	t.Run("convert dto to entity successfully", func(t *testing.T) {
		o := OrderDto{
			ClientId: nil,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 3,
				},
			},
		}

		order := o.ConvertToEntity()

		assert.Equal(t, 3, len(order.Products))
	})
}
