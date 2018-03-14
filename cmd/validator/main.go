package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/validators"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/parser"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/go-playground/validator.v9"
)

// Response messages.
const (
	ErrValidationMsg         = "Error de validación en Subscripción."
	ErrRequestMsg            = "Petición mal formada, contacte con soporte."
	ErrInternalMsg           = "Error Interno, contacte con soporte."
	ErrSubscripcionExistsMsg = "Ya se registró un usuario con ese Documento."
	SucSavingSubscipcionMsg  = "Subscripción registrada con éxito."
)

// Log messages.
const (
	ErrUUIDLog               = "Couldn't generate UUID."
	ErrRequestLog            = "Request body is invalid!"
	ErrValidationLog         = "Validation Error: %s\n"
	ErrParsingBodyToJSON     = "Error parsing body to JSON"
	ErrDynamoDBConnectionLog = "Error while trying to open connection to DynamoDB"
	ErrSavingSubscripcionLog = "Error while trying to write Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrSubscripcionExistsLog = "Subscripcion with 'Documento' %v already exists!\n"
	SucSavingSubscripcionLog = "Subscripcion with 'Documento' %v successfully saved\n"
)

type validationError struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param,omitempty"`
}

func (v validationError) Error() string {
	return fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", v.Field, v.Tag)
}

type body struct {
	LogID  string            `json:"log_id"`
	Msg    string            `json:"message"`
	Errors []validationError `json:"errors,omitempty"`
}

func handleReq(reqID string, req events.APIGatewayProxyRequest) (int, body, error) {
	b := req.Body
	validate := validators.Initialize()

	// Parse body
	subs, err := parser.Parse(b)
	if err != nil {
		fmt.Println(ErrRequestLog, err.Error())
		return 400, body{LogID: reqID, Msg: ErrRequestMsg}, nil
	}

	// Validate Subscripcion struct
	err = validate.Struct(subs)
	if err != nil {
		var errors []validationError

		for _, err := range err.(validator.ValidationErrors) {
			v := validationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
				Param: err.Param(),
			}
			fmt.Printf(ErrValidationLog, v.Error())
			errors = append(errors, v)
		}

		return 400, body{LogID: reqID, Msg: ErrValidationMsg, Errors: errors}, nil
	}

	// Open DynamoDB connection
	c := client.NewClient()
	err = c.Open()
	if err != nil {
		fmt.Println(ErrDynamoDBConnectionLog)
		return 500, body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	// Write new Subscripcion to DB.
	err = c.SubscripcionService().CreateSubscripcion(subs)
	if err != nil {
		if err == simposio.ErrSubscripcionExists {
			fmt.Printf(ErrSubscripcionExistsLog, subs.Documento)
			return 400, body{LogID: reqID, Msg: ErrSubscripcionExistsMsg}, nil
		}
		fmt.Printf(ErrSavingSubscripcionLog, subs.Documento)
		return 500, body{LogID: reqID, Msg: ErrInternalMsg}, err
	}

	fmt.Printf(SucSavingSubscripcionLog, subs.Documento)

	return 201, body{LogID: reqID, Msg: SucSavingSubscipcionMsg}, nil
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get AWS request ID
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID
	fmt.Printf("Starting request id: %s\n", reqID)

	st, b, err := handleReq(reqID, req)

	// Convert body to JSON.
	bs, jerr := json.Marshal(b)
	if jerr != nil {
		fmt.Println(ErrParsingBodyToJSON)
		st = 500
		bs = []byte{}
		err = jerr
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
		StatusCode: st,
		Body:       string(bs),
	}, err
}

func main() {
	lambda.Start(Handler)
}
