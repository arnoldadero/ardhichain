package main

// import (
// 	"crypto/sha256"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // The Owner information
// type Owner struct {
// 	Name, Address, NextOfKin string
// 	Id                       string
// 	Contact                  int
// }

// // Title deed information
// type LandInfo struct {
// 	LandSize                   float64
// 	LandLocation, Restrictions string
// 	Boundaries                 [][]int
// }

// // LandBlock information
// type LandBlock struct {
// 	Data         map[string]interface{}
// 	TitleDeedId  string
// 	TimeStamp    time.Time
// 	Hash         string
// 	PreviousHash string
// 	Pow          int
// }

// // Custom Blockchain
// type Blockchain struct {
// 	genesisBlock LandBlock
// 	Chain        []LandBlock
// 	Difficulty   int
// }

// // Generating the hash of a block
// func (b LandBlock) CalculateHash() string {
// 	// convert the Data field to json for easy conversion to string
// 	dataBytes, err := json.Marshal(b.Data)
// 	if err != nil {
// 		fmt.Println("Error converting data to json/n", err)
// 		os.Exit(1)

// 	}

// 	blockData := b.PreviousHash + string(dataBytes) + b.TimeStamp.String() + strconv.Itoa(b.Pow)
// 	blockHash := sha256.Sum256([]byte(blockData))

// 	return fmt.Sprintf("%x", blockHash)
// }

// // Mining new block
// func (b *LandBlock) Mine(difficulty int) {
// 	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
// 		b.Pow++
// 		b.Hash = b.CalculateHash()
// 	}
// }

// // Creating the genesis blockchain
// func CreateBlockChain(difficulty int) Blockchain {
// 	genesisBlock := LandBlock{
// 		Hash:      "0",
// 		TimeStamp: time.Now(),
// 	}
// 	return Blockchain{
// 		genesisBlock,
// 		[]LandBlock{genesisBlock},
// 		difficulty,
// 	}
// }

// //Adding newblocks to the blockchain

// func (b Blockchain) AddBlock(to, from string, amount float64) {
// 	landBlockData := map[string]interface{}{
// 		"from" : from,
// 		"to" : to,
// 		"amount" : amount,
// 	}
// 	lastBlock := b.Chain[len(b.Chain)-1]
// 	newBlock := LandBlock{
// 		Data : landBlockData,
// 		PreviousHash : lastBlock.Hash,
// 		TimeStamp : time.Now(),
// 	}
// 	newBlock.Mine(b.Difficulty)
// 	b.Chain = append(b.Chain, newBlock)
// }

// //Checking the validity of a blockchain
// func (b Blockchain) IsValid()bool{
// 	for i := range b.Chain[1:]{
// 		currentBlock := b.Chain[i +1]
// 		PreviousBlock := b.Chain[i]
// 		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.Hash != PreviousBlock.Hash{
// 			return false
// 		}

// 	}
// 	return true
// }

// //Main func
// func main(){
// 	//Create a Blockchain instance with a mining difficulty of 2
// 	blockchain := CreateBlockChain(2)

// 	//Record Transactions on the blockchain
// 	blockchain.AddBlock("Elijah", "Ben", 1)
// 	blockchain.AddBlock("Nick", "Stee", 5)
// 	fmt.Println(blockchain.IsValid())
// }
