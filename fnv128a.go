package main

import (
	"hash"
	"hash/fnv"
)

type meuHash struct {
	hash.Hash
}

func New() hash.Hash {
	h := fnv.New128a()
	return &meuHash{h}
}

func (h *meuHash) Sum(b []byte) []byte {
	return h.Hash.Sum(b)
}

func (h *meuHash) Write(b []byte) (int, error) {
	return h.Hash.Write(b)
}