package mutation

import (
	"github.com/mitsukomegumi/DespacitoNet-Go/src/common"
)

// Mutate - mutate specified byte array
func Mutate(b []byte, size int) {
	mutatedSize := 0

	for mutatedSize != size {
		randSelector := common.RandomBetween(0, len(b)-size)
		randVal := common.RandomValue(size)

		b[randSelector+mutatedSize] = randVal[0]
		mutatedSize++
	}
}
