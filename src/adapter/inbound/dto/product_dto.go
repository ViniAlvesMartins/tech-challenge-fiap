package dto

import (
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
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

func ValidateProduct(dto ProductDto) IValidateError {

	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(dto)

	var errList []Fields

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			errList = append(errList, Fields{
				Field:   err.Field(),
				Message: err.Param(),
			})
		}

		// from here you can create your own error messages in whatever language you wish
	}

	var validateError IValidateError

	validateError.Errors = errList

	return validateError
}

func ConvertDtoToDomain(dto ProductDto) domain.Product {

	var product = domain.Product{
		ID:          dto.ID,
		NameProduct: dto.NameProduct,
		Description: dto.Description,
		Price:       dto.Price,
		CategoryId:  dto.CategoryId,
		Active:      dto.Active,
	}

	return product
}
