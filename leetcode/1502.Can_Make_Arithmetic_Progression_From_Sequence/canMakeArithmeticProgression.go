package leetcode

import (
	"math"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if diff != (arr[i] - arr[i-1]) {
			return false
		}
	}

	return true
}

func canMakeArithmeticProgression1(arr []int) bool {
	m := make(map[int]bool)

	min := math.MaxInt32
	max := math.MinInt32
	for _, n := range arr {
		m[n] = true
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	if max-min == 0 {
		return true
	}

	if len(m) != len(arr) || (max-min)%(len(arr)-1) != 0 {
		return false
	}

	diff := (max - min) / (len(arr) - 1)

	for _, n := range arr {
		if (n-min)%diff != 0 {
			return false
		}
	}

	return true
}
