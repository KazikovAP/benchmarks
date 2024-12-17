package benchmarks_test

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"

	"github.com/KazikovAP/benchmarks/benchmarks"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)

	for i := range slice {
		n, _ := rand.Int(rand.Reader, big.NewInt(1000))
		slice[i] = int(n.Int64())
	}

	return slice
}

func BenchmarkSortingAlgorithms(b *testing.B) {
	sizes := []int{1000, 10_000, 100_000}
	algorithms := []struct {
		name string
		fn   func([]int)
	}{
		{"StandardSort", benchmarks.StandardSort},
		{"BubbleSort", benchmarks.BubbleSort},
	}

	for _, size := range sizes {
		slice := generateRandomSlice(size)

		for _, algo := range algorithms {
			b.Run(algo.name+"_"+strconv.Itoa(size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					copySlice := append([]int(nil), slice...)
					algo.fn(copySlice)
				}
			})
		}
	}
}
