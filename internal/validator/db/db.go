package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/friasdesign/xii-simposio-infra/internal/validator"
)

// SubscripcionService contains CRUD methods for Subscription type.
type SubscripcionService struct {
	DB *dynamodb.DynamoDB
}

// Subscripcion fetches a Subscripcion by Documento.
func (s *SubscripcionService) Subscripcion(doc int) *validator.Subscripcion {
	result, err := s.DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"documento": {
				N: aws.String(string(doc)),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	item := validator.Subscripcion{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &item
}

// CreateSubscripcion creates a new Subscripcion.
func (s *SubscripcionService) CreateSubscripcion(subs validator.Subscripcion) error {
	av, err := dynamodbattribute.MarshalMap(subs)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err = s.DB.PutItem(input)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully added 'Subscripcion', id: %v\n", subs.Documento)
	return nil
}