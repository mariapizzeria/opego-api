package validators

import "github.com/go-playground/validator/v10"

func IsValid[T any](body T) error {
	valid := validator.New()
	err := valid.Struct(body)
	return err
}
