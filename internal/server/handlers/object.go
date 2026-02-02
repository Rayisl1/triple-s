package server

import (
	"net/http"

	"triple-s/internal/xmlfmt"
)

type Handler struct {
	baseDir string
}

func NewHandler(baseDir string) *Handler {
	return &Handler{baseDir: baseDir}
}
func (h *Handler) handlePutObject(w http.ResponseWriter, r *http.Request, bucket, objectKey string) {
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "put object is not implemented yet: "+bucket+"/"+objectKey)
}

func (h *Handler) handleGetObject(w http.ResponseWriter, r *http.Request, bucket, objectKey string) {
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "get object is not implemented yet: "+bucket+"/"+objectKey)
}

func (h *Handler) handleDeleteObject(w http.ResponseWriter, r *http.Request, bucket, objectKey string) {
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "delete object is not implemented yet: "+bucket+"/"+objectKey)
}
