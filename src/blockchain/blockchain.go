package blockchain

import (
	"github.com/mitsukomegumi/DespacitoNet-Go/src/block"
	"github.com/mitsukomegumi/DespacitoNet-Go/src/transaction"
)

// Blockchain - chain representation of blocks, mutations, txs
type Blockchain struct {
	Blocks []*block.Block `json:"blocks"`

	MaxSupply  *int `json:"maxsupply"`
	CircSupply *int `json:"circulatingsupply"`

	Mutations *int `json:"mutations"`

	DespacitoSrc *[]byte `json:"despacito"`

	UncomfTxs []*transaction.Transaction
}

// Mine - infinitely mine
func (blockchain Blockchain) Mine(minerWallet string) error {
	transcodeable := false
	for !transcodeable {
		if len(blockchain.Blocks) > 0 {
			block.NewBlock(10, minerWallet, blockchain.Blocks[len(blockchain.Blocks)-1].DespacitoSrc, blockchain.Blocks[len(blockchain.Blocks)-1].Version, blockchain.UncomfTxs)
		} else {
			block.NewBlock(10, minerWallet)
		}
	}
	return nil
}
