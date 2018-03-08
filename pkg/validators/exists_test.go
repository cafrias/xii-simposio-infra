package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/test"
	"github.com/friasdesign/xii-simposio-infra/internal/validator/validators"
)

func TestExists_ValidatesCorrectly(t *testing.T) {
	fix := []test.Case{
		test.Case{
			Input:    "asd",
			Expected: true,
		},
		test.Case{
			Input:    123,
			Expected: true,
		},
		test.Case{
			Input:    "",
			Expected: false,
		},
		test.Case{
			Input:    nil,
			Expected: false,
		},
		test.Case{
			Input: " 	",
			Expected: false,
		},
	}

	for _, val := range fix {
		o := validators.Exists(val.Input)
		if o != val.Expected {
			t.Fatal("Wrong validation!", val.Input)
		}
	}
}
