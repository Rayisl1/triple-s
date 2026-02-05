package server

import (
	"encoding/xml"
	"net/http"
	"strconv"
	"time"

	"triple-s/internal/storage"
	"triple-s/internal/validate"
	"triple-s/internal/xmlfmt"
)

func (h *Handler) handlePutObject(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	if err := validate.ObjectName(objectKey); err != nil {
		xmlfmt.WriteError(w, http.StatusBadRequest, "InvalidObjectName", err.Error())
		return
	}

	exists, err := storage.IsExistObject(h.baseDir, bucket, objectKey)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if exists {
		xmlfmt.WriteError(w, http.StatusConflict, "ObjectAlreadyExists", "bucket already exists")
		return
	}

	if err := storage.CreateObjectFile(h.baseDir, bucket, objectKey); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stram"
	}
	contentLength := r.ContentLength
	if contentLength < 0 {
		contentLengthStr := r.Header.Get("Content-Length")
		if contentLengthStr != "" {
			if cl, err := strconv.ParseInt(contentLengthStr, 10, 64); err == nil {
				contentLength = cl
			}
		}
	}

	if err := storage.AddObject(h.baseDir, bucket, objectKey, storage.ObjectMeta{
		Name:         objectKey,
		Size:         contentLength,
		ContentType:  contentType,
		CreationDate: time.Now().UTC().Format(time.RFC3339),
	}); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleGetObject(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")
	objects, err := storage.ListObjects(h.baseDir, bucket)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	var xmlObjects []xmlfmt.Object
	for _, o := range objects {
		xmlObjects = append(xmlObjects, xmlfmt.Object{
			Name:         o.Name,
			Size:         o.Size,
			ContentType:  o.ContentType,
			CreationDate: o.CreationDate,
		})
	}

	resp := xmlfmt.ListAllMyObjectsResult{
		Objects: xmlfmt.Objects{
			Object: xmlObjects,
		},
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	_ = xml.NewEncoder(w).Encode(resp)
	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "get object is not implemented yet: "+bucket+"/"+objectKey)
}

func (h *Handler) handleDeleteObject(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("ObjectName")
	objectKey := r.PathValue("ObjectKey")

	exists, err := storage.IsExistObject(h.baseDir, bucket, objectKey)
	if err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	if !exists {
		xmlfmt.WriteError(w, http.StatusNotFound, "NoSuchObject", "object not found")
		return
	}

	if err := storage.RemoveObjectFile(h.baseDir, bucket, objectKey); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	if err := storage.DeleteObjectFromCSV(h.baseDir, bucket, objectKey); err != nil {
		xmlfmt.WriteError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}

	xmlfmt.WriteError(w, http.StatusNotImplemented, "NotImplemented", "delete object is not implemented yet: "+bucket+"/"+objectKey)
}
