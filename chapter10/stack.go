package chapter10

import "errors"

func stackEmpty(S []int) bool {
	if len(S) == 0 {
		return true
	} else {
		return false
	}
}

func push(S []int, x int) (S_ []int) {
	S_ = append(S, x)
	return
}

func pop(S []int) (S_ []int, result int, err error) {
	if len(S) == 0 {
		err = errors.New("underflow")
		return
	}

	l := len(S)-1
	result = S[l]
	S_ = S[:l]
	return
}