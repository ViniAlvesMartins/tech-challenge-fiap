package enum

import (
	"slices"
)

type StatusOrder string

const (
	OrderStatusAwaitingPayment StatusOrder = "AWAITING_PAYMENT"
	OrderStatusPaid            StatusOrder = "PAID"
	OrderStatusCanceled        StatusOrder = "CANCELED"
	OrderStatusReceived        StatusOrder = "RECEIVED"
	OrderStatusPreparing       StatusOrder = "PREPARING"
	OrderStatusReady           StatusOrder = "READY"
	OrderStatusFinished        StatusOrder = "FINISHED"
)

func ValidateStatus(val string) bool {
	validStatus := []StatusOrder{OrderStatusAwaitingPayment, OrderStatusPaid, OrderStatusCanceled, OrderStatusReceived, OrderStatusPreparing, OrderStatusReady, OrderStatusFinished}
	return slices.Contains(validStatus, StatusOrder(val))
}
