package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// this is struct Block, which is the content of the block on the blochchain
type Block struct {
	data         map[string]interface{} 
	hash         string 
	previousHash string 
	timestamp    time.Time  
	pow          int 
}

// this is the struct Blockchain, which is the blockchain itself.
type Blockchain struct {
	genesisBlock Block 
	chain        []Block 
	difficulty   int 
}

// this is a method of struct Block, which calculates the hash of the block. (b.hash)
func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow) 
	blockHash := sha256.Sum256([]byte(blockData)) 
	return fmt.Sprintf("%x", blockHash)
}

// this is a method of struct Block, which mines the block. 
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++ 
		b.hash = b.calculateHash() 
	}
}

// this is a method of struct Blockchain, which creates this blockchain.
func CreateBlockchain(difficulty int) Blockchain {

	genesisBlock := Block{ 
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// this is a method of struct Blockchain, which adds a block to the blockchain.
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

// this is a method of struct Blockchain, which checks if the blockchain is valid.
func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (b Blockchain) getData() {
	for {
		var id int
		fmt.Print("Input block id: ")
		fmt.Scan(&id)

		fmt.Println(b.chain[id].data)
	}
}


func main() {
    // create a new blockchain instance with a mining difficulty of 2
    blockchain := CreateBlockchain(3)

    // record transactions on the blockchain for Gbolahan
    blockchain.addBlock("Gbolahan", "Galad", 5)

    // check if the blockchain is valid; expecting true
    fmt.Printf("%v\n\n",blockchain.isValid())

	// print all the blocks on the blockchain
	for _, block := range blockchain.chain {	
		fmt.Println(block)
	}
	fmt.Println("")

	// print block.data for Gbolahan
	fmt.Printf("Gbolahan's Block data: %v\n", blockchain.chain[1].data)

	fmt.Printf("Block 1's proof of work: %v\n\n\n", blockchain.chain[1].pow)

	for {
		var choice string
		fmt.Print("Do you want to add a new block? (y/n): ")
		fmt.Scanln(&choice)

		if choice == "y" {
			var from string
			var to string
			var amount float64
			fmt.Print("From: ")
			fmt.Scanln(&from)
			fmt.Print("To: ")
			fmt.Scanln(&to)
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			blockchain.addBlock(from, to, amount)
		}
		if choice == "n" {
			break
		}
	}

	for {
		var choice string 
		fmt.Print("Do you want to check data on g-blockchain? (y/n): ")
		fmt.Scanln(&choice)

		if choice == "y" {
			blockchain.getData()
		}
		if choice == "n" {
			break
		}
	}

	for {
		var choice string
		fmt.Print("Do you want to check if the blockchain is valid? (y/n): ")
		fmt.Scanln(&choice)

		if choice == "y" {
			fmt.Printf("%v\n", blockchain.isValid())
		}
		if choice == "n" {
			break
		}
	}
}