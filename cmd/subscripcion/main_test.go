package main_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/friasdesign/xii-simposio-infra/cmd/subscripcion/messages"
	"gopkg.in/gavv/httpexpect.v1"
)

const HTTPEndpoint = "http://127.0.0.1:3000"
const tableName = "xii-simposio-dev-subscripciones"

func tearDown(t *testing.T) {
	svc := dynamodb.New(session.New(&aws.Config{
		Region: aws.String("sa-east-1"),
	}))
	sOut, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		t.Fatal("Error while scanning DynamoDB table, ", err)
	}

	var wReqs []*dynamodb.WriteRequest
	for _, item := range sOut.Items {
		wReqs = append(wReqs, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					"documento": item["documento"],
				},
			},
		})
	}
	// items := dynamodb.
	// subsTable := []*dynamodb.WriteRequest{}
	// subsTable[tableName] = &dynamodb.DeleteRe
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			tableName: wReqs,
		},
	}
	_, err = svc.BatchWriteItem(input)
	if err != nil {
		t.Fatal("Error while deleting items from DynamoDB table, ", err)
	}
}

func TestPOST_ErrSubscripcionExistsMsg(t *testing.T) {
	tearDown(t)
	// e := httpexpect.New(t, HTTPEndpoint)

	// // e.POST("/subscripcion")

	// e.POST("/subscripcion").
	// 	Expect().
	// 	Status(http.StatusCreated).
	// 	JSON().Object().ContainsKey("message").ValueEqual("message", messages.ErrSubscripcionNotFoundMsg)
}

// func TestGET(t *testing.T) {
// 	e := httpexpect.New(t, HTTPEndpoint)

// 	e.GET("/subscripcion").WithQuery("doc", 1234).
// 		Expect().
// 		Status(http.StatusOK).
// 		Body().Contains(messages.ErrSubscripcionNotFoundMsg)
// }

func TestGET_ErrSubscripcionNotFoundMsg(t *testing.T) {
	e := httpexpect.New(t, HTTPEndpoint)

	e.GET("/subscripcion").WithQuery("doc", 1234).
		Expect().
		Status(http.StatusBadRequest).
		JSON().Object().ContainsKey("message").ValueEqual("message", messages.ErrSubscripcionNotFoundMsg)
}
