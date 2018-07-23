package types

import (
	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/mitsukomegumi/Despacito-Go/src/mutation"
)

// Blockchain - chain representation of blocks, mutations, txs
type Blockchain struct {
	Blocks *[]Block `json:"blocks"`

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
		if len(*blockchain.Blocks) > 0 {
			blocks := *blockchain.Blocks
			dest, err := NewBlock(10, minerWallet, blocks[len(blocks)-1].DespacitoSrc, blocks[len(blocks)-1].Version, blockchain.UncomfTxs)

			if err != nil {
				return err
			}

			mutation.Mutate(*dest.DespacitoSrc, 4)

			mErr := mutation.VerifyMutation(*dest.DespacitoSrc)

			if mErr == nil {
				transcodeable = true
			}
		} else {
			despacito, err := common.ReadDespacito(common.GetCurrentDir())

			if err != nil {
				return err
			}

			dest, err := NewBlock(10, minerWallet, despacito, 0, blockchain.UncomfTxs)

			if err != nil {
				return err
			}

			mutation.Mutate(*dest.DespacitoSrc, 4)

			mErr := mutation.VerifyMutation(*dest.DespacitoSrc)

			if mErr == nil {
				transcodeable = true
			}
		}
	}
	return nil
}

// PublishBlock - push specified block to blockchain
func (blockchain Blockchain) PublishBlock(block Block) {

}
