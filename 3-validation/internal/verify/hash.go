package verify

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
