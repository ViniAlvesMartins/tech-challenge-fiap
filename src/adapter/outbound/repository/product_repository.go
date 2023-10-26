package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"gorm.io/gorm"
	"log/slog"
)

type ProductRepository struct {
	db *gorm.DB

	logger *slog.Logger
}

func NewProductRepository(db *gorm.DB, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{
		db:     db,
		logger: logger,
	}
}

func (p ProductRepository) Create(product domain.Product) (domain.Product, error) {
	result := p.db.Create(&product)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return product, errors.New("create product from repository has failed")
	}

	return product, nil
}

func (p ProductRepository) Update(product domain.Product) (domain.Product, error) {
	result := p.db.Save(&product)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return product, errors.New("update product from repository has failed")
	}

	return product, nil
}

func (p ProductRepository) Delete(id int) error {
	var product domain.Product

	result := p.db.Model(&product).Where("id = ?", id).Update("active", false)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return errors.New("delete product from repository has failed")
	}

	return nil
}
