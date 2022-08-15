

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		utils.go
//	Author:			Zeke
//	Date:			2022.08.15
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

func Int2Hex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
