package goblockchain

import (
	"crypto/sha256"
	"time"
	"fmt"
	"strconv"
)

// Block is a block
type Block struct {
	hash string
	previousHash string
	timestamp time.Time
	nonce int
	data string
}

func (b *Block) GetData() string {
	return b.data
}

func (b *Block) GetNonce() int {
	return b.nonce
}

func NewBlock(previousHash, data string) Block {
	var b Block
	nonce := 0
	now := time.Now()
	b.hash = calculateHash(previousHash, now, nonce, data)
	b.previousHash = previousHash
	b.timestamp = now
	b.nonce = nonce
	b.data = data
	return b
}

func calculateHash(previousHash string, timestamp time.Time, nonce int, data string) string {
	return fmt.Sprintf("%064x", (sha256.Sum256([]byte(previousHash + timestamp.String() + strconv.Itoa(nonce) + data))))
}

func (b *Block) recalculateHash() {
	b.nonce++
	b.timestamp = time.Now()
	b.hash = calculateHash(b.previousHash, b.timestamp, b.nonce, b.data)
}

func (b Block) String() string {
	return "Hash: " + b.hash +
	", Previous hash: " + b.previousHash +
	", Timestamp: " + b.timestamp.String() +
	", Nonce: " + strconv.Itoa(b.nonce) +
	", Data: " + b.data
}