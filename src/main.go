package main

import (
	"reflect"
	"time"

	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/mitsukomegumi/Despacito-Go/src/core/types"
)

func main() {
	var chain *types.Blockchain

	despacito, err := common.ReadDespacito(common.GetCurrentDir())

	if err != nil {
		panic(err)
	}

	chain, err = types.ReadChainFromMemory(common.GetCurrentDir())

	if err != nil || reflect.ValueOf(chain.Blocks).IsNil() {
		max := 21000000
		zero := 0

		emptyTx := types.Transaction{Amount: 0, SenderAddr: "", ReceiverAddr: "", Timestamp: time.Now().UTC(), BlockReward: false}

		emptyTxArr := []types.Transaction{emptyTx}

		emptyBlock := types.Block{Reward: 0, MinerAddress: "", Version: 0, DespacitoSrc: despacito, Transactions: &emptyTxArr, Hash: ""}

		emptyBlocks := []types.Block{emptyBlock}
		emptyTxs := []types.Transaction{emptyTx}

		nChain := types.Blockchain{Blocks: &emptyBlocks, MaxSupply: &max, CircSupply: &zero, DespacitoSrc: despacito, UncomfTxs: &emptyTxs}

		nChain.WriteChainToMemory(common.GetCurrentDir())

		chain, err = types.ReadChainFromMemory(common.GetCurrentDir())

		if err != nil {
			panic(err)
		}
	}

	chain.Mine("asdf")
}
