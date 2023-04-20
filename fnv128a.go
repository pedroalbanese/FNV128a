package main

import (
	"hash"
)

type fnv128a struct {
	h1, h2 uint64
}

func NewFNV128a() hash.Hash {
	var f fnv128a
	f.Reset()
	return &f
}

func (f *fnv128a) Sum(b []byte) []byte {
	s := make([]byte, 16)
	encodeUint64(s, f.h1)
	encodeUint64(s[8:], f.h2)
	return append(b, s...)
}

func (f *fnv128a) Reset() {
	f.h1 = 0x6c62272e07bb0142
	f.h2 = 0x62d4087733b4b9ae
}

func (f *fnv128a) Size() int {
	return 16
}

func (f *fnv128a) BlockSize() int {
	return 1
}

func (f *fnv128a) Write(data []byte) (int, error) {
	const (
		prime1 = 0x0000013b
		prime2 = 0x000000b3
		prime3 = 0x00000063
		prime4 = 0x00000029
		prime5 = 0x00000011
	)

	var h1, h2 = f.h1, f.h2

	for _, c := range data {
		h1 ^= uint64(c)
		h1 *= prime1
		h2 ^= h1
		h2 *= prime2

		h1 ^= h2
		h1 *= prime3
		h2 ^= h1
		h2 *= prime4

		h1 ^= h2
		h1 *= prime5
		h2 ^= h1
		h2 *= prime5
	}

	f.h1, f.h2 = h1, h2
	return len(data), nil
}

func encodeUint64(dst []byte, x uint64) {
	dst[0] = byte(x >> 56)
	dst[1] = byte(x >> 48)
	dst[2] = byte(x >> 40)
	dst[3] = byte(x >> 32)
	dst[4] = byte(x >> 24)
	dst[5] = byte(x >> 16)
	dst[6] = byte(x >> 8)
	dst[7] = byte(x)
}
