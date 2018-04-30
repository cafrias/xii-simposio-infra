package mailClient

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/friasdesign/xii-simposio-infra/internal/mailer"
)

// New creates a new instace of MailClient
func New() (*MailClient, error) {
	mail := MailClient{
		Sender: aws.String(os.Getenv("EMAIL")),
		Receivers: []*string{
			aws.String(os.Getenv("EMAIL")),
		},
		Charset: aws.String("UTF-8"),
		Subject: aws.String("Nueva Subscripcion XII Simposio"),
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Error while setting up AWS session.")
		return nil, err
	}

	mail.client = ses.New(sess)

	return &mail, nil
}

// MailClient represents an email client.
type MailClient struct {
	Sender    *string
	Receivers []*string
	Charset   *string
	Subject   *string
	client    *ses.SES
}

// Send sends the email with an specified template.
func (m *MailClient) Send(t string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: m.Receivers,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: m.Charset,
					Data:    aws.String(t),
				},
			},
			Subject: &ses.Content{
				Charset: m.Charset,
				Data:    m.Subject,
			},
		},
		Source: m.Sender,
	}

	_, err := m.client.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}

// Check if MailClient implements interface.
var _ mailer.Client = &MailClient{}
