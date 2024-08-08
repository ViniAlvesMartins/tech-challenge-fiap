package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
)

type ProductUseCase struct {
	productRepository contract.ProductRepository
}

func NewProductUseCase(productRepository contract.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}

func (p *ProductUseCase) Create(product entity.Product) (*entity.Product, error) {
	product.Active = true
	productNew, err := p.productRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &productNew, nil
}

func (p *ProductUseCase) Update(product entity.Product, id int) (*entity.Product, error) {
	product.Active = true
	product.ID = id

	productUpdated, err := p.productRepository.Update(product)
	if err != nil {
		return nil, err
	}

	return productUpdated, nil
}

func (p *ProductUseCase) Delete(id int) error {
	return p.productRepository.Delete(id)
}

func (p *ProductUseCase) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	prod, err := p.productRepository.GetProductByCategory(categoryId)
	if err != nil {
		return nil, err
	}

	return prod, nil
}

func (p *ProductUseCase) GetById(id int) (*entity.Product, error) {
	prod, err := p.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
