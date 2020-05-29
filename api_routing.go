package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index Route
//
// Params according to HandleFunc pattern: func(ResponseWriter, *Request))
func index(indexResponse http.ResponseWriter, indexRequest *http.Request) {
	fmt.Fprintf(indexResponse, "Check-CPF API by Heitor Peralles!")
}

// Routing service start function
func startRouting() {

	log.Print("Initializing Check-CPF API")

	// Init router
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api/v1/verify", verify).Methods("POST")

	// Start server
	port := ":8000"
	log.Print("Starting to listen on port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
