package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timeStamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.timeStamp)
	fmt.Printf("nonce     %d\n", b.nonce)
	fmt.Printf("previous_hash     %x\n", b.previousHash)
	fmt.Printf("transactions     %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64 `json:"timestamp"`
		Nonce        int  `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json: "transactions"`
	}{
		Timestamp:    b.timeStamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
	/*
		blockChain := NewBlockChain()
		blockChain.Print()
		blockChain.CreateBlock(5, "hash 1")
		blockChain.Print()
		blockChain.CreateBlock(2, "hash 2")
		blockChain.Print()
	*/
}
