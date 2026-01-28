package server

import (
	"net/http"
	"strings"
)

type Handler struct {
	baseDir string
}

func NewHandler(baseDir string) *Handler {
	return &Handler{baseDir: baseDir}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")

	if path == "" {
		if r.Method == http.MethodGet {
			h.handleListBuckets(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(path, "/")
	bucket := parts[0]

	if len(parts) == 1 {
		switch r.Method {
		case http.MethodPut:
			h.handleCreateBucket(w, r, bucket)
			return
		case http.MethodDelete:
			h.handleDeleteBucket(w, r, bucket)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	objectKey := strings.Join(parts[1:], "/")
	switch r.Method {
	case http.MethodPut:
		h.handlePutObject(w, r, bucket, objectKey)
		return
	case http.MethodGet:
		h.handleGetObject(w, r, bucket, objectKey)
		return
	case http.MethodDelete:
		h.handleDeleteObject(w, r, bucket, objectKey)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
