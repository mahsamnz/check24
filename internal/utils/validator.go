package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(object interface{}) error {
	validate := validator.New()

	err := validate.Struct(object)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, vErr := range validationErrors {
				errMsg += fmt.Sprintf("Field: %s, Tag: %s, Actual Value: %v\n",
					vErr.Field(), vErr.Tag(), vErr.Value())
			}
			return fmt.Errorf("validation errors: \n%s", errMsg)
		}
		return err
	}
	return nil
}
