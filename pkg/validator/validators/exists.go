package validators

import (
	"strings"

	"github.com/friasdesign/xii-simposio-infra/pkg/validator"
)

// Exists checks whether i exists.
func Exists(i interface{}) (bool, error) {
	if s, ok := i.(string); ok {
		if len(strings.TrimSpace(s)) > 0 {
			return true, nil
		}
	} else if i != nil {
		return true, nil
	}

	return false, validator.ValidationError{
		Msg:   "Required but value is empty",
		Check: validator.Required,
		Value: i,
	}
}
