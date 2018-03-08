package client

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

// Client represents a client to the underlying DynamoDB data store.
type Client struct {
	Region              string
	Endpoint            string
	TableName           string
	Now                 func() time.Time
	subscripcionService SubscripcionService
	db                  *dynamodb.DynamoDB
}

// NewClient creates a new Client.
func NewClient() *Client {
	c := &Client{
		Region:    "sa-east-1",
		Endpoint:  "dynamodb.sa-east-1.amazonaws.com",
		TableName: os.Getenv("TABLE_NAME"),
		Now:       time.Now,
	}
	c.subscripcionService.client = c
	return c
}

// Open opens and initializes the DynamoDB connection.
func (c *Client) Open() error {
	// AWS Config
	conf := &aws.Config{
		Region: aws.String(c.Region),
	}
	conf.WithEndpoint(c.Endpoint)

	// AWS session
	sess, err := session.NewSession(conf)
	if err != nil {
		return err
	}

	c.db = dynamodb.New(sess)

	return nil
}

// DB returns the db connection
func (c *Client) DB() *dynamodb.DynamoDB { return c.db }

// SubscripcionService returns the Subscripcion service associated with the client.
func (c *Client) SubscripcionService() simposio.SubscripcionService { return &c.subscripcionService }
