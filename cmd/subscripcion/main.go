package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/friasdesign/xii-simposio-infra/cmd/subscripcion/messages"
	"github.com/friasdesign/xii-simposio-infra/internal/api"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/validators"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/parser"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlePOST(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	b := req.Body

	// Parse body
	subs, err := parser.ParseSubscripcion(b)
	if err != nil {
		fmt.Println(messages.ErrRequestLog, err.Error())
		return 400, api.Body{LogID: reqID, Msg: messages.ErrRequestMsg}, nil
	}

	// Validate Subscripcion struct
	errors := validators.ValidateSubscripcion(subs)
	if errors != nil {
		return 400, api.Body{LogID: reqID, Msg: messages.ErrValidationMsg, Errors: errors}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	err = c.SubscripcionService().CreateSubscripcion(subs)
	if err != nil {
		if err == simposio.ErrSubscripcionExists {
			fmt.Printf(messages.ErrSubscripcionExistsLog, subs.Documento)
			return 400, api.Body{LogID: reqID, Msg: messages.ErrSubscripcionExistsMsg}, nil
		}
		fmt.Printf(messages.ErrSavingSubscripcionLog, subs.Documento)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	fmt.Printf(messages.SucSavingSubscripcionLog, subs.Documento)

	return 201, api.Body{LogID: reqID, Msg: messages.SucSavingSubscipcionMsg}, nil
}

func handlePUT(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	b := req.Body

	// Parse body
	subs, err := parser.ParseSubscripcion(b)
	if err != nil {
		fmt.Println(messages.ErrRequestLog, err.Error())
		return 400, api.Body{LogID: reqID, Msg: messages.ErrRequestMsg}, nil
	}

	// Validate Subscripcion struct
	errors := validators.ValidateSubscripcion(subs)
	if errors != nil {
		return 400, api.Body{LogID: reqID, Msg: messages.ErrValidationMsg, Errors: errors}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	err = c.SubscripcionService().UpdateSubscripcion(subs)
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

func handleGET(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	docStr := req.QueryStringParameters["doc"]
	doc, err := strconv.Atoi(docStr)
	if err != nil {
		fmt.Printf(messages.ErrQueryParamDocInvalidLog, docStr)
		return 400, api.Body{LogID: reqID, Msg: messages.ErrQueryParamDocInvalidMsg}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	subs, err := c.SubscripcionService().Subscripcion(doc)
	if err != nil {
		if err == simposio.ErrSubscripcionNotFound {
			fmt.Printf(messages.ErrSubscripcionNotFoundLog, doc)
			return 404, api.Body{LogID: reqID, Msg: messages.ErrSubscripcionNotFoundMsg}, nil
		}
		fmt.Printf(messages.ErrFetchingSubscripcionLog, doc)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	fmt.Printf(messages.SucFetchingSubscripcionLog, subs.Documento)
	return 200, api.Body{LogID: reqID, Msg: messages.SucFetchingSubscripcionMsg, Payload: subs}, nil
}

func handleDELETE(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	docStr := req.QueryStringParameters["doc"]
	doc, err := strconv.Atoi(docStr)
	if err != nil {
		fmt.Printf(messages.ErrQueryParamDocInvalidLog, docStr)
		return 400, api.Body{LogID: reqID, Msg: messages.ErrQueryParamDocInvalidMsg}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(messages.ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	err = c.SubscripcionService().DeleteSubscripcion(doc)
	if err != nil {
		if err == simposio.ErrSubscripcionNotFound {
			fmt.Printf(messages.ErrSubscripcionNotFoundLog, doc)
			return 404, api.Body{LogID: reqID, Msg: messages.ErrSubscripcionNotFoundMsg}, nil
		}
		fmt.Printf(messages.ErrDeletingSubscripcionLog, doc)
		return 500, api.Body{LogID: reqID, Msg: messages.ErrInternalMsg}, err
	}

	fmt.Printf(messages.SucDeletingSubscripcionLog, doc)
	return 200, api.Body{LogID: reqID, Msg: messages.SucDeletingSubscripcionMsg}, nil
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get AWS request ID
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID
	HTTPMethod := req.HTTPMethod
	fmt.Printf("Starting request id '%s' with HTTP method '%s'\n", reqID, HTTPMethod)

	var st int
	var b api.Body
	var err error
	switch HTTPMethod {
	case "POST":
		st, b, err = handlePOST(reqID, req)
	case "PUT":
		st, b, err = handlePUT(reqID, req)
	case "GET":
		st, b, err = handleGET(reqID, req)
	case "DELETE":
		st, b, err = handleDELETE(reqID, req)
	default:
		fmt.Printf(messages.ErrUnexpectedHTTPMethodLog, HTTPMethod)
		st = 400
		b = api.Body{}
		err = errors.New("Unexpected HTTP method")
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
