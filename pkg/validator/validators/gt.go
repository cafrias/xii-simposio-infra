package validators

import "github.com/friasdesign/xii-simposio-infra/pkg/validator"

// Gt checks that v is greater than th.
func Gt(v interface{}, th string) (bool, error) {
	vi, thi, err := parseInts(v, th)

	if err != nil {
		return false, err
	}

	if vi <= thi {
		return false, validator.ValidationError{
			Msg:       "Isn't greater than threshold",
			Check:     validator.Gt,
			Value:     vi,
			Threshold: thi,
		}
	}

	return true, nil
}
