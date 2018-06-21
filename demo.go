package main

import (
	"./goblockchain"
	"fmt"
	"strconv"
)

func main() {
	var b goblockchain.Block
	bc := goblockchain.NewBlockchain(5)
	fmt.Println("Simple blockchain demo")

	fmt.Println("Mining first block...")
	bc.MineBlock("First block")
	fmt.Println("Block mined.")
	b = bc.LastBlock()
	fmt.Printf("Difficulty: %v, nonce: %v\n", bc.GetDifficulty(), b.GetNonce())

	fmt.Println("Mining second block...")
	bc.MineBlock("Second block")
	fmt.Println("Block mined.")
	b = bc.LastBlock()
	fmt.Printf("Difficulty: %v, nonce: %v\n", bc.GetDifficulty(), b.GetNonce())

	fmt.Println("Mining third block...")
	bc.MineBlock("Third block")
	fmt.Println("Block mined.")
	b = bc.LastBlock()
	fmt.Printf("Difficulty: %v, nonce: %v\n", bc.GetDifficulty(), b.GetNonce())

	fmt.Println("Is blockchain valid? " + strconv.FormatBool(bc.IsValid()))

	fmt.Println("Printing blockchain...");
	fmt.Println(bc)
}