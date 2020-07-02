package chapter02

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	numsL := make([]int, mid)
	copy(numsL, nums[:mid])
	numsR := make([]int, len(nums)-mid)
	copy(numsR, nums[mid:])
	sortedL := mergeSort(numsL)
	sortedR := mergeSort(numsR)

	i := 0
	j := 0
	k := 0
	for i < len(sortedL) && j < len(sortedR) {
		if sortedL[i] <= sortedR[j] {
			nums[k] = sortedL[i]
			i++
		} else {
			nums[k] = sortedR[j]
			j++
		}
		k++
	}

	if i == len(sortedL) { // sortedR has remaining data
		copy(nums[k:], sortedR[j:])
	} else { // sortedL has remaining data
		copy(nums[k:], sortedL[i:])
	}

	return nums
}
