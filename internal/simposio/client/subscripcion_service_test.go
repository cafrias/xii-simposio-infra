package client_test

import (
	"reflect"
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

var subs simposio.Subscripcion

func setUp() {
	subs = simposio.Subscripcion{
		Documento:        1234,
		Apellido:         "Frias",
		Nombre:           "Carlos",
		Email:            "carlos.a.frias@gmail.com",
		Direccion:        "Rio Fuego 3490",
		Zip:              9420,
		Localidad:        "Rio Grande",
		Pais:             "Argentina",
		ArancelCategoria: "general",
		ArancelPago:      "Con cheque",
		Acompanantes:     0,
	}
}

func TestSubscripcionService_CreateSubscripcion(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

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

// Ensure duplicate dials are not allowed.
func TestSubscripcionService_CreateSubscripcion_ErrSubscripcionExists(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}
	if err := s.CreateSubscripcion(&subs); err != simposio.ErrSubscripcionExists {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionExists'")
	}
}

// Ensure duplicate dials are not allowed.
func TestSubscripcionService_CreateSubscripcion_ErrSubscripcionNotFound(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	if _, err := s.Subscripcion(1234); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}
