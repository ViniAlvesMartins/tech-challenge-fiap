package enum

type StatusOrder string

const (
	WAITING   StatusOrder = "WAITING"
	RECEIVED  StatusOrder = "RECEIVED"
	PREPARING StatusOrder = "PREPARING"
	READY     StatusOrder = "READY"
	FINISHED  StatusOrder = "FINISHED"
)
