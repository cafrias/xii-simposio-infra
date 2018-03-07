package validator

// Subscripcion errors.
const (
	ErrSubscripcionNotFound = Error("Subscripcion not found")
	ErrSubscripcionExists   = Error("Subscripcion already exists")
)

// Error represents a validator error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
