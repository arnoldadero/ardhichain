package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, block := range blockchain {
		if block.ID == id {
			json.NewEncoder(w).Encode(block)
			return
		}
	}

	http.Error(w, "Block not found", http.StatusNotFound)
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
}

var blockchain []Block