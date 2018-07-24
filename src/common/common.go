package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// GlobalMutationSize - predefined mutation size
const GlobalMutationSize = 6

// NilValue - global zero value
const NilValue = 0

// RandomBetween - generates random integer between two integers
func RandomBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomValue - generates random byte array in size
func RandomValue(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)

	return token
}

// SHA256 - generate sha256 string of specified string
func SHA256(b interface{}) (string, error) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(b)

	if err != nil {
		return "", err
	}

	algorithm := sha256.New()
	_, err = algorithm.Write(reqBodyBytes.Bytes())

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(algorithm.Sum(nil)), nil
}

// ReadDespacito - attempts to read despacito mp4 file
func ReadDespacito(dir string) (*[]byte, error) {
	data, err := ioutil.ReadFile(dir + "/despacito.mp4")

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetCurrentDir - retrieves current directory
func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

// WriteGob - create gob from specified object, at filePath
func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// ReadGob - read gob specified at path
func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
