package bpool

import "sync"

const (
	tinyBufferSize   = 512
	smallBufferSize  = 2 * 1024
	mediumBufferSize = 8 * 1024
	largeBufferSize  = 32 * 1024
)

var (
	defaultBufferPool = New()
)

func Get(size int) []byte {
	return defaultBufferPool.Get(size)
}

func Put(b []byte) {
	defaultBufferPool.Put(b)
}

type BufferPool struct {
	tiny   sync.Pool
	small  sync.Pool
	medium sync.Pool
	large  sync.Pool
}

func New() *BufferPool {
	return &BufferPool{
		tiny: sync.Pool{
			New: func() interface{} {
				return make([]byte, tinyBufferSize)
			},
		},
		small: sync.Pool{
			New: func() interface{} {
				return make([]byte, smallBufferSize)
			},
		},
		medium: sync.Pool{
			New: func() interface{} {
				return make([]byte, mediumBufferSize)
			},
		},
		large: sync.Pool{
			New: func() interface{} {
				return make([]byte, largeBufferSize)
			},
		},
	}
}

func (p *BufferPool) Get(size int) []byte {
	if size <= 0 {
		return nil
	}
	if size <= tinyBufferSize {
		return p.tiny.Get().([]byte)
	}
	if size <= smallBufferSize {
		return p.small.Get().([]byte)
	}
	if size <= mediumBufferSize {
		return p.medium.Get().([]byte)
	}
	if size <= largeBufferSize {
		return p.large.Get().([]byte)
	}
	return make([]byte, size)
}

func (p *BufferPool) Put(b []byte) {
	switch cap(b) {
	case tinyBufferSize:
		p.tiny.Put(b)
	case smallBufferSize:
		p.small.Put(b)
	case mediumBufferSize:
		p.medium.Put(b)
	case largeBufferSize:
		p.large.Put(b)
	default:
		// ignored
	}
}
