package port

import "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"

type ProductRepository interface {
	Create(product domain.Product) (domain.Product, error)
}

type ProductService interface {
	Create(description string, price float32, category string) (domain.Product, error)
}
