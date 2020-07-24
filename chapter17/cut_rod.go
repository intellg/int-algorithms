package chapter17

import "math"

func cutRod(p []int, n int) int {
	var q int
	if n != 0 {
		q = math.MinInt64
		for i := 1; i <= n; i++ {
			v := p[i] + cutRod(p, n-i)
			if q < v {
				q = v
			}
		}
	}
	return q
}

func memoizedCutRod(p []int, n int) int {
	r := make([]int, n+1)
	for i := 0; i <= n; i++ {
		r[i] = math.MinInt64
	}
	return memoizedCutRodAux(p, r, n)
}

func memoizedCutRodAux(p, r []int, n int) int {
	if r[n] >= 0 {
		return r[n]
	}
	var q int
	if n != 0 {
		q = math.MinInt64
		for i := 1; i <= n; i++ {
			v := p[i] + memoizedCutRodAux(p, r, n-i)
			if q < v {
				q = v
			}
		}
	}
	r[n] = q
	return q
}

func bottomUpCutRod(p []int, n int) int {
	r := make([]int, n+1)
	for j := 1; j <= n; j++ {
		q := math.MinInt64
		for i:=1; i<=j ;i++ {
			v := p[i] + r[j-i]
			if q < v {
				q = v
			}
		}
		r[j] = q
	}
	return r[n]
}