package chapter04

import "math"

func findMaxCrossSubarray(A []int, low, mid, high int) (maxLeft, maxRight, sum int) {
	leftSum := math.MinInt64
	sum = 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := math.MinInt64
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	sum = leftSum + rightSum
	return
}

func findMaximumSubarray(A []int, low, high int) (from, to, sum int) {
	if low == high {
		from = low
		to = low
		sum = A[low]
		return
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaximumSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossSubarray(A, low, mid, high)
	if leftSum > rightSum {
		if leftSum > crossSum {
			return leftLow, leftHigh, leftSum
		}
	} else if rightSum > crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}
