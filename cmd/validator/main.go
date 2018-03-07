package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response represents a response for AWS Lambda.
type Response struct {
	Message string `json:"message"`
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body: %v\n", req.Body)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Everything OK",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
