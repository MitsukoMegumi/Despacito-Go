package mutation

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

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

// VerifyMutation - verify byte array for encoding
func VerifyMutation(b []byte) error {
	err := SaveMutation(b)

	if err != nil {
		return err
	}

	rErr := ReadMutation(b)

	if rErr != nil {
		return rErr
	}

	return nil
}

// ReadMutation - read mutated despacito from current directory
func ReadMutation(b []byte) error {
	file, err := os.Open(common.GetCurrentDir() + "despacito.mp4") // For read access.
	if err != nil {
		return err
	}

	_, err = file.Read(b)
	if err != nil {
		return err
	}

	return nil
}

// SaveMutation - save specified mutation
func SaveMutation(b []byte) error {
	b64data := string(b)
	reader := strings.NewReader(b64data)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	decodedData, err := ioutil.ReadAll(decoder)

	if err != nil {
		return err
	}

	ioutil.WriteFile(common.GetCurrentDir()+"/despacito.mp4", decodedData, 0644)

	return nil
}
