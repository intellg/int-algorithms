package main

import (
	"fmt"
)

func main() {
	//A := make([]int, 100000)
	//A[0] = 1
	//A[99999] = 99999
	A := []int{0, 1, 2, 2, 3, 4, 3, 5, 1, 6, 0, 7, 8}
	result := cal(A)
	fmt.Println(result)
}

func cal(A []int) int {
	// write your code in Go 1.4
	locations := make(map[int]int)
	for _, v := range A {
		locations[v] = 0
	}
	total := len(locations)


	min := len(A)
	for i := 0; i < len(A); i++ {
		found := make(map[int]int)
		for j := i; j-i < min && j < len(A); j++ {
			found[A[j]] = 0
			if len(found) == total {
				min = j - i + 1
				break
			}
		}
	}
	return min
}