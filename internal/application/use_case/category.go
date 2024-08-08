package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
)

type CategoryUseCase struct {
	categoryRepository contract.CategoryRepository
}

func NewCategoryUseCase(categoryRepository contract.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryUseCase) GetById(id int) (*entity.Category, error) {
	category, err := c.categoryRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
