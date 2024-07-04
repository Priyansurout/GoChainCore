package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const TargetBits = 10

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(Block *Block) *ProofOfWork {

	Target := big.NewInt(1)
	Target.Lsh(Target, uint(256-TargetBits))
	return &ProofOfWork{Block, Target}

}

// / it is metho work on proofOfWork it take nonce and mixe with Block data and give the data in from of []byte
func (pow *ProofOfWork) dataHashed(nonce int) []byte {

	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", pow.Block.Index)),
		[]byte(pow.Block.Timestamp),
		[]byte(pow.Block.Data),
		[]byte(pow.Block.PrevHash),
		[]byte(fmt.Sprintf("%d", nonce)),
	}, []byte{})

	return data
}

func (pow *ProofOfWork) Run() {
	var hashInt big.Int
	var hash [32]byte

	for nonce := 0; ; nonce++ {
		data := pow.dataHashed(nonce)
		hash = sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])

		fmt.Printf("Nonce: %d\n", nonce)
		fmt.Printf("Data: %s\n", data)
		fmt.Printf("Hash: %x\n", hash)
		fmt.Printf("HashInt: %d\n", &hashInt)
		fmt.Printf("Target: %d\n\n", pow.Target)

		if hashInt.Cmp(pow.Target) == -1 {
			fmt.Println("Valid hash found!")
			pow.Block.Hash = hex.EncodeToString(hash[:])
			break
		}
	}
}

func main() {
	Block := &Block{
		Index:     1,
		Timestamp: "2024-07-02 00:00:00",
		Data:      "Sample Block Data best in this waorld",
		PrevHash:  "0000000000000000",
	}

	pow := NewProofOfWork(Block)
	pow.Run()

	fmt.Printf("Block Hash: %s\n", Block.Hash)
}
