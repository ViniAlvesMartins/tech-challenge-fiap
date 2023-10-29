package port

type CheckoutService interface {
	PayWithQRCode(id int) error
}
