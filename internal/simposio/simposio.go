package simposio

import "github.com/friasdesign/xii-simposio-infra/internal/dynamodb"

// Aranceles contains a map with the values of each arancel.
var Aranceles = map[string]float64{
	"estudiante_untdf":   0,
	"estudiante_otro":    350,
	"docente_untdf":      1600,
	"matriculado_cpcetf": 1000,
	"general":            2700,
}

// Subscripcion represents a single subscription to the event.
type Subscripcion struct {
	Documento           int     `json:"documento" validate:"required,gt=1000000"`
	Apellido            string  `json:"apellido" validate:"required"`
	Nombre              string  `json:"nombre" validate:"required"`
	Telefono            int     `json:"telefono,omitempty" validate:"omitempty,gt=0"`
	Celular             int     `json:"celular,omitempty" validate:"omitempty,gt=0"`
	Fax                 int     `json:"fax,omitempty" validate:"omitempty,gt=0"`
	Email               string  `json:"email" validate:"required,email"`
	Direccion           string  `json:"direccion" validate:"required"`
	Zip                 int     `json:"zip" validate:"required,gte=1000,lt=10000"`
	Localidad           string  `json:"localidad" validate:"required"`
	Pais                string  `json:"pais" validate:"required"`
	ArancelAdicional    float64 `json:"arancel_adicional,omitempty" validate:"gte=0"`
	ArancelCategoria    string  `json:"arancel_categoria" validate:"required,arancel"`
	ArancelPago         string  `json:"arancel_pago" validate:"required,gte=0"`
	PonenciaPresenta    bool    `json:"ponencia_presenta" validate:"required"`
	PonenciaTitulo      string  `json:"ponencia_titulo,omitempty"`
	PonenciaArea        string  `json:"ponencia_area,omitempty"`
	PonenciaCoautores   string  `json:"ponencia_coautores,omitempty"`
	PonenciaInstitucion string  `json:"ponencia_institucion,omitempty"`
	Acompanantes        int     `json:"acompanantes" validate:"required,gte=0"`
	Confirmado          bool    `json:"confirmado"`
}

// GetArancelBase returns the base values for arancel.
func (s *Subscripcion) GetArancelBase() float64 {
	return Aranceles[s.ArancelCategoria]
}

// GetArancelTotal returns the total value for the arancel.
func (s *Subscripcion) GetArancelTotal() float64 {
	return s.GetArancelBase() + s.ArancelAdicional
}

// Client creates a connection to the services.
type Client interface {
	dynamodb.DynamoDB
	SubscripcionService() SubscripcionService
}

// SubscripcionService represents a service for managing Subscripcion.
type SubscripcionService interface {
	Subscripcion(doc int) (*Subscripcion, error)
	CreateSubscripcion(subs *Subscripcion) error
	UpdateSubscripcion(subs *Subscripcion) error
	DeleteSubscripcion(doc int) error
}
