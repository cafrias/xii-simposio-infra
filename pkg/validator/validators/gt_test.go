package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/pkg/validator"

	"github.com/friasdesign/xii-simposio-infra/pkg/test"
	"github.com/friasdesign/xii-simposio-infra/pkg/validator/validators"
)

func TestGt_ValidatesCorrectly(t *testing.T) {
	type Input struct {
		Val interface{}
		Thr string
	}
	fix := []test.Case{
		test.Case{
			Input: Input{
				Val: "asd",
				Thr: "12",
			},
			Expected: false,
			Error:    validator.InvalidType,
		},
		test.Case{
			Input: Input{
				Val: 2,
				Thr: "1",
			},
			Expected: true,
		},
		test.Case{
			Input: Input{
				Val: 0,
				Thr: "1",
			},
			Expected: false,
		},
		test.Case{
			Input: Input{
				Val: 1,
				Thr: "1",
			},
			Expected: false,
		},
	}

	for _, val := range fix {
		in, ok := val.Input.(Input)
		if ok == false {
			t.Fatal("Error while parsin fixture!")
		}
		o, err := validators.Gt(in.Val, in.Thr)
		if val.Error != nil && err != val.Error {
			t.Fatal("Unexpected error!", in.Val, err.Error())
		} else if _, ok := err.(validator.ValidationError); err != nil && ok == false {
			t.Fatal("Unexpected error!", in.Val, err.Error())
		}

		if o != val.Expected {
			t.Fatal("Wrong validation!", val.Input)
		}
	}
}
