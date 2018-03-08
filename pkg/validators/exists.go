package validators

import (
	"strings"
)

// Exists checks whether i exists.
func Exists(i interface{}) bool {
	if s, ok := i.(string); ok {
		return len(strings.TrimSpace(s)) > 0
	}
	return i != nil
}
