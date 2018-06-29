package block

import "github.com/mitsukomegumi/DespacitoNet-Go/src/transaction"

// Block - blockchain block
type Block struct {
	Reward       int    `json:"reward"`
	MinerAddress string `json:"minerwallet"`

	Version int `json:"version"`

	DespacitoSrc []byte `json:"despacito"`

	Transactions []*transaction.Transaction `json:"transactions"`

	Hash string `json:"hash"`
}
