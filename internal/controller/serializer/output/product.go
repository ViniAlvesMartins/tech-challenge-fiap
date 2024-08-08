package output

import "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"

type ProductDto struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CategoryId  int     `json:"category_id"`
	Active      bool    `json:"active"`
}

func ProductFromEntity(product entity.Product) ProductDto {
	return ProductDto{
		ID:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CategoryId:  product.CategoryId,
		Active:      product.Active,
	}
}

func ProductListFromEntity(products []entity.Product) []ProductDto {
	var productsDto []ProductDto
	for _, p := range products {
		productsDto = append(productsDto, ProductFromEntity(p))
	}

	return productsDto
}
