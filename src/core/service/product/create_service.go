package service

import (
	"errors"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type productService struct {
	productRepository port.ProductRepository
}

func NewProductService(productRepository port.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (srv *productService) Create(product domain.Product) (domain.Product, error) {

	prod, err := srv.productRepository.Create(product)

	if err != nil {
		return domain.Product{}, errors.New("create product from repository has failed")
	}

	return prod, nil
}
