package chapter02

import (
	"fmt"
	"testing"
)

type insertionSortData struct {
	nums   []int
	result []int
}

var insertionSortItems = []insertionSortData{
	{
		nums:   []int{5, 2, 4, 6, 1, 3},
		result: []int{1, 2, 3, 4, 5, 6},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, testItem := range insertionSortItems {
		insertionSort(testItem.nums)
		if len(testItem.nums) != len(testItem.result) {
			fmt.Printf("Wrong for %v\n", testItem.nums)
			return
		}
		for i := 0; i < len(testItem.nums); i++ {
			if testItem.nums[i] != testItem.result[i] {
				fmt.Printf("Wrong for %v with index %d\n", testItem.nums, i)
				return
			}
		}
	}
	fmt.Println("Correct")
}
