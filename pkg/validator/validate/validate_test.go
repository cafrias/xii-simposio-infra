package validate_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/pkg/validator"
)

type Case struct {
	Tag      string
	Val      interface{}
	Expected bool
}

func TestValidateFromTag_ReturnTrueIfOK(t *testing.T) {
	fix := []Case{
		Case{
			Tag:      "required",
			Val:      "asd",
			Expected: true,
		},
		Case{
			Tag:      "required,gt=1",
			Val:      2,
			Expected: true,
		},
		Case{
			Tag:      "gte=0,lt=10",
			Val:      8,
			Expected: true,
		},
		Case{
			Tag:      "email",
			Val:      "pepe@pepe.com",
			Expected: true,
		},
	}

	for _, val := range fix {
		out, err := validator.ValidateFromTag(val.Tag, val.Val)
		if err != nil {
			t.Fatal("Shouldn't return error!")
		}
		if out != val.Expected {
			t.Fatal("Expected ", val.Expected)
		}
	}
}
