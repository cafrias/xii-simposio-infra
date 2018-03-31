package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"

	"github.com/friasdesign/xii-simposio-infra/internal/mailer/mailClient"
	"github.com/friasdesign/xii-simposio-infra/internal/mailer/parser"
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
			fmt.Printf("Processing key '%s' with type '%v'\n", key, value.DataType())
			// Parse DynamoDBAttribute to string
			valStr, err := parser.DDBAttributeValueToString(key, value)
			if err != nil {
				fmt.Print("Error while parsing DDBAttributeValue to String", err)
			}

			// Humanize label
			label := parser.HumanizeLabel(key)

			smap[label] = valStr
		}

		// Add Total
		base := simposio.Aranceles[record.Change.NewImage["arancel_categoria"].String()]
		adicional, _ := record.Change.NewImage["arancel_adicional"].Float()
		smap["Arancel Base"] = strconv.FormatFloat(base, 'f', -1, 64)
		smap["Total"] = strconv.FormatFloat(base+adicional, 'f', -1, 64)

		// Parse Template
		tStr, err := templates.ParseNewSubscripcion(smap)
		if err != nil {
			fmt.Println("Error while parsing template")
			return
		}

		// Initialize ses
		mailCli, err := mailClient.New()
		if err != nil {
			fmt.Println("Error while creating a new email client.", err)
			return
		}

		// Send email
		err = mailCli.Send(tStr)
		if err != nil {
			fmt.Println("Error while sending email.")
			return
		}

		fmt.Println("Successfully sent!")
	}
}

func main() {
	lambda.Start(Handler)
}
