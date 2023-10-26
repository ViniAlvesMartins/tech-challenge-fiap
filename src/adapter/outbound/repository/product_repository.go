package repository

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"log/slog"

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

func (repo *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	result := repo.db.Create(&product)

	if result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return product, nil
}

func (repo *ProductRepository) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	fmt.Println("Cheguei no Repositorio!")
	var product []entity.Product
	fmt.Println(product)

	if result := repo.db.Debug().Where("category_id=?", categoryId).Find(&product); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		repo.logger.Error("result.Error")
		return nil, errors.New("an error occurred from repository")
	}

	fmt.Println(product)
	fmt.Println(categoryId)

	return product, nil
}
