package server

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"triple-s/internal/storage"
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

func (h *Handler) handleCreateBucket(baseDir string, w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")

	exist, err := storage.IsExistBucket(filepath.Join(baseDir, "buckets.csv"))
	if err != nil {
		//TODO
	}
	if exist {
		return
	}
	err1 := os.Mkdir(bucketName, 0755)

	if err1 != nil {
		log.Fatal("Error creating directory:", err1)
	}
	fmt.Println("Directory created successfully!")
	path := filepath.Join(baseDir, "buckets.csv")
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "create bucket is not implemented yet: "+bucketName)
}

func (h *Handler) handleDeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "delete bucket is not implemented yet: "+bucketName)
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
