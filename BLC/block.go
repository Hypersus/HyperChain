

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		Block.go
//	Author:			Zeke
//	Date:			2022.08.15
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package BLC

import (
	"time"
	//"fmt"
	"strconv"
	"crypto/sha256"
	"bytes"
	"encoding/json"
)

type Block struct {
	// block height
	Height			int64
	// hash of last Block
	PrevBlockHash 	[]byte
	// transaction data
	Data 			[]byte
	TimeStamp		int64
	// hash of this Block
	Hash			[]byte
	//
	Nonce			int64
}

func (block *Block) String() string {
	//var str string
	bytes, err := json.MarshalIndent(block,"","\t")
	if err != nil {
		panic("Encode struct failed")
	}
	return string(bytes)
	//str += "\n"
	//str += fmt.Sprintf("\tHeight: %d\n", block.Height)
	//str += fmr.Sprintf("\tData: ")
}


func NewBlock(data string, height int64, prevBlockHash []byte) *Block{
	block := &Block{
		Height:			height,
		PrevBlockHash:	prevBlockHash,
		Data:			[]byte(data),
		TimeStamp:		time.Now().Unix(),
		Hash:			nil,
		Nonce:			0,
	}
	
	// set valid hash and nonce
	pow := NewPOW(block)

	hash, nonce := pow.Run()

	return block
}

func (block *Block) SetHash(){
	// convert timestamp to byte slice (ASCII format)
	timeString := strconv.FormatInt(block.TimeStamp,2)
	timeBytes := []byte(timeString)
	// convert height from int64 to byte slice
	height := Int2Hex(block.Height)
	blockBytes := bytes.Join([][]byte{height,block.PrevBlockHash,block.Data,timeBytes,block.Hash},nil)
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}


func CreateGenesisBlock(data string) *Block{
	return NewBlock(data, 1, make([]byte,32))	
}
