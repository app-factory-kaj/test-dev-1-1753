package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type GreetingResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := GreetingResponse{Message: "Hello, World!"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("error writing health response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)

	addr := ":9090"
	log.Printf("hello-api listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
