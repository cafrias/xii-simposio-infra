package client

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

// Ensure DialService implements wtf.DialService.
var _ simposio.SubscripcionService = &SubscripcionService{}

// SubscripcionService contains CRUD methods for Subscription type.
type SubscripcionService struct {
	client *Client
}

// Subscripcion fetches a Subscripcion by Documento.
func (s *SubscripcionService) Subscripcion(doc int) (*simposio.Subscripcion, error) {
	result, err := s.client.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.client.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"documento": {
				N: aws.String(strconv.Itoa(doc)),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(result.Item) == 0 {
		return nil, simposio.ErrSubscripcionNotFound
	}

	item := simposio.Subscripcion{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &item, nil
}

// CreateSubscripcion creates a new Subscripcion.
func (s *SubscripcionService) CreateSubscripcion(subs *simposio.Subscripcion) error {
	av, err := dynamodbattribute.MarshalMap(subs)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(documento)"),
		Item:                av,
		TableName:           aws.String(s.client.TableName),
	}
	_, err = s.client.db.PutItem(input)
	if err != nil {
		if _, ok := err.(awserr.RequestFailure); ok {
			return simposio.ErrSubscripcionExists
		}
		return err
	}

	return nil
}
