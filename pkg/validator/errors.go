package validator

// Generic errors
const (
	InvalidType = Error("Invalid type!")
)

// ValidationError represents a validator error.
type ValidationError struct {
	Msg       string
	Check     string
	Value     interface{}
	Threshold interface{}
}

// Error returns the error message.
func (e ValidationError) Error() string { return e.Msg }

// Error represents a generic string error.
type Error string

func (e Error) Error() string { return string(e) }
