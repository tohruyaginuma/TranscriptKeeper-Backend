package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Valid struct {
	validate *validator.Validate
}

func NewValidator() *Valid {
	validator := validator.New()

	return &Valid{
		validate: validator,
	}
}

func (v Valid) IsValid(value any) error {
	err := v.validate.Struct(value)
	if err == nil {
		return nil
	}

	return fmt.Errorf("invalid: %w", err)
}
