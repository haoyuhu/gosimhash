package utils

import (
	"hash"
	"github.com/dchest/siphash"
)

type SipHasher struct {
	siphash hash.Hash64
}

func NewSipHasher(key []byte) *SipHasher {
	return &SipHasher{siphash: siphash.New(key)}
}

func (hasher *SipHasher) Hash64(data string) uint64 {
	hasher.siphash.Reset()
	hasher.siphash.Write([]byte(data))
	return hasher.siphash.Sum64()
}
