package benchmarks

import "sort"

func StandardSort(slice []int) {
	sort.Ints(slice)
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
