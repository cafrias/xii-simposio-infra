package api

import "os"

// Headers represents the headers for a response to API Gateway
type Headers map[string]string

// Body represents the body for all responses to API Gateway
type Body struct {
	LogID string `json:"log_id"`
	Msg   string `json:"message"`
}

// DefaultHeaders generate shared headers for all responses.
func DefaultHeaders() Headers {
	return Headers{
		"Access-Control-Allow-Origin":      os.Getenv("ALLOWED_ORIGINS"),
		"Access-Control-Allow-Credentials": "true",
	}
}
