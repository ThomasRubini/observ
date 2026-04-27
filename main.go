package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Starting observ server on port " + port)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from observ!\n")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}