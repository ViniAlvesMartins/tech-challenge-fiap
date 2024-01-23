package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type OrderUseCase interface {
	Create(order entity.Order, products []*entity.Product) (*entity.Order, error)
	GetAll() (*[]entity.Order, error)
	GetById(id int) (*entity.Order, error)
	SetStatusToReceived(id int, status enum.StatusOrder) error
}

type CheckoutUseCase interface {
	PayWithQRCode(id int) error
}

type PaymentUseCase interface {
	Create(payment *entity.Payment) error
	Checkout(id int) error
	PayWithQRCode(order *entity.Order) error
	GetLastPaymentStatus(orderId int) (*enum.PaymentStatus, error)
}

type CategoryUseCase interface {
	GetById(id int) (*entity.Category, error)
}

type ClientUseCase interface {
	GetClientByCpf(cpf int) (*entity.Client, error)
	Create(client entity.Client) (*entity.Client, error)
	GetAlreadyExists(cpf int, email string) (*entity.Client, error)
}

type ProductUseCase interface {
	Create(product entity.Product) (*entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(id int) error
	GetProductByCategory(categoryId int) ([]entity.Product, error)
	GetById(int int) (*entity.Product, error)
}
