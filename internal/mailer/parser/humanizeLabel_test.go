package parser_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/mailer/parser"
)

func TestHumanizeLabel_ConvertsCorrectly(t *testing.T) {
	type testCase struct {
		Key    string
		Output string
	}
	fix := []testCase{
		testCase{
			Key:    "acompanantes",
			Output: "Acompañantes",
		},
		testCase{
			Key:    "algun_titulo",
			Output: "Algun Titulo",
		},
	}

	for _, item := range fix {
		out := parser.HumanizeLabel(item.Key)
		if out != item.Output {
			t.Fatal("Wrong output!", "Expected: ", item.Output, "Received: ", out)
		}
	}
}
