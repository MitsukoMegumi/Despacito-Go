package main

import (
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

	if err != nil {
		max := 21000000
		zero := 0

		emptyBlocks := []types.Block{}
		emptyTxs := []types.Transaction{}

		nChain := types.Blockchain{Blocks: &emptyBlocks, MaxSupply: &max, CircSupply: &zero, DespacitoSrc: despacito, UncomfTxs: &emptyTxs}

		chain = &nChain
	}

	chain.Mine("asdf")
}
