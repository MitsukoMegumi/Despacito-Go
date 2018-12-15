package types

import (
	"time"

	"github.com/dowlandaiello/Despacito-Go/src/common"
)

// Transaction - defines transactions on blockchain
type Transaction struct {
	Amount int `json:"amount"`

	SenderAddr   string `json:"sender"`
	ReceiverAddr string `json:"receiver"`

	Timestamp time.Time `json:"timestamp"`

	BlockReward bool `json:"blockreward"`

	Hash string `json:"hash"`
}

// NewTransaction - creates new instance of transaction struct
func NewTransaction(amount int, sender string, receiver string, reward bool) (*Transaction, error) {
	tx := Transaction{Amount: amount, SenderAddr: sender, ReceiverAddr: receiver, Timestamp: time.Now().UTC(), BlockReward: reward, Hash: ""}

	hash, err := common.SHA256(tx)

	if err != nil {
		return nil, err
	}

	tx.Hash = hash

	return &tx, nil
}
