package goblockchain

import (
	"fmt"
	"strings"
	"strconv"
)

type Blockchain struct {
	blocks []Block
	difficulty int
}

func (b *Blockchain) GetDifficulty() int {
	return b.difficulty
}

func (b Blockchain) String() string {
	r := "Blockchain - difficulty: " + strconv.Itoa(b.difficulty) + ", blocks: " + strconv.Itoa(len(b.blocks)) + "\n"
	for _, v := range b.blocks {
		r += v.String() + "\n"
	}
	return r
}

func (b *Blockchain) IsValid() bool {
	for i := 0; i < len(b.blocks) - 1; i++ {
		currentBlock := b.blocks[i]
		nextBlock := b.blocks[i + 1]

		if nextBlock.hash != calculateHash(nextBlock.previousHash, nextBlock.timestamp, nextBlock.nonce, nextBlock.data) {
			return false
		}

		if currentBlock.hash != nextBlock.previousHash {
			return false
		}

		if !strings.HasPrefix(nextBlock.hash, b.ExpectedHashPrefix()) {
			return false
		}
	}
	return true
}

func NewBlockchain(difficulty int) Blockchain {
	var b Blockchain
	b.difficulty = difficulty
	return b
}

func (b *Blockchain) LastBlock() Block {
	return b.blocks[len(b.blocks) - 1]
}

func (b *Blockchain) MineBlock(data string) {
	var previousHash string
	if len(b.blocks) == 0 {
		previousHash = fmt.Sprintf("%064x", 0)
	} else {
		previousHash = b.LastBlock().hash
	}

	newBlock := NewBlock(previousHash, data)

	for newBlock.hash[0:b.GetDifficulty()] != b.ExpectedHashPrefix() {
		newBlock.recalculateHash()
	}

	b.blocks = append(b.blocks, newBlock)
}

func (b *Blockchain) ExpectedHashPrefix() string {
	return strings.Repeat("0", b.difficulty)
}