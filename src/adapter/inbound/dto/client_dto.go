package dto

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/go-playground/validator/v10"
)

type ClientDto struct {
	ID    int    `json:"id"`
	Cpf   int    `json:"cpf" validate:"required" error:"Campo cpf é obrigatorio"`
	Name  string `json:"name" validate:"required" error:"Campo nome é obrigatorio"`
	Email string `json:"email" validate:"required" error:"Campo email é obrigatorio"`
}

type IValidateErrorClient struct {
	Errors []Fields
}

type ClientFields struct {
	Field   string
	Message string
}

func ValidateClient(dto ClientDto) IValidateError {

	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(dto)

	var errList []Fields

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
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

func ConvertClientDtoToDomain(dto ClientDto) domain.Client {

	var client = domain.Client{
		ID:    dto.ID,
		Name:  dto.Name,
		Cpf:   dto.Cpf,
		Email: dto.Email,
	}

	return client
}
