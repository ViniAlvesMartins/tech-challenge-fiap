package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate(dto interface{}) IValidateError {
	var validateError IValidateError
	var errList []Fields
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(dto)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			errList = append(errList, Fields{
				Field:   e.Field(),
				Message: e.Param(),
			})
		}
	}

	validateError.Errors = errList

	return validateError
}
