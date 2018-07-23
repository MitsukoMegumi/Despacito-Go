package main

import (
	"testing"

	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/mitsukomegumi/Despacito-Go/src/core/types"
)

func TestNewBlock(t *testing.T) {
	despacito, err := common.ReadDespacito(common.GetCurrentDir())

	test, err := types.NewBlock(10, "asdfasdf", despacito, 10, nil)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(test)
}
