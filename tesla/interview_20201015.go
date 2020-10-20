package main

import "fmt"

func main() {
	var nums []int
	sortedNums := sort(nums)
	// Let's start from the sorted num array
	sortedNums = []int{0, 1, 2, 3, 9, 9, 9, 7, 8, 9, 10}
	target := 7
	result := twoSum(sortedNums, target)
	fmt.Println(result, target)
}

func calculate2(sortedNums []int, target int) (indices []int) {
	for i := 0; i < len(sortedNums); i++ {
		if i > target {
			break
		}

		min := i + 1
		max := len(sortedNums) - 1
		previousJ := -1
		for {
			j := (min + max) / 2
			if j == previousJ {
				break
			}
			previousJ = j

			val := sortedNums[i] + sortedNums[j]
			if val == target {
				return []int{i, j}
			} else if val > target {
				max = j - 1
			} else { // val < target
				min = j + 1
			}
		}
	}
	return nil
}

func sort(nums []int) (sortedNum []int) {
	// This is a dummy function
	return nums
}

//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//Input: nums = [3,2,4], target = 6 Output:... by Iqbal Mohideen
//Iqbal Mohideen8:27 AM
//Input: nums = [3,2,4], target = 6
//Output: [1,2]

func twoSum(nums []int, target int) []int {
	var hashmap = map[int]int{}

	for idx, val := range nums {
		obj := target - val

		i, ok := hashmap[obj]
		if ok {
			return []int{idx, i}
		} else {
			hashmap[val] = idx
		}
	}
	return nil
}
