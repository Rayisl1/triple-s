package server

import (
	"fmt"
	"net/http"
	"triple-s/internal/config"
)

func Run(cfg config.Config) error {
	mux := http.NewServeMux()

	h := NewHandler(cfg.Dir)

	mux.HandleFunc("GET /", h.handleListBuckets)
	mux.HandleFunc("PUT /{BucketName}", h.handleCreateBucket)
	mux.HandleFunc("DELETE /{BucketName}", h.handleDeleteBucket)

	fmt.Printf("Server started on http://localhost:%d\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux)
}
