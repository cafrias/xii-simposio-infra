package parser

import (
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

// DDBAttributeValueToString converts a DynamoDBAttributeValue into string.
func DDBAttributeValueToString(key string, attr events.DynamoDBAttributeValue) (string, error) {
	switch attr.DataType() {
	case events.DataTypeBinary:
		return string(attr.Binary()), nil
	case events.DataTypeNull:
		return "", nil
	case events.DataTypeBoolean:
		if attr.Boolean() {
			return "Si", nil
		}
		return "No", nil
	case events.DataTypeNumber:
		if key == "arancel_adicional" {
			f, err := attr.Float()
			if err != nil {
				return "", err
			}
			return strconv.FormatFloat(f, 'f', -1, 64), nil
		}
		f, err := attr.Integer()
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(f)), nil
	case events.DataTypeString:
		return attr.String(), nil
	}

	return "", errors.New("Unknown type to parse")
}
