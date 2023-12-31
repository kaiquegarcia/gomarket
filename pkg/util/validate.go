package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(input interface{}) ([]string, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		fmt.Printf("could not validate the input. check the error:\n%s\n", err.Error())
		errorList := make([]string, 0)
		if errors, castable := err.(validator.ValidationErrors); castable {
			for _, err := range errors {
				errorList = append(errorList, err.Error())
			}
		}

		return errorList, err
	}

	return make([]string, 0), nil
}
