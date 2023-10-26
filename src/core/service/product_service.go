package service

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type ProductService struct {
	productRepository port.ProductRepository
	logger            *slog.Logger
}

func NewProductService(productRepository port.ProductRepository, logger *slog.Logger) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		logger:            logger,
	}
}

func (srv *ProductService) Create(product domain.Product) (domain.Product, error) {

	prod, err := srv.productRepository.Create(product)

	if err != nil {
		return domain.Product{}, errors.New("create product from repository has failed")
	}

	return prod, nil
}

func (srv *ProductService) GetProductByCategory(categoryId int) ([]domain.Product, error) {
	fmt.Println("Cheguei no service!")
	prod, err := srv.productRepository.GetProductByCategory(categoryId)


	if err != nil {
		return nil, err
	}

	return prod, nil

}