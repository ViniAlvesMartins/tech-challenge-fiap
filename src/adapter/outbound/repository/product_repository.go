package repository

import (
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fmt"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) Create(product domain.Product) (domain.Product, error) {

	if result := repo.db.Save(&product); result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return product, nil
}
