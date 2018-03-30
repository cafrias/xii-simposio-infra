package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/messages"

	"github.com/friasdesign/xii-simposio-infra/internal/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/friasdesign/xii-simposio-infra/internal/api/request"
)

type body struct {
	Message string
}

func handleListing(reqID string, variant string) (int, api.Body, error) {
	// Open DynamoDB connection
	c := client.NewClient()
	err := c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}
	sServ := c.SubscripcionService()

	var result []*simposio.Subscripcion
	switch variant {
	case "pending":
		result, err = sServ.SubscripcionesPendientes()
	case "all":
		result, err = sServ.Subscripciones()
	default:
		result, err = sServ.SubscripcionesConfirmadas()
	}

	if err != nil {
		fmt.Println(messages.ErrDynamoDBFetchingLog, err)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	fmt.Println(messages.SucFetchingListLog)

	return 200, api.Body{LogID: reqID, Msg: messages.SucFetchingListMsg, Payload: result}, nil
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, awsReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get AWS request ID
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID
	req := request.FromAPIGatewayProxyRequest(awsReq)
	fmt.Printf("Starting request id '%s'\n", reqID)

	query := req.GetQuery("query")

	var st int
	var b api.Body
	var err error
	switch query {
	case "pending":
		fallthrough
	case "confirmed":
		fallthrough
	case "all":
		st, b, err = handleListing(reqID, query)
	default:
		st = 400
		b = api.Body{
			LogID: reqID,
			Msg:   messages.ErrRequestMsg,
		}
		fmt.Printf(messages.ErrQueryParamInvalidLog, "query", query)
	}

	// Convert body to JSON.
	bs, jerr := json.Marshal(b)
	if jerr != nil {
		fmt.Println(messages.ErrParsingBodyToJSON)
		st = 500
		bs = []byte{}
		err = jerr
	}

	return events.APIGatewayProxyResponse{
		Headers:    api.DefaultHeaders(),
		StatusCode: st,
		Body:       string(bs),
	}, err
}

func main() {
	lambda.Start(Handler)
}
