package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

// Initialize initializes validators.
func Initialize() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("arancel", ArancelCategoria)

	return validate
}
