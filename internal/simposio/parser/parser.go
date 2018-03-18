package parser

import (
	"encoding/json"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

// ParseSubscripcion parses JSON string into simposio.Suscripcion
func ParseSubscripcion(str string) (*simposio.Subscripcion, error) {
	var subs simposio.Subscripcion

	if err := json.Unmarshal([]byte(str), &subs); err != nil {
		return nil, err
	}

	return &subs, nil
}
