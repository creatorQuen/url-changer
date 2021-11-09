package localservices

import (
	"encoding/hex"
	"math/rand"
)

type KeyGenerator struct {
}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

func (c KeyGenerator) Generate() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)
	return key
}
