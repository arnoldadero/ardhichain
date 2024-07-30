package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("searchHandler called") // Log when the handler is called

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for _, block := range blockchain {
		if block.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(block)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Block not found"})
}

func init() {
	// Initializing the blockchain with some dummy data
	blockchain = append(blockchain, Block{ID: 1, PreviousHash: "0", Hash: "hash1", Data: "Block 1 Data", Timestamp: "2023-07-01T00:00:00Z", Nonce: 0})
	blockchain = append(blockchain, Block{ID: 2, PreviousHash: "hash1", Hash: "hash2", Data: "Block 2 Data", Timestamp: "2023-07-02T00:00:00Z", Nonce: 0})
}

type Block struct {
	ID           int    `json:"id"`
	PreviousHash string `json:"previousHash"`
	Hash         string `json:"hash"`
	Data         string `json:"data"`
	Timestamp    string `json:"timestamp"`
	Nonce        int    `json:"nonce"`
	NationalID   string `json:"nationalId"`
}

var blockchain []Block
