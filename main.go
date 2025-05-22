package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Finished handling %s in %v", r.URL.Path, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()

	// Your routes
	r.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./helloworld.html")
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// Catch-all for unmatched routes
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("404 Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
	})

	// Wrap the router in the logger
	loggedRouter := loggingMiddleware(r)

	fmt.Println("Running Hello World server on Port 8080")
	http.ListenAndServe(":8080", loggedRouter)
}