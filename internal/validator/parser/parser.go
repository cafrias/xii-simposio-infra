package parser

import (
	"encoding/json"

	"github.com/friasdesign/xii-simposio-infra/internal/validator"
)

// Parse parses JSON string into validator.Suscripcion
func Parse(str string) (*validator.Subscripcion, error) {
	var subs validator.Subscripcion

	if err := json.Unmarshal([]byte(str), &subs); err != nil {
		return nil, err
	}

	return &subs, nil
}
