package chapter02

import (
	"testing"
)

type sortData struct {
	nums   []int
	result []int
}

var sortItems = []sortData{
	{
		nums:   []int{2, 1},
		result: []int{1, 2},
	},
	{
		nums:   []int{5, 2, 4, 6, 1, 3},
		result: []int{1, 2, 3, 4, 5, 6},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, testItem := range sortItems {
		insertionSort(testItem.nums)
		if len(testItem.nums) != len(testItem.result) {
			t.Errorf("Invalid for %v\n", testItem.nums)
			return
		}
		for i := 0; i < len(testItem.nums); i++ {
			if testItem.nums[i] != testItem.result[i] {
				t.Errorf("Invalid for %v with index %d\n", testItem.nums, i)
				return
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, testItem := range sortItems {
		mergeSort(testItem.nums)
		if len(testItem.nums) != len(testItem.result) {
			t.Errorf("Invalid for %v\n", testItem.nums)
			return
		}
		for i := 0; i < len(testItem.nums); i++ {
			if testItem.nums[i] != testItem.result[i] {
				t.Errorf("Invalid for %v with index %d\n", testItem.nums, i)
				return
			}
		}
	}
}
