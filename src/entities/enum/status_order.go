package enum

import "errors"

type StatusOrder string

const (
	AWAITING_PAYMENT StatusOrder = "AWAITING_PAYMENT"
	RECEIVED         StatusOrder = "RECEIVED"
	PREPARING        StatusOrder = "PREPARING"
	READY            StatusOrder = "READY"
	FINISHED         StatusOrder = "FINISHED"
)

func ValidateStatus(val string) (bool, error) {
	if v := StatusOrder(val); v == AWAITING_PAYMENT || v == RECEIVED || v == PREPARING || v == READY || v == FINISHED {
		return true, nil
	}
	return false, errors.New("invalid StatusType")
}
