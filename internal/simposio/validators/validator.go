package validators

import (
	"fmt"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"gopkg.in/go-playground/validator.v9"
)

// ValidateSubscripcion validates 'Subscripcion'
func ValidateSubscripcion(subs *simposio.Subscripcion) []ValidationError {
	validate := validator.New()
	validate.RegisterValidation("arancel", ArancelCategoria)

	err := validate.Struct(subs)
	if err != nil {
		var errors []ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			v := ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
				Param: err.Param(),
			}
			fmt.Printf(ErrValidationLog, v.Error())
			errors = append(errors, v)
		}

		return errors
	}

	return nil
}
