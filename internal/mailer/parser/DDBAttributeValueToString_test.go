package parser_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/friasdesign/xii-simposio-infra/internal/mailer/parser"
)

func TestDDBAttributeValueToString_ConvertsCorrectly(t *testing.T) {
	type testCase struct {
		Key       string
		AttrType  string
		AttrValue string
		Output    string
		Error     error
	}
	fix := []testCase{
		testCase{
			Key:       "arancel_adicional",
			AttrType:  "N",
			AttrValue: "\"21.21\"",
			Output:    "21.21",
			Error:     nil,
		},
		testCase{
			Key:       "algo",
			AttrType:  "BOOL",
			AttrValue: "false",
			Output:    "No",
			Error:     nil,
		},
		testCase{
			Key:       "otro",
			AttrType:  "BOOL",
			AttrValue: "true",
			Output:    "Si",
			Error:     nil,
		},
		testCase{
			Key:       "otro",
			AttrType:  "NULL",
			AttrValue: "\"\"",
			Output:    "",
			Error:     nil,
		},
		testCase{
			Key:       "otro",
			AttrType:  "S",
			AttrValue: "\"texto\"",
			Output:    "texto",
			Error:     nil,
		},
		testCase{
			Key:       "integer",
			AttrType:  "N",
			AttrValue: "\"22\"",
			Output:    "22",
			Error:     nil,
		},
	}

	for _, item := range fix {
		jsonStr := fmt.Sprintf("{\n \"%v\": %s\n}", item.AttrType, item.AttrValue)

		fmt.Println(jsonStr)

		var ddbAttr events.DynamoDBAttributeValue
		err := ddbAttr.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			t.Fatal("Something went wrong while parsing JSON", err)
		}
		pStr, err := parser.DDBAttributeValueToString(item.Key, ddbAttr)
		if err != nil {
			t.Fatal("Unexpected error!", err)
		}

		if pStr != item.Output {
			t.Fatal("Invalid output", "Expected: ", item.Output, "Received: ", pStr)
		}
	}
}
