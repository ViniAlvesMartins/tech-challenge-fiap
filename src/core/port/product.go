package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"

type ProductRepository interface {
	Create(product domain.Product) (domain.Product, error)
	GetProductByCategory(categoryId int) (domain.Product, error)
}

type ProductService interface {
	Create(product domain.Product) (domain.Product, error)
	GetProductByCategory(categoryId int) (domain.Product, error)
}
