package repository

import (
	"fmt"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"gorm.io/gorm"
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

func (repo *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	result := repo.db.Create(&product)

	if result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return product, nil
}

func (repo *ProductRepository) GetProductByCategory(categoryId int) (*domain.Product, error) {
	var product domain.Product

	result := repo.db.Where(&categoryId)

	if result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return &product, nil
}