package validators

import "fmt"

// Represents errors to log.
const (
	ErrValidationLog = "Validation Error: %s\n"
)

// ValidationError represents a validation error.
type ValidationError struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param,omitempty"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", v.Field, v.Tag)
}
