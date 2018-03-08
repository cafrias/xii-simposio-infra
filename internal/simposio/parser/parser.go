package parser

import (
	"encoding/json"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

// Parse parses JSON string into simposio.Suscripcion
func Parse(str string) (*simposio.Subscripcion, error) {
	var subs simposio.Subscripcion

	if err := json.Unmarshal([]byte(str), &subs); err != nil {
		return nil, err
	}

	return &subs, nil
}
