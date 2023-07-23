package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

var hash = sha256.New()

type SHA256 struct {
}

func (sha SHA256) Hash(s string) string {
	return hex.EncodeToString(hash.Sum([]byte(s)))
}
