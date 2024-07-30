package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println("newOwnerHandler called")

	var owner Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Assuming you add this owner to the blockchain or save it in your data store
	// Here, you can also create a new block in the blockchain with these details

	// For demonstration, we'll just print the owner details
	fmt.Printf("New Owner: %+v\n", owner)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "New owner created successfully"})
}
