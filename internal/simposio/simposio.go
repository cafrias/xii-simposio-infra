package simposio

import "github.com/friasdesign/xii-simposio-infra/internal/dynamodb"

// Subscripcion represents a single subscription to the event.
type Subscripcion struct {
	Documento           int     `json:"documento" validate:"required"`
	Apellido            string  `json:"apellido" validate:"required"`
	Nombre              string  `json:"nombre" validate:"required"`
	Telefono            int     `json:"telefono,omitempty"`
	Celular             int     `json:"celular,omitempty"`
	Fax                 int     `json:"fax,omitempty"`
	Email               string  `json:"email" validate:"required,email"`
	Direccion           string  `json:"direccion" validate:"required"`
	Zip                 int     `json:"zip" validate:"required,gte=1000,lt=10000"`
	Localidad           string  `json:"localidad" validate:"required"`
	Pais                string  `json:"pais" validate:"required"`
	ArancelAdicional    float64 `json:"arancel_adicional,omitempty" validate:"gte=0"`
	ArancelCategoria    int     `json:"arancel_categoria" validate:"required"`
	ArancelPago         string  `json:"arancel_pago" validate:"required,gte=0"`
	PonenciaPresenta    bool    `json:"ponencia_presenta" validate:"required"`
	PonenciaTitulo      string  `json:"ponencia_titulo,omitempty"`
	PonenciaArea        string  `json:"ponencia_area,omitempty"`
	PonenciaCoautores   string  `json:"ponencia_coautores,omitempty"`
	PonenciaInstitucion string  `json:"ponencia_institucion,omitempty"`
	Acompanantes        int     `json:"acompanantes" validate:"required,gte=0"`
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
}
