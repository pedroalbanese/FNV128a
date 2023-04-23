# FNV128a
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/FNV128a/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/edgetk?status.png)](http://godoc.org/github.com/pedroalbanese/FNV128a)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/FNV128a/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/FNV128a/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/FNV128a)](https://goreportcard.com/report/github.com/pedroalbanese/FNV128a)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/FNV128a)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/FNV128a)](https://github.com/pedroalbanese/FNV128a/releases)

### FNV128a Non-Cryptographic Recursive Hasher  
[draft-eastlake-fnv-18.txt](https://datatracker.ietf.org/meeting/114/materials/slides-114-secdispatch-the-fnv-non-cryptographic-hash-algorithm-00.pdf) (Under development)

FNV1a 128-bit is a variant of the Fowler–Noll–Vo hash function family that produces a 128-bit hash value. It is designed to be fast and has good dispersion properties, making it suitable for use in hash tables, checksums, and similar applications. It is based on the FNV1a algorithm, which uses the FNV prime and FNV offset values to transform a sequence of bytes into a hash value.
