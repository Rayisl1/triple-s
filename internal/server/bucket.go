package server

import (
	"encoding/xml"
	"net/http"
	"time"

	"triple-s/internal/storage"
	"triple-s/internal/validate"
	"triple-s/internal/xmlfmt"
)

type Handler struct {
	baseDir string
}

func NewHandler(baseDir string) *Handler {
	return &Handler{baseDir: baseDir}
}

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

func (h *Handler) handleCreateBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("BucketName")

	if err := validate.BucketName(bucket); err != nil {
		xmlfmt.WriteError(w, http.StatusBadRequest, "InvalidBucketName", err.Error())
		return
	}

	exists, err := storage.IsExistBucket(h.baseDir, bucket)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if exists {
		xmlfmt.WriteError(w, http.StatusConflict, "BucketAlreadyExists", "bucket already exists")
		return
	}

	if err := storage.CreateBucketDir(h.baseDir, bucket); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if err := storage.CreateObjectcsvinBucket(h.baseDir, bucket); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	if err := storage.AddBucket(h.baseDir, storage.BucketMeta{
		Name:         bucket,
		CreationDate: time.Now().UTC().Format(time.RFC3339),
	}); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleDeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("BucketName")

	exists, err := storage.IsExistBucket(h.baseDir, bucket)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if !exists {
		xmlfmt.WriteError(w, http.StatusNotFound, "NoSuchBucket", "bucket not found")
		return
	}

	empty, err := storage.IsBucketEmpty(h.baseDir, bucket)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if !empty {
		xmlfmt.WriteError(w, http.StatusConflict, "BucketNotEmpty", "bucket is not empty")
		return
	}

	if err := storage.RemoveBucketDir(h.baseDir, bucket); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	if err := storage.DeleteBucketFromCSV(h.baseDir, bucket); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
