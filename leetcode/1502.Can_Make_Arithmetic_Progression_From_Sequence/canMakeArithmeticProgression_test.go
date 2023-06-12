package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		id     int
		input  []int
		output bool
	}{
		{
			id:     1,
			input:  []int{1, 2, 4},
			output: false,
		},
		{
			id:     2,
			input:  []int{3, 5, 1},
			output: true,
		},
		{
			id:     3,
			input:  []int{0, 0},
			output: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, canMakeArithmeticProgression(test.input), test.id)
		assert.Equal(t, test.output, canMakeArithmeticProgression1(test.input), test.id)
	}
}
