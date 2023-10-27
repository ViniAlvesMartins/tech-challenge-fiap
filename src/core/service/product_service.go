package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
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

func (p *ProductService) Create(product entity.Product) (entity.Product, error) {
	product.Active = true
	productNew, err := p.productRepository.Create(product)

	if err != nil {
		return productNew, err
	}

	return productNew, nil
}

func (p *ProductService) Update(product entity.Product) (entity.Product, error) {
	product.Active = true
	productUpdated, err := p.productRepository.Update(product)

	if err != nil {
		return productUpdated, err
	}

	return productUpdated, nil
}

func (p *ProductService) Delete(id int) error {
	err := p.productRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductService) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	prod, err := p.productRepository.GetProductByCategory(categoryId)

	if err != nil {
		return nil, err
	}

	return prod, nil
}
