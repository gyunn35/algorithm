package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		id   int
		strs []string
	}{
		{
			id:   1,
			strs: []string{"apple", "banana", "watermelon"},
		},
		{
			id:   2,
			strs: []string{"|||||", "||||||", "||||||||||"},
		},
		{
			id:   3,
			strs: []string{"", "", ""},
		},
	}

	var codec Codec

	for _, test := range tests {
		assert.Equal(t, test.strs, codec.Decode(codec.Encode(test.strs)))
	}
}
