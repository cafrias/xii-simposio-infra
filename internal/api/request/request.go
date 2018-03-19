package request

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/friasdesign/xii-simposio-infra/internal/api"
)

var _ api.Request = &Request{}

// FromAPIGatewayProxyRequest creates a new Request object from given APIGatewayProxyRequest.
func FromAPIGatewayProxyRequest(req events.APIGatewayProxyRequest) *Request {
	return &Request{
		natReqObj: &req,
	}
}

// Request represents a Request.
type Request struct {
	natReqObj *events.APIGatewayProxyRequest
}

// GetQuery gets a query parameter.
func (r *Request) GetQuery(p string) string {
	return r.natReqObj.QueryStringParameters[p]
}

// GetBody returns the body of the request.
func (r *Request) GetBody() string {
	return r.natReqObj.Body
}
