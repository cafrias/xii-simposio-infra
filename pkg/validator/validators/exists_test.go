package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/pkg/validator"

	"github.com/friasdesign/xii-simposio-infra/pkg/test"
	"github.com/friasdesign/xii-simposio-infra/pkg/validator/validators"
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
		o, err := validators.Exists(val.Input)
		if _, ok := err.(validator.ValidationError); err != nil && ok == false {
			t.Fatal("Unexpected error!")
		}
		if o != val.Expected {
			t.Fatal("Wrong validation!", val.Input)
		}
	}
}
