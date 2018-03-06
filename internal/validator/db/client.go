package db

import (
	"os"
	"time"

	"github.com/friasdesign/xii-simposio-infra/internal/validator"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Client represents a client to the underlying DynamoDB data store.
type Client struct {
	Region              string
	TableName           string
	Now                 func() time.Time
	subscripcionService SubscripcionService
	db                  *dynamodb.DynamoDB
}

// NewClient creates a new Client.
func NewClient() *Client {
	c := &Client{
		Region:    "sa-east-1",
		TableName: os.Getenv("TABLE_NAME"),
		Now:       time.Now,
	}
	c.subscripcionService.client = c
	return c
}

// Open opens and initializes the DynamoDB connection.
func (c *Client) Open() error {
	// AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.Region),
	})
	if err != nil {
		return err
	}

	c.db = dynamodb.New(sess)

	return nil
}

// SubscripcionService returns the Subscripcion service associated with the client.
func (c *Client) SubscripcionService() validator.SubscripcionService { return &c.subscripcionService }
