package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func NewBlock(data string, prevhash string, index int) *Block {

	block := &Block{

		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevhash,
	}

	block.Hash = CalculateHash(block)

	return block

}

func CalculateHash(block *Block) string {

	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)

}
