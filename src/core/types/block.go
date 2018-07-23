package types

import (
	"errors"
	"reflect"

	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/mitsukomegumi/Despacito-Go/src/mutation"
)

// Block - blockchain block
type Block struct {
	Reward       int    `json:"reward"`
	MinerAddress string `json:"minerwallet"`

	Version int `json:"version"`

	DespacitoSrc *[]byte `json:"despacito"`

	Transactions *[]Transaction `json:"transactions"`

	Hash string `json:"hash"`
}

// NewBlock - creates instance of block struct
func NewBlock(reward int, miner string, despacito *[]byte, version int, transactions *[]Transaction) (*Block, error) {

	if reflect.ValueOf(despacito).IsNil() {
		return nil, errors.New("invalid block source")
	} else if len(*despacito) == 0 {
		return nil, errors.New("invalid block source")
	}

	mutation.Mutate(*despacito, common.GlobalMutationSize)
	blck := Block{Reward: reward, MinerAddress: miner, Version: version, DespacitoSrc: despacito, Transactions: transactions, Hash: ""}
	hash, err := common.SHA256(blck)

	if err != nil {
		return nil, err
	}

	blck.Hash = hash

	return &blck, nil
}
