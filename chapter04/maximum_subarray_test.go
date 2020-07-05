package chapter04

import (
	"testing"
)

type maximumSubarrayData struct {
	nums []int
	from int
	to   int
	sum  int
}

var maximumSubarrayItems = []maximumSubarrayData{
	{
		nums: []int{1, -2, 3, -4, 5, 6, -7, 4, 5},
		from: 4,
		to:   8,
		sum:  13,
	},
	{
		nums: []int{1, -4, 3, -4},
		from: 2,
		to:   2,
		sum:  3,
	},
}

func TestFindMaximumSubarray(t *testing.T) {
	for i, testItem := range maximumSubarrayItems {
		from, to, sum := findMaximumSubarray(testItem.nums, 0, len(testItem.nums)-1)
		if from != testItem.from {
			t.Errorf("Invalid 'from' for test item[%d]: expect %d but get %d", i, testItem.from, from)
			return
		}
		if to != testItem.to {
			t.Errorf("Invalid 'to' for test item[%d]: expect %d but get %d", i, testItem.to, to)
			return
		}
		if sum != testItem.sum {
			t.Errorf("Invalid 'sum' for test item[%d]: expect %d but get %d", i, testItem.sum, sum)
			return
		}
	}
}
