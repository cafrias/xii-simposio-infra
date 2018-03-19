package api

import (
	"os"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/validators"
)

// Headers represents the headers for a response to API Gateway
type Headers map[string]string

// Body represents the body for all responses to API Gateway
type Body struct {
	LogID   string                       `json:"log_id"`
	Msg     string                       `json:"message"`
	Payload interface{}                  `json:"payload,omitempty"`
	Errors  []validators.ValidationError `json:"errors,omitempty"`
}

// DefaultHeaders generate shared headers for all responses.
func DefaultHeaders() Headers {
	return Headers{
		"Access-Control-Allow-Origin":      os.Getenv("ALLOWED_ORIGINS"),
		"Access-Control-Allow-Credentials": "true",
	}
}

// Request represents a request object.
type Request interface {
	GetQuery(p string) string
	GetBody() string
}
