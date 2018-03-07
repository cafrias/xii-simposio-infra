package dynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB represents a connection to a DynamoDB database.
type DynamoDB interface {
	Open() error
	DB() *dynamodb.DynamoDB
}
