package transaction

import (
	"time"
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
