package client_test

import (
	"reflect"
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
)

var subs simposio.Subscripcion

func setUp() (*Client, simposio.SubscripcionService) {
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

	c := MustOpenClient()
	s := c.SubscripcionService()
	return c, s
}

func TestSubscripcionService_CreateSubscripcion(t *testing.T) {
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

	if err := s.CreateSubscripcion(&subs); err != nil {
		t.Fatal(err)
	}
	if err := s.CreateSubscripcion(&subs); err != simposio.ErrSubscripcionExists {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionExists'")
	}
}

// Ensure duplicate dials are not allowed.
func TestSubscripcionService_CreateSubscripcion_ErrSubscripcionNotFound(t *testing.T) {
	c, s := setUp()
	defer c.Close()

	if _, err := s.Subscripcion(1234); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_UpdateSubscripcion(t *testing.T) {
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

	// Remove non existing Subscripcion
	if err := s.DeleteSubscripcion(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_Confirmar(t *testing.T) {
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

	// Confirmar non existing Subscripcion
	if err := s.Confirmar(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_Pendiente(t *testing.T) {
	c, s := setUp()
	defer c.Close()

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
	c, s := setUp()
	defer c.Close()

	// Pendiente non existing Subscripcion
	if err := s.Pendiente(subs.Documento); err != simposio.ErrSubscripcionNotFound {
		t.Fatal("Doesn't throw expected error 'ErrSubscripcionNotFound'")
	}
}

func TestSubscripcionService_Subscripciones(t *testing.T) {
	c, s := setUp()
	defer c.Close()

	// Create two 'Subscripcion' items.
	fix := map[int]simposio.Subscripcion{
		1: simposio.Subscripcion{
			Documento: 1,
		},
		2: simposio.Subscripcion{
			Documento: 2,
		},
	}
	for _, i := range fix {
		if err := s.CreateSubscripcion(&i); err != nil {
			t.Fatal(err)
		}
	}

	// Fetch all subscripcion items.
	res, err := s.Subscripciones()
	if err != nil {
		t.Fatal("Unexpected error, ", err)
	}
	for _, actual := range res {
		expected := fix[actual.Documento]
		if reflect.DeepEqual(*actual, expected) == false {
			t.Fatal("Expected and actual aren't the same.", actual, expected)
		}
	}
}

func TestSubscripcionService_SubscripcionesPendientes(t *testing.T) {
	c, s := setUp()
	defer c.Close()

	// Create two 'Subscripcion' items.
	fix := map[int]simposio.Subscripcion{
		1: simposio.Subscripcion{
			Documento:  1,
			Confirmado: true,
		},
		2: simposio.Subscripcion{
			Documento:  2,
			Confirmado: false,
		},
	}
	for _, i := range fix {
		if err := s.CreateSubscripcion(&i); err != nil {
			t.Fatal(err)
		}
	}

	// Fetch all subscripcion items.
	res, err := s.SubscripcionesPendientes()
	if err != nil {
		t.Fatal("Unexpected error, ", err)
	}
	if n := len(res); n != 1 {
		t.Fatal("Unexpected number of items, ", n)
	}
	if reflect.DeepEqual(*res[0], fix[2]) == false {
		t.Fatal("Expected and actual aren't the same.", res[0], fix[2])
	}
}

func TestSubscripcionService_SubscripcionesConfirmadas(t *testing.T) {
	c, s := setUp()
	defer c.Close()

	// Create two 'Subscripcion' items.
	fix := map[int]simposio.Subscripcion{
		1: simposio.Subscripcion{
			Documento:  1,
			Confirmado: true,
		},
		2: simposio.Subscripcion{
			Documento:  2,
			Confirmado: false,
		},
	}
	for _, i := range fix {
		if err := s.CreateSubscripcion(&i); err != nil {
			t.Fatal(err)
		}
	}

	// Fetch all subscripcion items.
	res, err := s.SubscripcionesConfirmadas()
	if err != nil {
		t.Fatal("Unexpected error, ", err)
	}
	if n := len(res); n != 1 {
		t.Fatal("Unexpected number of items, ", n)
	}
	if reflect.DeepEqual(*res[0], fix[1]) == false {
		t.Fatal("Expected and actual aren't the same.", res[0], fix[1])
	}
}
