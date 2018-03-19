package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio/client"

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

func setUp(t *testing.T) *client.Client {
	os.Setenv("TABLE_NAME", tableName)
	c := client.NewClient()
	err := c.Open()
	if err != nil {
		t.Fatal("Error while openning connection to DB, ", err)
	}
	return c
}

func ReadJSONFixture(path string, structure interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, structure)
	if err != nil {
		return err
	}

	return nil
}

func TestGET(t *testing.T) {
	c := setUp(t)
	defer tearDown(t)

	var subs simposio.Subscripcion
	err := ReadJSONFixture("testdata/OK.json", &subs)
	if err != nil {
		t.Fatal("Error while reading fixture, ", err)
	}

	// Populate DB with fixture.
	err = c.SubscripcionService().CreateSubscripcion(&subs)
	if err != nil {
		t.Fatal("Error while writing to DynamoDB, ", err)
	}

	e := httpexpect.New(t, HTTPEndpoint)

	e.GET("/subscripcion").WithQuery("doc", subs.Documento).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("payload").ValueEqual("payload", subs)
}

func TestGET_ErrSubscripcionNotFoundMsg(t *testing.T) {
	e := httpexpect.New(t, HTTPEndpoint)

	e.GET("/subscripcion").WithQuery("doc", 1234).
		Expect().
		Status(http.StatusBadRequest).
		JSON().Object().
		ContainsKey("message").ValueEqual("message", messages.ErrSubscripcionNotFoundMsg)
}
