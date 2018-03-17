package client

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

// UpdateSubscripcion updates a Subscripcion.
func (s *SubscripcionService) UpdateSubscripcion(subs *simposio.Subscripcion) error {
	av, err := dynamodbattribute.MarshalMap(subs)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		ConditionExpression: aws.String("documento = :doc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":doc": {
				N: aws.String(strconv.Itoa(subs.Documento)),
			},
		},
		Item:      av,
		TableName: aws.String(s.client.TableName),
	}
	_, err = s.client.db.PutItem(input)
	if err != nil {
		if _, ok := err.(awserr.RequestFailure); ok {
			return simposio.ErrSubscripcionNotFound
		}
		return err
	}

	return nil
}

// DeleteSubscripcion removes a Subscripcion.
func (s *SubscripcionService) DeleteSubscripcion(doc int) error {
	docStr := aws.String(strconv.Itoa(doc))
	input := &dynamodb.DeleteItemInput{
		ConditionExpression: aws.String("documento = :doc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":doc": {
				N: docStr,
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"documento": {
				N: docStr,
			},
		},
		TableName: aws.String(s.client.TableName),
	}
	_, err := s.client.db.DeleteItem(input)
	if err != nil {
		if _, ok := err.(awserr.RequestFailure); ok {
			return simposio.ErrSubscripcionNotFound
		}
		return err
	}

	return nil
}

// Confirmar updates field Confirmado for given Subscripcion.
func (s *SubscripcionService) Confirmar(doc int) error {
	return s.setConfirmado(doc, true)
}

// Pendiente updates field Confirmado for given Subscripcion.
func (s *SubscripcionService) Pendiente(doc int) error {
	return s.setConfirmado(doc, false)
}

func (s *SubscripcionService) setConfirmado(doc int, val bool) error {
	docStr := aws.String(strconv.Itoa(doc))
	input := &dynamodb.UpdateItemInput{
		ConditionExpression: aws.String("documento = :doc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":doc": {
				N: docStr,
			},
			":conf": {
				BOOL: aws.Bool(val),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"documento": {
				N: docStr,
			},
		},
		TableName:        aws.String(s.client.TableName),
		UpdateExpression: aws.String("set confirmado = :conf"),
	}
	_, err := s.client.db.UpdateItem(input)
	if err != nil {
		if _, ok := err.(awserr.RequestFailure); ok {
			return simposio.ErrSubscripcionNotFound
		}
		return err
	}

	return nil
}

// Subscripciones returns all 'Subscripcion' items in database.
func (s *SubscripcionService) Subscripciones() ([]*simposio.Subscripcion, error) {
	expr, err := expression.NewBuilder().Build()
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(s.client.TableName),
	}
	result, err := s.client.db.Scan(input)
	if err != nil {
		return nil, err
	}

	var ret []*simposio.Subscripcion
	for _, i := range result.Items {
		var item simposio.Subscripcion
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, err
		}

		ret = append(ret, &item)
	}

	return ret, nil
}
