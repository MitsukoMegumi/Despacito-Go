package mutation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nareix/joy4/av"

	"github.com/mitsukomegumi/Despacito-Go/src/common"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
)

// Mutate - mutate specified byte array
func Mutate(b []byte, size int) (string, error) {

	if size == 0 || len(b) < size {
		return "", errors.New("invalid size or input")
	}

	mutatedSize := 0

	var mutated []string

	for mutatedSize != size {
		randSelector := common.RandomBetween(0, len(b)-size)
		randVal := common.RandomValue(size)

		mutated = append(mutated, string(randVal[0]))

		b[randSelector+mutatedSize] = randVal[0]
		mutatedSize++
	}

	return strings.Join(mutated, " "), nil
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
	format.RegisterAll()
	file, err := avutil.Open("despacito.mp4")

	if err != nil {
		return err
	}

	streams, err := file.Streams()

	if err != nil {
		return err
	}

	stream := streams[0]

	vstream := stream.(av.CodecData)

	fmt.Println(vstream.Type())

	if vstream.Type() != av.H264 {
		return errors.New("invalid")
	}

	ioutil.WriteFile(common.GetCurrentDir()+"\\despacito.mp4", b, 0644)

	return nil
}
