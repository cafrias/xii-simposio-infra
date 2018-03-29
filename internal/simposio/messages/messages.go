package messages

// Response messages.
const (
	ErrValidationMsg           = "Error de validación en Subscripción."
	ErrRequestMsg              = "Petición mal formada, contacte con soporte."
	ErrInternalMsg             = "Error Interno, contacte con soporte."
	ErrSubscripcionExistsMsg   = "Ya se registró un usuario con ese Documento."
	ErrSubscripcionNotFoundMsg = "No se encuentra subscripción con ese Documento."
	ErrQueryParamDocInvalidMsg = "Debe indicar un documento valido para buscar."
	SucSavingSubscipcionMsg    = "Subscripción registrada con éxito."
	SucFetchingSubscripcionMsg = "Subscripción encontrada con éxito."
	SucFetchingListMsg         = "Subscripciones encontras con éxito."
	SucDeletingSubscripcionMsg = "Subscripción eliminada con éxito."
)

// Log messages.
const (
	ErrUUIDLog                 = "Couldn't generate UUID."
	ErrRequestLog              = "Request body is invalid!"
	ErrValidationLog           = "Validation Error\n"
	ErrParsingBodyToJSON       = "Error parsing body to JSON"
	ErrDynamoDBConnectionLog   = "Error while trying to open connection to DynamoDB"
	ErrSavingSubscripcionLog   = "Error while trying to write Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrFetchingSubscripcionLog = "Error while trying to fetch Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrDynamoDBFetchingLog     = "Error while trying to fetch data from DynamoDB, '%v'\n"
	ErrDeletingSubscripcionLog = "Error while trying to delete Subscripcion with 'Documento' %v to DynamoDB\n"
	ErrSubscripcionExistsLog   = "Subscripcion with 'Documento' %v already exists!\n"
	ErrSubscripcionNotFoundLog = "Subscripcion with 'Documento' %v not found!\n"
	ErrUnexpectedHTTPMethodLog = "Unexpected HTTP method '%s'\n"
	ErrQueryParamDocInvalidLog = "Invalid 'doc' query param '%v'\n"
	ErrQueryParamInvalidLog    = "Invalid query param '%s' with value '%v'\n"
	SucSavingSubscripcionLog   = "Subscripcion with 'Documento' %v successfully saved\n"
	SucFetchingSubscripcionLog = "Subscripcion with 'Documento' %v successfully fetched\n"
	SucDeletingSubscripcionLog = "Subscripcion with 'Documento' %v successfully deleted\n"
	SucFetchingListLog         = "Subscripcion list fetched successfully."
)
