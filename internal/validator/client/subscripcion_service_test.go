package client_test

import (
	"reflect"
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/validator"
)

func TestSubscripcionService_CreateSubscripcion(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	subs := validator.Subscripcion{
		Documento:        1234,
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

	// Create new Subscripcion.
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Retrieve Subscripcion and compare.
	other, err := s.Subscripcion(1234)
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(&subs, other) {
		t.Fatalf("unexpected Subscripcion: %#v", other)
	}
}
