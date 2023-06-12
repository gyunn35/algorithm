package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{6, 3, 2},
		},
		{
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, productExceptSelf(test.input))
	}
}
