package client_test

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"
)

// Now is the mocked current time for testing.
var Now = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

// Client is a test wrapper for db.Client.
type Client struct {
	*client.Client
}

func NewClient() *Client {
	c := &Client{
		Client: client.NewClient(),
	}
	c.Region = "localhost"
	c.TableName = "subscripciones"
	c.Endpoint = "http://localhost:8000"
	c.Now = func() time.Time { return Now }

	return c
}

// MustOpenClient returns an new, open instance of Client.
func MustOpenClient() *Client {
	c := NewClient()

	// Open connection
	if err := c.Open(); err != nil {
		panic(err)
	}

	// Create table
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("documento"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("documento"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(c.TableName),
	}

	database := c.Client.DB()
	_, err := database.CreateTable(input)
	if err != nil {
		panic(err)
	}

	return c
}

// Close closes the client and removes the underlying database.
func (c *Client) Close() error {
	database := c.Client.DB()
	_, err := database.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(c.TableName),
	})
	if err != nil {
		return err
	}

	return nil
}
