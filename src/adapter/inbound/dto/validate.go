package dto

import (
	"github.com/go-playground/validator/v10"
)

func Validate(dto interface{}) IValidateError {
	var validateError IValidateError
	var errList []Fields

	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(dto); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errList = append(errList, Fields{
				Field:   e.Field(),
				Message: e.Param(),
			})
		}
	}

	validateError.Errors = errList

	return validateError
}
