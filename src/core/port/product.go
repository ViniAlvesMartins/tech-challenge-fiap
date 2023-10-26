package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	GetProductByCategory(categoryId int) ([]entity.Product, error)
}

type ProductService interface {
	Create(product entity.Product) (entity.Product, error)
	GetProductByCategory(categoryId int) ([]entity.Product, error)
}
