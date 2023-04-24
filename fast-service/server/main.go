package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const maxSize = 26214400 // 25MB

func main() {
}

func downloadHandler() http.HandlerFunc {
	src := rand.NewSource(0)
	return func(w http.ResponseWriter, r *http.Request) {
		max, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil {
			max = maxSize
		}

		if _, err := io.CopyN(w, rand.New(src), int64(max)); err != nil {
			log.Printf("Failed to write random data: %v.", err)
			return
		}
	}
}

func uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentLength := r.ContentLength
		if contentLength > maxSize {
			contentLength = maxSize
		}

		if _, err := io.CopyN(io.Discard, r.Body, contentLength); err != nil {
			log.Printf("Failed to write body: %v.", err)
			return
		}
	}
}
