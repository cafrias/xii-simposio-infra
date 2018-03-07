package validators

import (
	"strconv"
)

// Documento validates whether string i represents valid Documento.
func Documento(i string) bool {
	// Convert to int
	_, err := strconv.Atoi(i)
	if err != nil {
		return false
	}
	return true
}
