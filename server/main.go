package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/fernand-o/fb2img"
)

func serverHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "url param not found", http.StatusBadRequest)
		return
	}

	img, err := fb2img.CreateImage(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(img)))
	w.Write(img)
}

func main() {
	http.HandleFunc("/fb2img", serverHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
