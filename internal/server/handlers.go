package server

import (
	"encoding/xml"
	"net/http"

	"triple-s/internal/storage"
	"triple-s/internal/xmlfmt"
)

func (h *Handler) handleListBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := storage.ListBuckets(h.baseDir)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	var xmlBuckets []xmlfmt.Bucket
	for _, b := range buckets {
		xmlBuckets = append(xmlBuckets, xmlfmt.Bucket{
			Name:         b.Name,
			CreationDate: b.CreationDate,
		})
	}

	resp := xmlfmt.ListAllMyBucketsResult{
		Buckets: xmlfmt.Buckets{
			Bucket: xmlBuckets,
		},
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	_ = xml.NewEncoder(w).Encode(resp)
}

func (h *Handler) handleCreateBucket(w http.ResponseWriter, r *http.Request, bucket string) {
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "create bucket is not implemented yet: "+bucket)
}

func (h *Handler) handleDeleteBucket(w http.ResponseWriter, r *http.Request, bucket string) {
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "delete bucket is not implemented yet: "+bucket)
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
