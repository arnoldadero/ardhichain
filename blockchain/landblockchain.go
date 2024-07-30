package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// The Owner information
type Owner struct {
	Name, Address, NextOfKin string
	Id                       string
	Contact                  int
}

// Title deed information
type LandInfo struct {
	Owner                      Owner
	LandSize                   string
	LandLocation, Restrictions string
	Boundaries                 [][]int
}

// Mockland contract to simulate transactions
type MockLandContract struct {
	lands        map[string]LandInfo
	transactions map[string]bool
	mu           sync.Mutex
}

// Adding a transaction struct
type Transaction struct {
	ID          string
	From        string  // Sender or current owner
	To          string  // New owner
	LandID      string  // Identifier of the land being transacted
	Amount      float64 // Amount of land (if applicable)
	Description string  // Additional details about the transaction
}

// LandBlock information
type LandBlock struct {
	Data         map[string]interface{}
	TitleDeedId  string
	TimeStamp    time.Time
	Hash         string
	PreviousHash string
	Pow          int
	Transactions []Transaction
}

// Custom Blockchain and adding ethereum client
type Blockchain struct {
	genesisBlock LandBlock
	Chain        []LandBlock
	Difficulty   int
	mockContract *MockLandContract
}

// Generating the hash of a block
func (b LandBlock) CalculateHash() string {
	// convert the Data field to json for easy conversion to string
	dataBytes, err := json.Marshal(b.Data)
	if err != nil {
		fmt.Println("Error converting data to json/n", err)
		os.Exit(1)

	}

	blockData := b.PreviousHash + string(dataBytes) + b.TimeStamp.String() + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}

// Mining new block
func (b *LandBlock) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.CalculateHash()
	}
}

// Creating the genesis blockchain
func CreateBlockChain(difficulty int) Blockchain {
	mockContract := &MockLandContract{
		lands:        make(map[string]LandInfo),
		transactions: make(map[string]bool),
	}
	genesisBlock := LandBlock{
		Hash:      "0",
		TimeStamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]LandBlock{genesisBlock},
		difficulty,
		mockContract,
	}
}

// //Adding newblocks to the blockchain

func (b Blockchain) AddBlock(id string, landInfo LandInfo) {
	if _, exists := b.mockContract.GetLand(id); exists {
		fmt.Println("Land already registered. Use a different ID.")
		return
	}
	data := map[string]interface{}{
		"id":           id,
		"size":         landInfo.LandSize,
		"location":     landInfo.LandLocation,
		"restrictions": landInfo.Restrictions,
	}

	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := LandBlock{
		Data:         data,
		PreviousHash: lastBlock.Hash,
		TimeStamp:    time.Now(),
	}
	newBlock.Mine(b.Difficulty)
	b.mockContract.RegisterLand(id, landInfo)
	b.Chain = append(b.Chain, newBlock)

}

// Retrieving land information
func (b *Blockchain) GetLand(id string) (LandInfo, bool) {
	return b.mockContract.GetLand(id)
}

// Checking the validity of a blockchain
func (b Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		currentBlock := b.Chain[i+1]
		PreviousBlock := b.Chain[i]
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.Hash != PreviousBlock.Hash {
			return false
		}

	}
	return true
}

// Print the blockchain for debugging
func PrintBlockchain(chain []LandBlock) {
	for i, block := range chain {
		fmt.Printf("Block #%d\n", i)
		fmt.Printf("Timestamp: %s\n", block.TimeStamp)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Proof-of-Work: %d\n", block.Pow)
		fmt.Printf("Data: %v\n", block.Data)
		fmt.Println()
	}
}

// Main func
func main() {
	rand.Seed(uint64(time.Now().UnixNano()))
	blockchain := CreateBlockChain(2)

	// Register initial land
	initialLandID := fmt.Sprintf("land-%d", rand.Intn(1000))
	blockchain.AddBlock(initialLandID, LandInfo{
		LandSize:     "500 sqm",
		LandLocation: "Kanyakwar",
		Restrictions: "Yes",
	})

	blockchain.AddBlock("0001", LandInfo{
		LandSize:     "50 sqm",
		LandLocation: "Kisumu Central",
		Restrictions: "Yes",
	})

	// Example transaction

	transaction := Transaction{
		ID:          "001",
		From:        "Bob",
		To:          "Mary",
		LandID:      initialLandID,
		Amount:      1000.0,
		Description: "Transfer of land ownership",
	}
	transaction_2 := Transaction{
		ID:          "002",
		From:        "Bob",
		To:          "Barrack",
		LandID:      initialLandID,
		Amount:      900,
		Description: "Transfer of ownership",
	}

	// Add the transaction to the blockchain
	blockchain.AddTransaction(transaction)
	blockchain.AddTransaction(transaction_2)
	// Validate the blockchain
	if blockchain.IsValid() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
	PrintBlockchain(blockchain.Chain)
}
