package parser_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/simposio"
	"github.com/friasdesign/xii-simposio-infra/internal/simposio/parser"
)

func TestParse_ReturnsSubscripcionIfOK(t *testing.T) {
	subs := simposio.Subscripcion{
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

	in, err := json.Marshal(subs)
	if err != nil {
		t.Fatal("Failed setting up scenario: Cannot parse JSON.")
	}

	output, err := parser.Parse(string(in))
	if err != nil {
		t.Fatal("Returned error!")
	}

	if reflect.DeepEqual(output, &subs) == false {
		t.Fatal("Did not parsed value correctly!", output)
	}
}

func TestParse_ReturnsSubscripcionIfEmpty(t *testing.T) {
	subs := simposio.Subscripcion{}

	in, err := json.Marshal(subs)
	if err != nil {
		t.Fatal("Failed setting up scenario: Cannot parse JSON.")
	}

	output, err := parser.Parse(string(in))
	if err != nil {
		t.Fatal("Returned error!")
	}

	if reflect.DeepEqual(output, &subs) == false {
		t.Fatal("Did not parsed value correctly!", output)
	}
}

func TestParse_FailsWithInvalidJSON(t *testing.T) {
	in := "I am invalid JSON"

	output, err := parser.Parse(in)
	if output != nil {
		t.Fatal("Expected output to be nil!")
	}

	if err == nil {
		t.Fatal("Expected to return an error!")
	}
}

func TestParse_FailsWithInvalidTypes(t *testing.T) {
	var subs struct {
		Documento string `json:"documento"`
	}
	subs.Documento = "1234"

	in, err := json.Marshal(subs)

	output, err := parser.Parse(string(in))
	if output != nil {
		t.Fatal("Expected output to be nil!")
	}

	if err == nil {
		t.Fatal("Expected to return an error!")
	}
}
