package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/test"
	"github.com/friasdesign/xii-simposio-infra/internal/validator/validators"
)

func TestDocumento_ValidatesCorrectly(t *testing.T) {
	fix := []test.Case{
		test.Case{
			Input:    "1234",
			Expected: true,
		},
		test.Case{
			Input:    "pepe",
			Expected: false,
		},
		test.Case{
			Input:    "1234asd",
			Expected: false,
		},
		test.Case{
			Input:    "0",
			Expected: true,
		},
	}

	for _, val := range fix {
		s, ok := val.Input.(string)
		if ok == false {
			t.Fatal("Invalid type for Test Case", val.Input)
		}
		o := validators.Documento(s)
		if o != val.Expected {
			t.Fatal("Wrong validation!", s)
		}
	}
}
