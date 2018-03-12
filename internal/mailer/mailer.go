package mailer

// Client represents an email client.
type Client interface {
	Send(template string) error
}
