package localservices

import (
	"encoding/hex"
	"log"
	"math/rand"
)

type KeyGenerator struct {
	myRand myRandInterface
}

func (c KeyGenerator) Generate() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)

	return key
}

type myRand struct{}

func NewKeyGenerator() *KeyGenerator {
	myRand := new(myRand)
	return &KeyGenerator{myRand: myRand}
}

func (mr *myRand) Read(b []byte) (n int, err error) {
	return rand.Read(b)
}
