

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		blockchain.go
//	Author:			Zeke
//	Date:			2022.08.17
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package BLC

import (
	"fmt"
)

type BlockChain struct {
	Blocks	[]*Block
}

func (blc *BlockChain) String() string{
	var str string
	for i, block := range blc.Blocks {
		str+=fmt.Sprintf("Block %d: %v\n", i, block)
	}
	return str
}

func CreateBlockChain() *BlockChain {	
	genesisBlock := CreateGenesisBlock("Genesis block created")
	return &BlockChain{[]*Block{genesisBlock}}
}

func (blc *BlockChain) Add(data string) {
	prevBlock := blc.Blocks[len(blc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Height + 1, prevBlock.Hash)  	
	blc.Blocks = append(blc.Blocks, newBlock)
}
