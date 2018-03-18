package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/friasdesign/xii-simposio-infra/internal/api"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/validators"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/parser"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response messages.
const (
	ErrValidationMsg           = "Error de validación en Subscripción."
	ErrRequestMsg              = "Petición mal formada, contacte con soporte."
	ErrInternalMsg             = "Error Interno, contacte con soporte."
	ErrSubscripcionExistsMsg   = "Ya se registró un usuario con ese Documento."
	ErrSubscripcionNotFoundMsg = "No se encuentra subscripción con ese Documento."
	ErrQueryParamDocInvalidMsg = "Debe indicar un documento valido para buscar."
	SucSavingSubscipcionMsg    = "Subscripción registrada con éxito."
	SucFetchingSubscripcionMsg = "Subscripción encontrada con éxito."
)

// Log messages.
const (
	ErrUUIDLog                 = "Couldn't generate UUID."
	ErrRequestLog              = "Request body is invalid!"
	ErrValidationLog           = "Validation Error\n"
	ErrParsingBodyToJSON       = "Error parsing body to JSON"
	ErrDynamoDBConnectionLog   = "Error while trying to open connection to DynamoDB"
	ErrSavingSubscripcionLog   = "Error while trying to write Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrFetchingSubscripcionLog = "Error while trying to fetch Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrSubscripcionExistsLog   = "Subscripcion with 'Documento' %v already exists!\n"
	ErrSubscripcionNotFoundLog = "Subscripcion with 'Documento' %v not found!\n"
	ErrUnexpectedHTTPMethodLog = "Unexpected HTTP method '%s'\n"
	ErrQueryParamDocInvalidLog = "Invalid 'doc' query param '%v'\n"
	SucSavingSubscripcionLog   = "Subscripcion with 'Documento' %v successfully saved\n"
	SucFetchingSubscripcionLog = "Subscripcion with 'Documento' %v successfully fetched\n"
)

func handlePOST(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	b := req.Body

	// Parse body
	subs, err := parser.ParseSubscripcion(b)
	if err != nil {
		fmt.Println(ErrRequestLog, err.Error())
		return 400, api.Body{LogID: reqID, Msg: ErrRequestMsg}, nil
	}

	// Validate Subscripcion struct
	errors := validators.ValidateSubscripcion(subs)
	if errors != nil {
		return 400, api.Body{LogID: reqID, Msg: ErrValidationMsg, Errors: errors}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	err = c.SubscripcionService().CreateSubscripcion(subs)
	if err != nil {
		if err == simposio.ErrSubscripcionExists {
			fmt.Printf(ErrSubscripcionExistsLog, subs.Documento)
			return 400, api.Body{LogID: reqID, Msg: ErrSubscripcionExistsMsg}, nil
		}
		fmt.Printf(ErrSavingSubscripcionLog, subs.Documento)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	fmt.Printf(SucSavingSubscripcionLog, subs.Documento)

	return 201, api.Body{LogID: reqID, Msg: SucSavingSubscipcionMsg}, nil
}

func handlePUT(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	b := req.Body

	// Parse body
	subs, err := parser.ParseSubscripcion(b)
	if err != nil {
		fmt.Println(ErrRequestLog, err.Error())
		return 400, api.Body{LogID: reqID, Msg: ErrRequestMsg}, nil
	}

	// Validate Subscripcion struct
	errors := validators.ValidateSubscripcion(subs)
	if errors != nil {
		return 400, api.Body{LogID: reqID, Msg: ErrValidationMsg, Errors: errors}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	err = c.SubscripcionService().UpdateSubscripcion(subs)
	if err != nil {
		if err == simposio.ErrSubscripcionNotFound {
			fmt.Printf(ErrSubscripcionNotFoundLog, subs.Documento)
			return 400, api.Body{LogID: reqID, Msg: ErrSubscripcionNotFoundMsg}, nil
		}
		fmt.Printf(ErrSavingSubscripcionLog, subs.Documento)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	fmt.Printf(SucSavingSubscripcionLog, subs.Documento)
	return 201, api.Body{LogID: reqID, Msg: SucSavingSubscipcionMsg}, nil
}

func handleGET(reqID string, req events.APIGatewayProxyRequest) (int, api.Body, error) {
	docStr := req.QueryStringParameters["doc"]
	doc, err := strconv.Atoi(docStr)
	if err != nil {
		fmt.Printf(ErrQueryParamDocInvalidLog, docStr)
		return 400, api.Body{LogID: reqID, Msg: ErrQueryParamDocInvalidMsg}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(ErrDynamoDBConnectionLog)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	subs, err := c.SubscripcionService().Subscripcion(doc)
	if err != nil {
		if err == simposio.ErrSubscripcionNotFound {
			fmt.Printf(ErrSubscripcionNotFoundLog, doc)
			return 400, api.Body{LogID: reqID, Msg: ErrSubscripcionNotFoundMsg}, nil
		}
		fmt.Printf(ErrFetchingSubscripcionLog, doc)
		return 500, api.Body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	fmt.Printf(SucFetchingSubscripcionLog, subs.Documento)
	return 200, api.Body{LogID: reqID, Msg: SucFetchingSubscripcionMsg, Payload: subs}, nil
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
	default:
		fmt.Printf(ErrUnexpectedHTTPMethodLog, HTTPMethod)
		st = 400
		b = api.Body{}
		err = errors.New("Unexpected HTTP method")
	}

	// Convert body to JSON.
	bs, jerr := json.Marshal(b)
	if jerr != nil {
		fmt.Println(ErrParsingBodyToJSON)
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
