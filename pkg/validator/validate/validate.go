package validate

import (
	"strings"

	"github.com/friasdesign/xii-simposio-infra/pkg/validator"
	"github.com/friasdesign/xii-simposio-infra/pkg/validator/validators"
)

// ValidateFromTag executes validators specified in a tag over value.
func ValidateFromTag(tag string, value interface{}) (bool, error) {
	args := strings.Split(tag, ",")

	for _, arg := range args {
		vals := strings.Split(arg, "=")
		switch vals[0] {
		case validator.Required:
			return validators.Exists(value)
		case validator.Gt:
			return validators.Gt(value, vals[1])
		case "gte":
			vi, thi, err := parseInts(value, vals[1])
			if err != nil {
				return false, err
			}

			if vi < thi {
				return false, ValidationError{
					Msg:       "Isn't greater than or equals threshold",
					Check:     vals[0],
					Value:     value,
					Threshold: vals[1],
				}
			}
		case "lt":
			vi, thi, err := parseInts(value, vals[1])
			if err != nil {
				return false, err
			}

			if vi >= thi {
				return false, ValidationError{
					Msg:       "Isn't less than threshold",
					Check:     vals[0],
					Value:     value,
					Threshold: vals[1],
				}
			}
		case "lte":
			vi, thi, err := parseInts(value, vals[1])
			if err != nil {
				return false, err
			}

			if vi > thi {
				return false, ValidationError{
					Msg:       "Isn't less than or equals threshold",
					Check:     vals[0],
					Value:     value,
					Threshold: vals[1],
				}
			}
		// case "min":
		// 	ok := validators.Min(value, arg[1])
		// 	if ok == false {
		// 		return false, InvalidMin
		// 	}
		// case "max":
		// 	ok := validators.Max(value, arg[1])
		// 	if ok == false {
		// 		return false, InvalidMax
		// 	}
		// }
	}
	return true, nil
}
