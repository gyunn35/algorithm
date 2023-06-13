package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		id     int
		input  string
		output bool
	}{
		{
			id:     1,
			input:  "A man, a plan, a canal: Panama",
			output: true,
		},
		{
			id:     2,
			input:  "race a car",
			output: false,
		},
		{
			id:     3,
			input:  "0o",
			output: false,
		},
		{
			id:     4,
			input:  "a",
			output: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, isPalindrome(test.input), test.id)
	}
}
