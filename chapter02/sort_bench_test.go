package chapter02

import (
	"testing"
)

const length = 60000

var testItem = make([]int, length)

func init() {
	for i := 0; i < length; i++ {
		testItem[i] = length - i
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertionSort(testItem)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeSort(testItem)
	}
}
