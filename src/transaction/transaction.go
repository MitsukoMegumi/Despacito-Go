package transaction

import (
	"time"
	"github.com/mitsukomegumi/DespacitoNet-Go/src/common"
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

func NewTransaction(amount int, sender string, receiver string, reward bool) *Transaction {
	tx := Transaction{Amount: amount, SenderAddr: sender, ReceiverAddr: receiver, Timestamp: time.Now().UTC(), BlockReward: reward, Hash: ""}

	tx.Hash := common.SHA256(tx)

	return &tx
}
