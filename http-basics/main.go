package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler called, method:", r.Method)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hey from go!"}`))
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("userHandler called, method:", r.Method)
	// 1. Allow only POST method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"only POST method is allowed"}`))
		return
	}

	// 2. Parse JSON body into User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid JSON"}`))
		return
	}

	// 3. Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 4. Create response body
	response := map[string]interface{}{
		"status": "recieved",
		"user":   user,
	}

	// 5. Encode response as JSON
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Server started on port 8080")

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)

	http.ListenAndServe(":8080", nil)
}
