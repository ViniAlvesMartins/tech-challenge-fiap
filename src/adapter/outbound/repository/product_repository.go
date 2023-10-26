package repository

import (
	"fmt"
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

func (repo *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	fmt.Println("post")
	result := repo.db.Create(&product)

	if result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return product, nil
}

func (repo *ProductRepository) Update(product domain.Product) (domain.Product, error) {
	fmt.Println("patch")
	fmt.Println(&product)
	result := repo.db.Save(&product)

	if result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return product, nil
}
