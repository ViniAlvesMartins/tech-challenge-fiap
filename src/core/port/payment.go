package port

type PaymentService interface {
	Checkout() (bool, error)
}
