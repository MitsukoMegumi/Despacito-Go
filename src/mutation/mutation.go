package mutation

import (
	"fmt"

	"github.com/dowlandaiello/Despacito-Go/src/common"
)

// Mutate - mutate specified byte array
func Mutate(b []byte, size int) {
	mutatedSize := 0

	var randVal []byte
	var randSelector int

	for mutatedSize != size {
		randSelector = common.RandomBetween(0, len(b))
		randVal = common.RandomValue(size)

		b[randSelector+mutatedSize] = randVal[0]

		fmt.Println(randVal[0])
		fmt.Println(randSelector)

		mutatedSize++
	}
}
