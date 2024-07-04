package main

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {

	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {

	return NewBlock("Genesis Block", "", 0)
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1)
	bc.Blocks = append(bc.Blocks, newBlock)
}
