

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

type BlockChain struct {
	Blocks	[]*Block
}

func CreateBlockChain() *BlockChain {	
	genesisBlock := CreateGenesisBlock("Genesis block created")
	return &BlockChain{[]*Block{genesisBlock}}
}

func (blc *Blockchain) Add(data string) {
	
}
