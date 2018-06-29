package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

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
