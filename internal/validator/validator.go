package validator

import "github.com/friasdesign/xii-simposio-infra/internal/dynamodb"

// Subscripcion represents a single subscription to the event.
type Subscripcion struct {
	Documento           int     `json:"documento"`
	Apellido            string  `json:"apellido"`
	Nombre              string  `json:"nombre"`
	Telefono            int     `json:"telefono,omitempty"`
	Celular             int     `json:"celular,omitempty"`
	Fax                 int     `json:"fax,omitempty"`
	Email               string  `json:"email"`
	Direccion           string  `json:"direccion"`
	Zip                 int     `json:"zip"`
	Localidad           string  `json:"localidad"`
	Pais                string  `json:"pais"`
	ArancelAdicional    float64 `json:"arancel_adicional,omitempty"`
	ArancelCategoria    int     `json:"arancel_categoria"`
	ArancelPago         float64 `json:"arancel_pago"`
	PonenciaPresenta    bool    `json:"ponencia_presenta"`
	PonenciaTitulo      string  `json:"ponencia_titulo,omitempty"`
	PonenciaArea        string  `json:"ponencia_area,omitempty"`
	PonenciaCoautores   string  `json:"ponencia_coautores,omitempty"`
	PonenciaInstitucion string  `json:"ponencia_institucion,omitempty"`
	Acompanantes        int     `json:"acompanantes"`
}

// Client creates a connection to the services.
type Client interface {
	dynamodb.DynamoDB
	SubscripcionService() SubscripcionService
}

// SubscripcionService represents a service for managing Subscripcion.
type SubscripcionService interface {
	Subscripcion(doc int) (*Subscripcion, error)
	CreateSubscripcion(subs Subscripcion) error
}
