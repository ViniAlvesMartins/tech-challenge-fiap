package enum

type PaymentType string

type PaymentStatus string

const (
	PaymentTypeCredit PaymentType = "CREDIT"
	PaymentTypeDebit  PaymentType = "DEBIT"
	PaymentTypeCash   PaymentType = "CASH"
	PaymentTypePix    PaymentType = "PIX"
	PaymentTypeQRCode PaymentType = "QRCODE"

	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusConfirmed PaymentStatus = "CONFIRMED"
	PaymentStatusCanceled  PaymentStatus = "CANCELED"
)
