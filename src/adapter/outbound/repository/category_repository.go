package repository

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
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

	result := c.db.Model(&category).Where("id= ?", id).Find(&category)

	if result.Error != nil {
		c.logger.Error(fmt.Sprintf("get category by id (%s) from repository has failed", strconv.Itoa(id)))
		return nil, errors.New("get category by id from repository has failed")
	}

	return &category, nil
}
