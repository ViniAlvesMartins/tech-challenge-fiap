package service

import (
	"errors"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type productService struct {
	productRepository port.ProductRepository
	logger            *slog.Logger
}

func NewProductService(productRepository port.ProductRepository, logger *slog.Logger) *productService {
	return &productService{
		productRepository: productRepository,
		logger:            logger,
	}
}

func (srv *productService) Create(product domain.Product) (domain.Product, error) {

	prod, err := srv.productRepository.Create(product)

	if err != nil {
		return domain.Product{}, errors.New("create product from repository has failed")
	}

	return prod, nil
}

func (srv *productService) GetProductByCategory(categoryId int) (*domain.Product, error) {
	prod, err := srv.productRepository.GetProductByCategory(categoryId)

	if err != nil {
		return nil, err
	}

	return client, nil

}