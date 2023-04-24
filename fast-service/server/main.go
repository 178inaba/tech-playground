package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

const maxSize = 26214400 // 25MB

func main() {
	ctx := context.Background()
	idleConnsClosed := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc("/download", downloadHandler())
	mux.HandleFunc("/upload", uploadHandler())
	srv := http.Server{
		Handler: mux,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v.", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v.", err)
	}

	<-idleConnsClosed
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
