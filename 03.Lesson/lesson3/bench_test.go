package lesson3

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < 10; i++ {
		b.Run(fmt.Sprintf("Count workers %d", i), func(b *testing.B) {
			WorkerPool(i)
		})
	}
}
