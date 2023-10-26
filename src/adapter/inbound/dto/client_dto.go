package dto

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/go-playground/validator/v10"
)

type ClientDto struct {
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
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			errList = append(errList, Fields{
				Field:   err.Field(),
				Message: err.Param(),
			})
		}
	}

	var validateError IValidateError

	validateError.Errors = errList

	return validateError
}

func ConvertClientDtoToDomain(dto ClientDto) entity.Client {

	var client = entity.Client{
		Name:  dto.Name,
		Cpf:   dto.Cpf,
		Email: dto.Email,
	}

	return client
}
