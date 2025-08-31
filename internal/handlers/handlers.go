package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// set the key in the front end as "media"
	file, fileHeader, err := r.FormFile("media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	t := time.Now()

	// in this part we need to store the file to the disk
	stored, err := os.Create(fmt.Sprintf("%v_%v", t.String(), fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	println(stored.Name())

	w.Write([]byte("uploaded"))
}
