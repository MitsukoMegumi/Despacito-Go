package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

// RandomBetween - generates random integer between two integers
func RandomBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
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
