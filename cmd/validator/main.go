package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/friasdesign/xii-simposio-infra/internal/validator"
	"github.com/friasdesign/xii-simposio-infra/internal/validator/db"
)

// Response represents a response for AWS Lambda.
type Response struct {
	Message string `json:"message"`
}

// Handler is used by AWS Lambda to handle request.
func Handler() (Response, error) {
	var c validator.Client = db.NewClient()
	err := c.Open()
	if err != nil {
		fmt.Println(err.Error())
		return Response{
			Message: "Something went wrong while opening database connection!",
		}, err
	}

	subsService := c.SubscripcionService()

	subs := validator.Subscripcion{
		Documento:        34186552,
		Apellido:         "Frias",
		Nombre:           "Carlos",
		Email:            "carlos.a.frias@gmail.com",
		Direccion:        "Rio Fuego 3490",
		Zip:              9420,
		Localidad:        "Rio Grande",
		Pais:             "Argentina",
		ArancelCategoria: 1,
		ArancelPago:      1245.1234,
		Acompanantes:     0,
	}

	err = subsService.CreateSubscripcion(subs)
	if err != nil {
		return Response{
			Message: "Something went wrong while writing to database!",
		}, err
	}

	return Response{
		Message: "Done! Check in the DB",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
