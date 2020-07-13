package chapter10

import (
	"testing"
)

func TestStack(t *testing.T) {
	S := []int{1, 2, 3}
	S = push(S, 4)
	if len(S) == 4 && S[3] == 4 {
		t.Log("Push is Correct")
	} else {
		t.Fatal("Push is Wrong")
	}

	S, result, _ := pop(S)
	if len(S) == 3 && result == 4 {
		t.Log("Pop is Correct")
	} else {
		t.Fatal("Pop is Wrong")
	}
}
