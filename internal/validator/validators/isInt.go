package validators

import (
	"strconv"
)

// IsInt validates whether string i represents valid Documento.
func IsInt(i string) bool {
	// Convert to int
	_, err := strconv.Atoi(i)
	if err != nil {
		return false
	}
	return true
}
