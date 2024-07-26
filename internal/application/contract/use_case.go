//go:generate mockgen -destination=mock/use_case.go -source=use_case.go -package=mock
package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
)

type OrderUseCase interface {
	Create(order entity.Order) (*entity.Order, error)
	GetAll() (*[]entity.Order, error)
	GetById(id int) (*entity.Order, error)
	UpdateStatusById(id int, status enum.StatusOrder) error
	CancelExpiredOrders(threshold int) error
	GetByStatus(status enum.StatusOrder) ([]*entity.Order, error)
}

type CategoryUseCase interface {
	GetById(id int) (*entity.Category, error)
}

type ClientUseCase interface {
	GetByCpf(cpf int) (*entity.Client, error)
	GetById(id *int) (*entity.Client, error)
	Create(client entity.Client) (*entity.Client, error)
	GetByCpfOrEmail(cpf int, email string) (*entity.Client, error)
}

type ProductUseCase interface {
	Create(product entity.Product) (*entity.Product, error)
	Update(product entity.Product, id int) (*entity.Product, error)
	Delete(id int) error
	GetProductByCategory(categoryId int) ([]entity.Product, error)
	GetById(int int) (*entity.Product, error)
}
