package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		id     int
		input  []int
		output [][]int
	}{
		{
			id:     1,
			input:  []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			id:     2,
			input:  []int{0, 0, 1},
			output: [][]int{},
		},
		{
			id:     3,
			input:  []int{0, 0, 0, 0, 0, 0, 0},
			output: [][]int{{0, 0, 0}},
		},
		{
			id:     4,
			input:  []int{-2, -2, -2, 0, 0, 0, 0, 2, 2, 2, 2},
			output: [][]int{{-2, 0, 2}, {0, 0, 0}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, threeSum(test.input), test.id)
	}
}
