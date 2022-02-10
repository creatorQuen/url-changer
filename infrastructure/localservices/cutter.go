package localservices

import (
	"encoding/hex"
	"log"
	"math/rand"
)

type KeyGenerator struct {
	myRand myRandInterface
}

func (c *KeyGenerator) Generate() string {
	bytes := make([]byte, 8)
	if _, err := c.myRand.Read(bytes); err != nil {
		log.Fatal("Generate: ", err)
		return "Generate error."
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
