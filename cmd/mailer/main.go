package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/friasdesign/xii-simposio-infra/internal/mailer/templates"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler handles the request for new Subscripcion events on DymanoDB.
func Handler(ctx context.Context, e events.DynamoDBEvent) {
	for _, record := range e.Records {
		fmt.Printf("Processing EventID: %v\n", record.EventID)

		// Convert to string map
		smap := make(map[string]string)

		for key, value := range record.Change.NewImage {
			var valStr string
			switch value.DataType() {
			case events.DataTypeBoolean:
				if value.Boolean() {
					valStr = "Si"
				} else {
					valStr = "No"
				}
			case events.DataTypeNumber:
				if key == "arancel_adicional" {
					f, _ := value.Float()
					valStr = strconv.FormatFloat(f, 'f', -1, 64)
				} else {
					f, _ := value.Integer()
					valStr = strconv.Itoa(int(f))
				}
			case events.DataTypeString:
				valStr = value.String()
			}
			var label string
			switch key {
			case "acompanantes":
				label = "Acompañantes"
			default:
				label = strings.Title(strings.Replace(key, "_", " ", -1))
			}

			smap[label] = valStr
		}

		// Add Total
		base := simposio.Aranceles[record.Change.NewImage["arancel_categoria"].String()]
		adicional, _ := record.Change.NewImage["arancel_adicional"].Float()
		smap["Arancel Base"] = strconv.FormatFloat(base, 'f', -1, 64)
		smap["Total"] = strconv.FormatFloat(base+adicional, 'f', -1, 64)

		t, err := template.New("subs").Parse(templates.NewSubscripcion)
		if err != nil {
			fmt.Println("Error while parsing template")
			return
		}
		tb := bytes.NewBufferString("")

		err = t.Execute(tb, smap)

		// Initialize ses
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})
		if err != nil {
			fmt.Println("Error while setting up AWS session.")
		}

		svc := ses.New(sess)

		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{},
				ToAddresses: []*string{
					aws.String("carlos.a.frias@gmail.com"),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Html: &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(tb.String()),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("Nueva Subscripcion XII Simposio"),
				},
			},
			Source: aws.String("carlos.a.frias@gmail.com"),
		}

		result, err := svc.SendEmail(input)
		if err != nil {
			fmt.Println("Error while sending email!")
		}

		fmt.Println("Successfully sent!", result)
	}
}

func main() {
	lambda.Start(Handler)
}
