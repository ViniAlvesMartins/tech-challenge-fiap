package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"gorm.io/gorm/clause"

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

func (p *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	if result := p.db.Create(&product); result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

func (p *ProductRepository) Update(product entity.Product) (*entity.Product, error) {
	var uptProduct entity.Product
	result := p.db.Model(&uptProduct).Where("id = ?", product.ID).Clauses(clause.Returning{}).Updates(entity.Product{
		NameProduct: product.NameProduct,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Active:      product.Active,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &uptProduct, nil
}

func (p *ProductRepository) GetById(id int) (*entity.Product, error) {
	var product entity.Product
	result := p.db.Model(&product).Where("id = ?", id).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &product, nil
}

func (p *ProductRepository) Delete(id int) error {
	result := p.db.Model(&entity.Product{}).Where("id = ?", id).Update("active", false)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepository) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	var product []entity.Product

	if result := p.db.Where("category_id=? AND active=true", categoryId).Find(&product); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return product, nil
}
