package blockchain

import (
	"github.com/mitsukomegumi/DespacitoNet-Go/src/block"
	"github.com/mitsukomegumi/DespacitoNet-Go/src/common"
	"github.com/mitsukomegumi/DespacitoNet-Go/src/transaction"
)

// Blockchain - chain representation of blocks, mutations, txs
type Blockchain struct {
	Blocks *[]block.Block `json:"blocks"`

	MaxSupply  *int `json:"maxsupply"`
	CircSupply *int `json:"circulatingsupply"`

	Mutations *int `json:"mutations"`

	DespacitoSrc *[]byte `json:"despacito"`

	UncomfTxs *[]transaction.Transaction
}

// Mine - infinitely mine
func (blockchain Blockchain) Mine(minerWallet string) error {
	transcodeable := false
	for !transcodeable {
		if len(*blockchain.Blocks) > 0 {
			blocks := *blockchain.Blocks
			block.NewBlock(10, minerWallet, blocks[len(blocks)-1].DespacitoSrc, blocks[len(blocks)-1].Version, blockchain.UncomfTxs)
		} else {
			despacito, err := common.ReadDespacito(common.GetCurrentDir())

			if err != nil {
				return err
			}

			block.NewBlock(10, minerWallet, despacito, 0, blockchain.UncomfTxs)
		}
	}
	return nil
}
