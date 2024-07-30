package main

import (
	"fmt"
	"time"
)

// New MockLandContract initilizes new smart contract.
func NewMockLandContract() *MockLandContract {
	return &MockLandContract{
		lands: make(map[string]LandInfo),
	}
}

// Register land in the smart contract
func (m *MockLandContract) RegisterLand(id string, info LandInfo) {
	if _, exists := m.lands[id]; exists{
		fmt.Println("Land already Registered")
	}
	m.lands[id] = info
}

// Transfer ownership represents transferring ownership of land
func (m *MockLandContract) TransferLand(id string, newOwner string) {
	if land, exists := m.lands[id]; exists {
		fmt.Printf("Ownership of land %s transferred to %v \n", id, newOwner)
		//Simulating a transfer
		land.Owner.Name = newOwner
		m.lands[id] = land
	} else {
		fmt.Println("Land not found")
		return
	}

}

// Function GetLand simulates getting land from the smart contract.

func (m *MockLandContract) GetLand(id string) (LandInfo, bool) {
	info, exists := m.lands[id]
	return info, exists
}

// AddTransaction adds a transaction to the blockchain's latest block
func (b *Blockchain) AddTransaction(tx Transaction) {
	// Get the latest block
	lastBlock := b.Chain[len(b.Chain)-1]

	// Create a new block for the transaction
	newBlock := LandBlock{
		Data: map[string]interface{}{
			"transaction": tx,
		},
		PreviousHash: lastBlock.Hash,
		TimeStamp:    time.Now(),
	}

	// Mine the new block
	newBlock.Hash = newBlock.CalculateHash()
	newBlock.Mine(b.Difficulty)

	// Add the new block to the blockchain
	b.Chain = append(b.Chain, newBlock)
}

// UpdateLand simulates updating an existing land record
func (m *MockLandContract) UpdateLand(id string, updates LandInfo) {
	if land, exists := m.lands[id]; exists {
		if updates.LandSize != "" {
			land.LandSize = updates.LandSize
		}
		if updates.LandLocation != "" {
			land.LandLocation = updates.LandLocation
		}
		if updates.Restrictions != "" {
			land.Restrictions = updates.Restrictions
		}
		m.lands[id] = land
		fmt.Printf("Land %s updated.\n", id)
	} else {
		fmt.Println("Land not found!")
	}
}