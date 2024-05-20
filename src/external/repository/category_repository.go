package repository

import (
	"errors"
	"fmt"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"

	"log/slog"
	"strconv"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewCategoryRepository(db *gorm.DB, logger *slog.Logger) *CategoryRepository {
	return &CategoryRepository{
		db:     db,
		logger: logger,
	}
}

func (c *CategoryRepository) GetById(id int) (*entity.Category, error) {
	var category entity.Category

	result := c.db.Model(&category).Where("id= ?", id).First(&category)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error(fmt.Sprintf("get category by id (%s) from repository has failed", strconv.Itoa(id)))
		return nil, result.Error
	}

	return &category, nil
}
