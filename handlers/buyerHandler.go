package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

type Owner struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	IDNumber          string `json:"idNumber"`
	NextOfKinName     string `json:"nextOfKinName"`
	NextOfKinIdNumber string `json:"nextOfKinIdNumber"`
	Amount            int    `json:"amount"`
	Timestamp         string `json:"timestamp"`
}

// newOwnerHandler is a handler function that creates a new owner in the blockchain
func NewOwnerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("newOwnerHandler called")

	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	log.Println("Request Body:", string(body))

	// Unmarshal the JSON into the Owner struct
	var owner Owner
	if len(body) == 0 {
		tmpl, err := template.ParseFiles("templates/buyer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, "")
	}
	err = json.Unmarshal(body, &owner)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Log the received owner details
	log.Printf("Received new owner: %+v\n", owner)

	// For demonstration, we'll just print the owner details
	fmt.Printf("New Owner: %+v\n", owner)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "New owner created successfully"})
}
