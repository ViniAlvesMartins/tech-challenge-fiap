package dto

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/go-playground/validator/v10"
)

type ProductDto struct {
	ID          int     `json:"id"`
	NameProduct string  `json:"name_product"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required,gt=0" error:"O Número deve ser maior que zero"`
	CategoryId  int     `json:"category_id" validate:"required"`
	Active      bool    `json:"active"`
}

var validate *validator.Validate

type IValidateError struct {
	Errors []Fields
}

type Fields struct {
	Field   string
	Message string
}

func ConvertDtoToDomain(dto ProductDto) entity.Product {
	var product = entity.Product{
		ID:          dto.ID,
		NameProduct: dto.NameProduct,
		Description: dto.Description,
		Price:       dto.Price,
		CategoryId:  dto.CategoryId,
		Active:      dto.Active,
	}

	return product
}
