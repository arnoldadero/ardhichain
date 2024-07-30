package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifyHandler called") // Log when the handler is called

	var req struct {
		NationalID string `json:"nationalId"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	for _, block := range blockchain {
		if block.NationalID == req.NationalID {
			w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct {
				Verified bool `json:"verified"`
			}{Verified: true})
			return
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(struct {
		Verified bool `json:"verified"`
	}{Verified: false})
}
