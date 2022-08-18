//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		pow.go
//	Author:			Zeke
//	Date:			2022.08.18
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package BLC

import (
	"math/big"
)

const targetBit = 16

type POW struct {
	Block	*Block
	target	*big.Int
}

func NewPOW(block *Block) *POW{
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBit)
	return &POW{
		Block:	block
		target:	target
	}
}

func (pow *POW) Run() {
	
}

