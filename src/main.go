package main

import (
	"fmt"

	"github.com/dowlandaiello/Despacito-Go/src/common"
	"github.com/dowlandaiello/Despacito-Go/src/core/types"
)

func main() {
	despacito, err := common.ReadDespacito(common.GetCurrentDir()) // Read local despacito copy

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	test, err := types.NewBlock(10, "asdfasdf", despacito, 10, nil) // Generate test block

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	fmt.Println(test) // Log test block
}
