package bpool

import (
	"math/rand"
	"testing"
	"time"
)

func TestBufferPool(t *testing.T) {
	for i := 0; i < 1<<16; i++ {
		Put(make([]byte, i))
		b := Get(i)
		if len(b) != i {
			t.Error("get: wrong size")
		}
		Put(b)
	}
}

func BenchmarkBufferPool(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		n := rand.Intn(20)
		size := 1 << n
		buf := Get(size)
		if len(buf) != size {
			b.Error("get: wrong size")
		}
		Put(buf)
	}
}
