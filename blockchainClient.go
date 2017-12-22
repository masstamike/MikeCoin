package main

import (
	"blockchain/models"
	"fmt"
)

func main() {
	str := "MikeCoin, is the best coin!"
	bc := models.BlockChain{}
	fmt.Println("Initializing blockchain with first node")
	bc.Init(str)
	bc.ToString()
	fmt.Println("Adding second node to blockchain")
	bc.AddToChain("Or is it?")
	bc.ToString()
	fmt.Println(bc.IsValid())
}
