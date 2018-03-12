package mailer

import (
	"os"

	"github.com/friasdesign/xii-simposio-infra/internal/mailer/templates"
)

// Mailer represents the mailer service.
type Mailer struct {
	Sender     string
	Recipients []string
	Subject    string
	HTMLBody   string
	Client     client
}

func (m *Mailer) Send() {
	// Parse
}

// NewSubscripcionEvent creates Mailer with preconfigured data for handling a new Subscripcion Event.
func NewSubscripcionEvent() *Mailer {
	return &Mailer{
		Sender: os.Getenv("EMAIL_BOT"),
		Recipients: []string{
			os.Getenv("EMAIL_SIMPOSIO"),
		},
		Subject:  "Subscripción a XII Simposio",
		HTMLBody: templates.NewSubscripcion,
	}
}

// Check interface implementation.
// var _ Service = &Mailer{}

// Service represents a mailer service.
type Service interface {
	Send(map[string]interface{}) error
}

type client struct {
}
