package models

import (
	"blockchain/digest"
	"fmt"
	"os"
)

type block interface {
	mineBlock()
	getHash() string
	getNonce() string
	toString() string
	fullBlock(string) string
	getPreviousBlockHash() string
}
type messageBlock struct {
	previousBlockHash string
	message string
	nonce string
	hash string
}

func NewMessageBlock(previousBlock block, message string) block {
	prevBlockHash := ""
	if previousBlock != nil {
		prevBlockHash = previousBlock.getHash()
	}

	var mb = messageBlock {
		previousBlockHash:prevBlockHash,
		message:message,
	}
	mb.mineBlock()
	mb.hash = digest.Hash(mb.fullBlock(mb.nonce))
	var b block = &mb

	return b
}

func (block *messageBlock) getHash() string {
	return block.hash
}

func (block *messageBlock) getNonce() string {
	return block.nonce
}

func (block *messageBlock) getPreviousBlockHash() string {
	return block.previousBlockHash
}

func (block *messageBlock) mineBlock () {
	var err error
	block.nonce,err = digest.FindNonce(block.message + block.previousBlockHash)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (block *messageBlock) toString() string {
	str := "********************************************************************************" + "\n"
	str += "* Message: " + block.message + "\n"
	str += "* Nonce: " + block.nonce + "\n"
	str += "* Hash: " + block.hash + "\n"
	str += "* Previous Block Hash: " + block.previousBlockHash + "\n"
	str += "********************************************************************************" + "\n"

	return str
}

func (block *messageBlock) fullBlock(nonce string ) string {
	return block.message + block.previousBlockHash + nonce
}

type BlockChain struct {
	blocks []block
}

func (bc *BlockChain) Init (msg string) {
	bc.blocks = make([]block, 0)
	bc.blocks = append(bc.blocks, NewMessageBlock(nil, msg))
}

func (bc *BlockChain) AddToChain (msg string) {
	bc.blocks = append(bc.blocks, NewMessageBlock(bc.blocks[len(bc.blocks)-1], msg))
}

func (bc *BlockChain) ToString() {
	for block := range bc.blocks {
		fmt.Println(bc.blocks[block].toString())
	}
}

func (bc *BlockChain) IsValid() bool {
	for block := range bc.blocks {
		if !digest.ValidNonce(bc.blocks[block].getNonce(), bc.blocks[block].fullBlock("")) ||
			(bc.blocks[block].getPreviousBlockHash() != "" &&
			bc.blocks[block].getPreviousBlockHash() != bc.blocks[block-1].getHash()) {

			return false
		}
	}
	return true
}
