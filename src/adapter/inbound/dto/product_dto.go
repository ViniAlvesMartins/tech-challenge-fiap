package dto

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
)

type ProductDto struct {
	ID          int     `json:"id"`
	NameProduct string  `json:"name_product"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required,gt=0" error:"O Número deve ser maior que zero"`
	CategoryId  int     `json:"category_id" validate:"required"`
	Active      bool    `json:"active"`
}

func (p *ProductDto) ConvertToEntity() entity.Product {
	return entity.Product{
		ID:          p.ID,
		NameProduct: p.NameProduct,
		Description: p.Description,
		Price:       p.Price,
		CategoryId:  p.CategoryId,
		Active:      p.Active,
	}
}
