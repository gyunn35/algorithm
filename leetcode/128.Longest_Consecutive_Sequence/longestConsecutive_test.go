package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		id     int
		input  []int
		output int
	}{
		{
			id:     1,
			input:  []int{1, 2, 3, 4, 6, 6, 6},
			output: 4,
		},
		{
			id:     2,
			input:  []int{},
			output: 0,
		},
		{
			id:     3,
			input:  []int{1},
			output: 1,
		},
		{
			id:     4,
			input:  []int{0, 0},
			output: 1,
		},
		{
			id:     5,
			input:  []int{1, 0, 1, 2},
			output: 3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, longestConsecutive(test.input), test.id)
	}
}
