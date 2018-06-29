package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// SHA256 - generate sha256 string of specified string
func SHA256(b interface{}) string {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(b)

	algorithm := sha256.New()
	algorithm.Write(reqBodyBytes.Bytes())

	return hex.EncodeToString(algorithm.Sum(nil))
}
