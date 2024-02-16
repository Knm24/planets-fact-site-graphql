package validator

import (
	_validator "github.com/go-playground/validator/v10"
)

func New() *_validator.Validate {
	validate := _validator.New()

	return validate
}
