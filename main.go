package main

import (
	"log"
	"net/http"
	"strconv"

	"image/png"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		width, errw := strconv.ParseInt(req.URL.Query().Get("width"), 10, 32)
		if errw != nil {
			width = 400
		}
		height, errh := strconv.ParseInt(req.URL.Query().Get("height"), 10, 32)
		if errh != nil {
			height = 200
		}
		maxIterations, errm := strconv.ParseInt(req.URL.Query().Get("max"), 10, 32)
		if errm != nil {
			maxIterations = 255
		}
		zoom, errz := strconv.ParseFloat(req.URL.Query().Get("zoom"), 32)
		if errz != nil {
			zoom = 1.0
		}

		png.Encode(res, createImage(int(width), int(height), int(maxIterations), float32(zoom)))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
