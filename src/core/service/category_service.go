package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
)

type CategoryService struct {
	categoryRepository port.CategoryRepository
	logger             *slog.Logger
}

func NewCategoryService(categoryRepository port.CategoryRepository, logger *slog.Logger) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
		logger:             logger,
	}
}

func (c *CategoryService) GetById(id int) (*entity.Category, error) {

	category, err := c.categoryRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}
