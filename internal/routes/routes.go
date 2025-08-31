package routes

import (
	"distributed-file-system/internal/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /upload", handlers.UploadHandler)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello there!"))
	})
	return mux
}
