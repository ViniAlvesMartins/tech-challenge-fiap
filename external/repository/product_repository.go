package repository

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/entity"

	"log/slog"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewProductRepository(db *gorm.DB, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{
		db:     db,
		logger: logger,
	}
}

func (p *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	result := p.db.Create(&product)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return product, errors.New("create product from repository has failed")
	}

	return product, nil
}

func (p *ProductRepository) Update(product entity.Product) (entity.Product, error) {
	fmt.Println(product)
	result := p.db.Save(&product)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return product, errors.New("update product from repository has failed")
	}

	return product, nil
}

func (p *ProductRepository) GetById(id int) (*entity.Product, error) {
	var product entity.Product

	result := p.db.Model(&product).Where("id = ?", id).First(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		p.logger.Error("result.Error")
		return nil, errors.New("get product by id from repository has failed")
	}

	return &product, nil
}

func (p *ProductRepository) Delete(id int) error {
	var product entity.Product

	result := p.db.Model(&product).Where("id = ?", id).Update("active", false)

	if result.Error != nil {
		p.logger.Error("result.Error")
		return errors.New("delete product from repository has failed")
	}

	return nil
}

func (p *ProductRepository) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	var product []entity.Product

	if result := p.db.Where("category_id=? AND active=true", categoryId).Find(&product); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		p.logger.Error("result.Error")
		return nil, errors.New("an error occurred from repository")
	}

	return product, nil
}
