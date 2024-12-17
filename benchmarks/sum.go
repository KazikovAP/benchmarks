package benchmarks

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func SumWithChannel(slice []int) int {
	numWorkers := runtime.NumCPU()
	ch := make(chan int, numWorkers)

	wg := sync.WaitGroup{}

	chunkSize := (len(slice) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := min(start+chunkSize, len(slice))

		go func(s []int) {
			defer wg.Done()

			for _, v := range s {
				ch <- v
			}
		}(slice[start:end])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for v, ok := <-ch; ok; v, ok = <-ch {
		total += v
	}

	return total
}

func SumWithMutex(slice []int) int {
	var mu sync.Mutex

	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()

	total := 0
	chunkSize := (len(slice) + numWorkers - 1) / numWorkers

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := min(start+chunkSize, len(slice))

		go func(s []int) {
			defer wg.Done()

			for _, v := range s {
				mu.Lock()
				total += v
				mu.Unlock()
			}
		}(slice[start:end])
	}

	wg.Wait()

	return total
}

func SumWithWaitGroup(slice []int) int {
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	chunkSize := (len(slice) + numWorkers - 1) / numWorkers
	partialSums := make([]int, numWorkers)

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		go func(idx int, s []int) {
			defer wg.Done()

			for _, v := range s {
				partialSums[idx] += v
			}
		}(i, slice[start:end])
	}

	wg.Wait()

	total := 0
	for _, v := range partialSums {
		total += v
	}

	return total
}

func SumWithAtomic(slice []int) int {
	var total atomic.Int64

	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	chunkSize := (len(slice) + numWorkers - 1) / numWorkers

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		go func(s []int) {
			defer wg.Done()

			for _, v := range s {
				total.Add(int64(v))
			}
		}(slice[start:end])
	}

	wg.Wait()

	return int(total.Load())
}
