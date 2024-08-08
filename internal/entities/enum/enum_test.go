package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateStatus(t *testing.T) {
	t.Run("validate status successfully", func(t *testing.T) {
		assert.True(t, ValidateOrderStatus(string(OrderStatusAwaitingPayment)))
	})

	t.Run("invalid order status", func(t *testing.T) {
		assert.False(t, ValidateOrderStatus("INVALID_STATUS"))
	})
}
