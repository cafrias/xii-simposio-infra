package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/friasdesign/xii-simposio-infra/internal/api"
	"github.com/friasdesign/xii-simposio-infra/internal/api/request"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/messages"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/parser"
)

func handleConfirm(reqID string, req *request.Request) (int, api.Body, error) {
	b := req.GetBody()

	// Parse body
	subs, err := parser.ParseSubscripcion(b)
	if err != nil {
		fmt.Println(messages.ErrRequestLog, err.Error())
		return 400, api.Body{LogID: reqID, Msg: messages.ErrRequestMsg}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	subsServ := c.SubscripcionService()
	if subs.Confirmado {
		err = subsServ.Confirmar(subs.Documento)
	} else {
		err = subsServ.Pendiente(subs.Documento)
	}
	if err != nil {
		if err == simposio.ErrSubscripcionNotFound {
			fmt.Printf(messages.ErrSubscripcionNotFoundLog, subs.Documento)
			return 404, api.Body{LogID: reqID, Msg: messages.ErrSubscripcionNotFoundMsg}, nil
		}
		fmt.Printf(messages.ErrSavingSubscripcionLog, subs.Documento)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	fmt.Printf(messages.SucSavingSubscripcionLog, subs.Documento)
	return 201, api.Body{LogID: reqID, Msg: messages.SucSavingSubscipcionMsg}, nil
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, awsReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get AWS request ID
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID
	req := request.FromAPIGatewayProxyRequest(awsReq)
	fmt.Printf("Starting request id '%s'\n", reqID)

	st, b, err := handleConfirm(reqID, req)

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
