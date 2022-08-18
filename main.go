

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		main.go
//	Author:			Zeke
//	Date:			2022.08.15
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"BLC"
	"fmt"
)

func main(){
//	block := BLC.CreateGenesisBlock("Genesis Block")
//	block.SetHash()
//	fmt.Printf("%x\n",block.Hash)
	blockchain := BLC.CreateBlockChain()
	blockchain.Add("block2")
	blockchain.Add("block3")
	fmt.Printf("%v\n",blockchain)
}
