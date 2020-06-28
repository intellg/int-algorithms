package chapter02

func insertionSort(nums []int) {
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		// Insert A[j] into the sorted sequence A[1..j-1]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i--
		}
		nums[i+1] = key
	}
}
