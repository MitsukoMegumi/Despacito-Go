package types

import (
	"fmt"
	"path/filepath"
	"reflect"

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

	fmt.Println("started mining...")

	for !transcodeable {
		if reflect.ValueOf(blockchain.Blocks).IsNil() {
			blocks := *blockchain.Blocks
			dest, err := NewBlock(10, minerWallet, blocks[len(blocks)-1].DespacitoSrc, blocks[len(blocks)-1].Version, blockchain.UncomfTxs)

			if err != nil {
				return err
			}

			mutatedResult, err := mutation.Mutate(*dest.DespacitoSrc, common.GlobalMutationSize)

			if err != nil {
				return err
			}

			fmt.Println("\nfound mutation: " + mutatedResult)

			mErr := mutation.VerifyMutation(*dest.DespacitoSrc)

			if mErr == nil {
				transcodeable = true
			} else {
				fmt.Println("solution invalid")
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

			mutatedResult, err := mutation.Mutate(*dest.DespacitoSrc, common.GlobalMutationSize)

			if err != nil {
				return err
			}

			fmt.Println("\nfound mutation: " + mutatedResult)

			mErr := mutation.VerifyMutation(*dest.DespacitoSrc)

			if mErr == nil {
				transcodeable = true
			} else {
				fmt.Println("solution invalid")
			}
		}
	}

	blockchain.WriteChainToMemory(common.GetCurrentDir())

	return nil
}

// PublishBlock - push specified block to blockchain
func (blockchain Blockchain) PublishBlock(block Block) {

}

// WriteChainToMemory - create serialized instance of specified chain in specified path (string)
func (blockchain Blockchain) WriteChainToMemory(path string) error {
	err := common.WriteGob(path+filepath.FromSlash("/Chain.gob"), blockchain)

	if err != nil {
		return err
	}

	return nil
}

// ReadChainFromMemory - read serialized object of specified chain from specified path
func ReadChainFromMemory(path string) (*Blockchain, error) {
	tempChain := new(Blockchain)

	err := common.ReadGob(path+filepath.FromSlash("/Chain.gob"), tempChain)
	if err != nil {
		return nil, err
	}

	return tempChain, nil
}
