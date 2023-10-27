package enum

type StatusOrder string

const (
	AWAITING_PAYMENT StatusOrder = "AWAITING_PAYMENT"
	RECEIVED         StatusOrder = "RECEIVED"
	PREPARING        StatusOrder = "PREPARING"
	READY            StatusOrder = "READY"
	FINISHED         StatusOrder = "FINISHED"
)
