package types

import (
	v "github.com/go-playground/validator"
)

var validator *v.Validate

func init() {
	validator = v.New()
}

// Validate - validates an object based on it's tags
func Validate(t interface{}) error {
	return validator.Struct(t)
}