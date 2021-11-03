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
	////TODO generator
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//key := r1.Intn(100)
	//return strconv.Itoa(key)

	bytes := make([]byte, 6) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)
	return key
}
