package main

import (
	"fmt"

	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/mitsukomegumi/Despacito-Go/src/core/types"
)

func main() {
	despacito, err := common.ReadDespacito(common.GetCurrentDir())
	test, err := types.NewBlock(10, "asdfasdf", despacito, 10, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(test)
}
