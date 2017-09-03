package utils

type Hasher interface {
	Hash64(data string) uint64
}
