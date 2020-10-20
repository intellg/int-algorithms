package main

import (
	"fmt"
	"strconv"
)

func main() {
	A := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1}
	result, total := solution(A)
	fmt.Println(result)
	fmt.Println(total)
}

func Solution(A []int) int {
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

func solution(A []int) (result []int, total int) {
	if len(A) == 0 {
		return nil, 0
	}

	init := A[0]
	flag := init
	counters := make([]int, 1)
	index := 0
	for _, v := range A {
		if v == flag {
			counters[index] ++
		} else {
			flag = v
			index++
			counters = append(counters, 1)
		}
	}

	i1 := 0
	if init == 1 {
		i1 = 1
	}
	i0 := len(counters) - 1
	if len(counters)%2 != i1 {
		i0--
	}

	for i1 < i0 {
		if counters[i1] <= counters[i0] {
			total += counters[i1]
			i1 += 2
		} else {
			total += counters[i0]
			i0 -= 2
		}
	}

	return
}

// 2nd round interview
func calculate(expression string) (result int, index int) {
	numbers := make([]int, 0)
	actions := make([]int32, 0)
	for i, v := range expression {
		numberFrom := -1
		if v == '1' { // TODO: is number
			if numberFrom == -1 {
				numberFrom = i
			}
		} else if v == '+' { // TODO: + - * /
			number, _ := strconv.Atoi(expression[numberFrom : i-1])
			numbers = append(numbers, number)
			actions = append(actions, v)
		} else if v == '(' {
			result, index := calculate(expression[i:])
			numbers = append(numbers, result)
			if index != len(expression) {
				i = index
			}
		} else if v == ')' {
			return doCal(numbers, actions), i
		}
	}
	return doCal(numbers, actions), len(expression)
}

func doCal(numbers []int, actions []int32) int {
	for i, v := range actions {
		if v == '*' {
			numbers[i] = numbers[i] * numbers[i+1]
			// 删除numbers[i+1] numbers.pop(i+1)
		} else if v == '/' {
			numbers[i] = numbers[i] / numbers[i+1]
			// 删除numbers[i+1]
		}
	}
	for i, v := range actions {
		if v == '+' {
			numbers[i] = numbers[i] + numbers[i+1]
			// 删除numbers[i+1]
		} else if v == '-' {
			numbers[i] = numbers[i] - numbers[i+1]
			// 删除numbers[i+1]
		}
	}

	return numbers[0]
}
