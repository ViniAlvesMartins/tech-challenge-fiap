package input

import "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"

type ProductDto struct {
	ProductName string  `json:"product_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required,gt=0" error:"O Número deve ser maior que zero"`
	CategoryId  int     `json:"category_id" validate:"required"`
}

func (p *ProductDto) ConvertToEntity() entity.Product {
	return entity.Product{
		ProductName: p.ProductName,
		Description: p.Description,
		Price:       p.Price,
		CategoryId:  p.CategoryId,
	}
}
