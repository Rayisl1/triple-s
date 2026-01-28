package xmlfmt

import (
	"encoding/xml"
	"net/http"
)

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}

func WriteError(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)

	_ = xml.NewEncoder(w).Encode(ErrorResponse{
		Code:    code,
		Message: message,
	})
}
