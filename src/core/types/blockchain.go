package types

import (
	"github.com/mitsukomegumi/Despacito-Go/src/common"
)

// Blockchain - chain representation of blocks, mutations, txs
type Blockchain struct {
	Blocks       []*Block `json:"blocks"`
	BlockMempool []*Block `json:"block mempool"`

	MaxSupply  *int `json:"maxsupply"`
	CircSupply *int `json:"circulatingsupply"`

	Mutations *int `json:"mutations"`

	DespacitoSrc *[]byte `json:"despacito"`

	UncomfTxs *[]Transaction
}

// Mine - infinitely mine
func (blockchain Blockchain) Mine(minerWallet string) error {
	transcodeable := false
	for !transcodeable {
		if len(blockchain.Blocks) > 0 {
			blocks := blockchain.Blocks
			NewBlock(10, minerWallet, blocks[len(blocks)-1].DespacitoSrc, blocks[len(blocks)-1].Version, blockchain.UncomfTxs)
		} else {
			despacito, err := common.ReadDespacito(common.GetCurrentDir())

			if err != nil {
				return err
			}

			NewBlock(10, minerWallet, despacito, 0, blockchain.UncomfTxs)
		}
	}
	return nil
}
