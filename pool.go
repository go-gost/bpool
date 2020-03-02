package bpool

import (
	pool "github.com/libp2p/go-buffer-pool"
)

var (
	DefaultBufferPool BufferPool = new(pool.BufferPool)
)

func Get(size int) []byte {
	return DefaultBufferPool.Get(size)
}

func Put(b []byte) {
	DefaultBufferPool.Put(b)
}

type BufferPool interface {
	Get(size int) []byte
	Put(b []byte)
}
