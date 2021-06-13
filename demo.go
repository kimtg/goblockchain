package main

import (
	"fmt"
	"strconv"

	"github.com/kimtg/goblockchain/goblockchain"
)

func main() {
	bc := goblockchain.NewBlockchain(5)
	fmt.Println("Simple blockchain demo")
	fmt.Println()

	fmt.Println("Mining first block...")
	bc.MineBlock("First block")
	fmt.Println("Block mined.")
	fmt.Println(bc.LastBlock())
	fmt.Println()

	fmt.Println("Mining second block...")
	bc.MineBlock("Second block")
	fmt.Println("Block mined.")
	fmt.Println(bc.LastBlock())
	fmt.Println()

	fmt.Println("Mining third block...")
	bc.MineBlock("Third block")
	fmt.Println("Block mined.")
	fmt.Println(bc.LastBlock())

	fmt.Println()

	fmt.Println("Is blockchain valid? " + strconv.FormatBool(bc.IsValid()))
	fmt.Println()

	fmt.Println("Printing blockchain...")
	fmt.Println(bc)
}
