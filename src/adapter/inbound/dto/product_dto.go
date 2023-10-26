package dto

import (
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/go-playground/validator/v10"
)

type ProductDto struct {
	NameProduct string  `json:"name_product"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required,gt=0" error:"O Número deve ser maior que zero"`
	CategoryId  int     `json:"category_id" validate:"required"`
}

var validate *validator.Validate

type IValidateError struct {
	Errors []Fields
}

type Fields struct {
	Field   string
	Message string
}

func ValidateProduct(dto ProductDto) IValidateError {
	var validateError IValidateError
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(dto)

	var errList []Fields

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			errList = append(errList, Fields{
				Field:   err.Field(),
				Message: err.Param(),
			})
		}
	}

	validateError.Errors = errList

	return validateError
}

func ConvertDtoToDomain(dto ProductDto) entity.Product {
	var product = entity.Product{
		NameProduct: dto.NameProduct,
		Description: dto.Description,
		Price:       dto.Price,
		CategoryId:  dto.CategoryId,
	}

	return product
}
