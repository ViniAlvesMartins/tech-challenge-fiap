//go:generate mockgen -destination=mock/repository.go -source=repository.go -package=mock
package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
)

type CategoryRepository interface {
	GetById(id int) (*entity.Category, error)
}

type ClientRepository interface {
	Create(client entity.Client) (*entity.Client, error)
	GetById(id *int) (*entity.Client, error)
	GetByCpf(cpf int) (*entity.Client, error)
	GetByCpfOrEmail(cpf int, email string) (*entity.Client, error)
}

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	GetAll() ([]entity.Order, error)
	GetById(id int) (*entity.Order, error)
	UpdateStatusById(id int, status enum.StatusOrder) error
}

type PaymentRepository interface {
	Create(payment entity.Payment) (entity.Payment, error)
	GetLastPaymentStatus(orderId int) (*entity.Payment, error)
}

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (*entity.Product, error)
	Delete(id int) error
	GetProductByCategory(categoryId int) ([]entity.Product, error)
	GetById(id int) (*entity.Product, error)
}