package dynamodb

// DynamoDB represents a connection to a DynamoDB database.
type DynamoDB interface {
	Open() error
}
