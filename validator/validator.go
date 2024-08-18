package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validator validator.Validate

func InitInstanceValidator() {
	instance := validator.New()

	Validator = *instance
}

func InitValidator() error {
	InitInstanceValidator()

	if err := RegisterCustomValidators(); err != nil {
		return err
	}

	return nil
}

func RegisterCustomValidators() error {
	err := Validator.RegisterValidation("user-exist", func(fl validator.FieldLevel) bool {
		return CheckUserExist(fl)
	})

	if err != nil {
		return err
	}

	return nil
}
