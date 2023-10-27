package repository

import (
	"errors"
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
func (repo *ProductRepository) GetProductByCategory(categoryId int) ([]domain.Product, error) {
	fmt.Println("Cheguei no Repositorio!")
	var product []domain.Product
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
