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
	transactions []*Transactions
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transactions) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.transactions = transactions
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.timeStamp)
	fmt.Printf("nonce         %d\n", b.nonce)
	fmt.Printf("previous_hash %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64           `json:"timestamp"`
		Nonce        int             `json:"nonce"`
		PreviousHash [32]byte        `json:"previous_hash"`
		Transactions []*Transactions `json:"transactions"`
	}{
		Timestamp:    b.timeStamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transactions
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	genesis := NewBlock(0, [32]byte{}, nil)
	bc.chain = append(bc.chain, genesis)
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transactions{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Transactions struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transactions {
	return &Transactions{sender, recipient, value}
}

func (t *Transactions) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_address     %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address     %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value     %.1f\n", t.value)
}

func (t *Transactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"reciever_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockChain()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash1 := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash1)

	previousHash2 := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash2)

	blockChain.Print()
}
