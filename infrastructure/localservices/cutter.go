package localservices

import (
	"math/rand"
	"strconv"
	"time"
)

type KeyGenerator struct {
}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

func (c KeyGenerator) Generate() string {
	//TODO generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	key := r1.Intn(100)
	return strconv.Itoa(key)
}
