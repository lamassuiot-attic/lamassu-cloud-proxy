package endpoint

// io.lamassu.ca.create & io.lamassu.ca.import
type LamassuCaCreateEvent struct {
	CaName       string `json:"name"`
	CaCert       string `json:"cert"`
	SerialNumber string `json:"serial_number"`
}

// io.lamassu.ca.delete
type LamassuCaDeleteEvent struct {
	CaName string `json:"name"`
	Status string `json:"status"`
}

// io.lamassu.delete.cert
type LamassuDeleteCertEvent struct {
	CaName       string `json:"name"`
	Status       string `json:"status"`
	SerialNumber string `json:"serial_number"`
}
