package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/validator/validators"
)

type Case struct {
	Input    string
	Expected bool
}

func TestEmail_ValidatesCorrectly(t *testing.T) {
	fix := []Case{
		Case{
			Input:    "pepe_pepe@pepe.com",
			Expected: true,
		},
		Case{
			Input:    "pepe_@pepe.com",
			Expected: true,
		},
		Case{
			Input:    "",
			Expected: false,
		},
		Case{
			Input:    "asd",
			Expected: false,
		},
		Case{
			Input:    "mario\\marielo@pepe.com",
			Expected: false,
		},
	}

	for _, val := range fix {
		o := validators.Email(val.Input)
		if o != val.Expected {
			t.Fatal("Wrong validation!", val.Input)
		}
	}
}
