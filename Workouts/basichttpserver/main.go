package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", greet)
	http.HandleFunc("/health", healthstatus)

	// This message confirms the code has finished compiling and is now running
	log.Println("Server starting on :8080...")
	
	// log.Fatal will print the error and stop the program if the port is busy
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher!"))
}

func healthstatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"UP"}`))
}