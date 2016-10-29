package main

import (
	"github.com/fernand-o/fb2img/fb2img"
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		// log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		// log.Println("unable to write image.")
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")

	imgpath, htmlpath := fb2img.CreateImage(url)
	defer func() {
		os.Remove(imgpath)
		os.Remove(htmlpath)
	}()

	imgbuffer, _ := os.Open(imgpath)
	defer imgbuffer.Close()

	img_as_png, _, _ := image.Decode(imgbuffer)

	writeImage(w, &img_as_png)
}

func main() {
	http.HandleFunc("/fb2img", serverHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
