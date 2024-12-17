package benchmarks_test

import (
	"strconv"
	"testing"

	"github.com/KazikovAP/benchmarks/benchmarks"
)

func generateRandomArray(size int) []int {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i
	}

	return slice
}

func BenchmarkSumWithChannel(b *testing.B) {
	sliceSizes := []int{1000, 10_000, 100_000}
	for _, size := range sliceSizes {
		b.Run("Channel_"+strconv.Itoa(size), func(b *testing.B) {
			slice := generateRandomArray(size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarks.SumWithChannel(slice)
			}
		})
	}
}

func BenchmarkSumWithMutex(b *testing.B) {
	sliceSizes := []int{1000, 10_000, 100_000}
	for _, size := range sliceSizes {
		b.Run("Mutex_"+strconv.Itoa(size), func(b *testing.B) {
			slice := generateRandomArray(size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarks.SumWithMutex(slice)
			}
		})
	}
}

func BenchmarkSumWithWaitGroup(b *testing.B) {
	sliceSizes := []int{1000, 10_000, 100_000}
	for _, size := range sliceSizes {
		b.Run("WaitGroup_"+strconv.Itoa(size), func(b *testing.B) {
			slice := generateRandomArray(size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarks.SumWithWaitGroup(slice)
			}
		})
	}
}

func BenchmarkSumWithAtomic(b *testing.B) {
	sliceSizes := []int{1000, 10_000, 100_000}
	for _, size := range sliceSizes {
		b.Run("Atomic_"+strconv.Itoa(size), func(b *testing.B) {
			slice := generateRandomArray(size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarks.SumWithAtomic(slice)
			}
		})
	}
}
