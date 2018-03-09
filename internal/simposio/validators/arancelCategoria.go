package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

// ArancelCategoria validates it represents a valid category.
func ArancelCategoria(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "estudiante_untdf":
		fallthrough
	case "estudiante_otro":
		fallthrough
	case "docente_untdf":
		fallthrough
	case "matriculado_cpcetf":
		fallthrough
	case "general":
		return true
	}

	return false
}
