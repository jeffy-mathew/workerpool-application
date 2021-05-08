package services

import (
	"encoding/hex"
	"hash"
)

type HasherInterface interface {
	HashInput(input []byte) string
}

type Hasher struct {
	hash hash.Hash
}

func NewHasher(hash hash.Hash) HasherInterface {
	return &Hasher{hash: hash}
}

func (h *Hasher) HashInput(input []byte) string {
	return hex.EncodeToString(h.hash.Sum(input))
}
