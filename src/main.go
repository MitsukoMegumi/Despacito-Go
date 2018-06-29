package main

import (
	"fmt"

	"github.com/mitsukomegumi/DespacitoNet-Go/src/block"
)

func main() {
	test, err := block.NewBlock(10, "asdfasdf", nil, 10, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(test)
}
