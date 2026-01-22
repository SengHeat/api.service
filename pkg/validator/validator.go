package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	var errors []string

	for _, e := range validationErrors {
		field := e.Field()
		tag := e.Tag()
		errors = append(errors, fmt.Sprintf("%s: %s", field, getErrorMessage(tag)))
	}

	return fmt.Errorf(strings.Join(errors, ", "))
}

func getErrorMessage(tag string) string {
	switch tag {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email format"
	case "min":
		return "value is too short"
	case "max":
		return "value is too long"
	case "oneof":
		return "invalid value"
	default:
		return "invalid value"
	}
}
