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
		Confirmado:       false,
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

func TestSubscripcionService_UpdateSubscripcion(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Create new Subscripcion.
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Update Subscripcion
	subs.Email = "mynewemail@email.com"
	if err := s.UpdateSubscripcion(&subs); err != nil {
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

func TestSubscripcionService_UpdateSubscripcion_ErrSubscripcionNotFound(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Create new Subscripcion.
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Update Subscripcion
	subs.Documento = subs.Documento + 1
	subs.Email = "mynewemail@email.com"
	if err := s.UpdateSubscripcion(&subs); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_DeleteSubscripcion(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Create new Subscripcion.
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Remove Subscripcion
	if err := s.DeleteSubscripcion(subs.Documento); err != nil {
		t.Fatal(err)
	}

	// Try retrieve Subscripcion.
	if _, err := s.Subscripcion(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_DeleteSubscripcion_ErrSubscripcionNotFound(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Remove non existing Subscripcion
	if err := s.DeleteSubscripcion(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_Confirmar(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Create new Subscripcion.
	subs.Confirmado = false
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Confirmar
	if err := s.Confirmar(subs.Documento); err != nil {
		t.Fatal(err)
	}

	// Try retrieve Subscripcion.
	confSubs, err := s.Subscripcion(subs.Documento)
	if err != nil {
		t.Fatal("Unexpected error, ", err)
	}

	if confSubs.Confirmado != true {
		t.Fatal("Didn't set 'confirmado' to true.")
	}
}

func TestSubscripcionService_Confirmar_ErrSubscripcionNotFound(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Confirmar non existing Subscripcion
	if err := s.Confirmar(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_Pendiente(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Create new Subscripcion.
	subs.Confirmado = true
	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}

	// Pendiente
	if err := s.Pendiente(subs.Documento); err != nil {
		t.Fatal(err)
	}

	// Try retrieve Subscripcion.
	confSubs, err := s.Subscripcion(subs.Documento)
	if err != nil {
		t.Fatal("Unexpected error, ", err)
	}

	if confSubs.Confirmado != false {
		t.Fatal("Didn't set 'confirmado' to false.")
	}
}

func TestSubscripcionService_Pendiente_ErrSubscripcionNotFound(t *testing.T) {
	setUp()
	c := MustOpenClient()
	defer c.Close()
	s := c.SubscripcionService()

	// Pendiente non existing Subscripcion
	if err := s.Pendiente(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}
