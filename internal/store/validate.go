package store

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateStruct(structValue any) error {
	newValidator := validator.New()

	err := newValidator.Struct(structValue)
	if err == nil {
		return nil
	}

	compiledErrors := []string{}

	errors := err.(validator.ValidationErrors)
	for _, fieldErrors := range errors {
		compiledErrors = append(compiledErrors, customErrorMessage(fieldErrors))
	}

	return fmt.Errorf("validation error: %s", strings.Join(compiledErrors, ", "))
}

func customErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fieldError.Field(), fieldError.Param())
	case "email":
		return "Invalid email format"
		// Add more cases as needed
	}
	return fieldError.Error()
}
