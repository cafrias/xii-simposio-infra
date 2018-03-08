package validators

import (
	"errors"
	"strconv"
)

// parseInts parse interface and string to two integers
func parseInts(v interface{}, th string) (int, int, error) {
	vi, ok := v.(int)
	if ok == false {
		return 0, 0, errors.New("Cannot validate 'v', should be of type int")
	}

	thi, err := strconv.Atoi(th)
	if err != nil {
		return 0, 0, errors.New("Cannot parse 'th', should represent value of type int")
	}

	return vi, thi, nil
}
