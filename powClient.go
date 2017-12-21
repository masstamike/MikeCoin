package main

import (
	"fmt"
	"blockchain/digest"
	"os"
)

func main() {
	str := "MikeCoin, is the best coin!"
	nonce, err := digest.FindNonce(str)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(nonce)
	if digest.ValidNonce(nonce, str) {
		fmt.Println("That's correct!")
	} else {
		fmt.Println("That's wrong!!!")
	}
}
