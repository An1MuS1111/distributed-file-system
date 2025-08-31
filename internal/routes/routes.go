package routes

import (
	"distributed-file-system/internal/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /upload", handlers.UploadHandler)
	return mux
}
