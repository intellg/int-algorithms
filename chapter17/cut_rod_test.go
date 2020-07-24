package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	prices = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30,
		34, 39, 43, 48, 54, 60, 65, 71, 75, 80,
		86, 91, 99, 105, 112, 119, 125, 130, 135, 140}
)

func TestCutRod(t *testing.T) {
	t.Run("test cutRod", func(t *testing.T) {
		result := cutRod(prices, 30)
		assert.Equal(t, 140, result)
	})

	t.Run("test memoizedCutRod", func(t *testing.T) {
		result := memoizedCutRod(prices, 30)
		assert.Equal(t, 140, result)
	})

	t.Run("test bottomUpCutRod", func(t *testing.T) {
		result := bottomUpCutRod(prices, 30)
		assert.Equal(t, 140, result)
	})
}

func BenchmarkCutRod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cutRod(prices, 20)
	}
}

func BenchmarkMemorizedCutRod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoizedCutRod(prices, 20)
	}
}

func BenchmarkBottomUpCutRod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bottomUpCutRod(prices, 20)
	}
}
